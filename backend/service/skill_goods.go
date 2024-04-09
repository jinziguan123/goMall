/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-08 14:07:29
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-09 00:10:35
 * @FilePath: /goMall/backend/service/skill_goods.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"goMall/backend/pkg/e"
	"goMall/backend/repository/cache"
	"goMall/backend/repository/database/dao"
	"goMall/backend/repository/database/model"
	"goMall/backend/repository/mq"
	"goMall/backend/serializer"
	"log"
	"math/rand"
	"mime/multipart"
	"strconv"
	"time"

	xlsx "github.com/360EntSecGroup-Skylar/excelize"
	logging "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

type SkillGoodsImport struct {
}

// 限购一个
type SkillGoodsService struct {
	SkillGoodsId uint   `json:"skill_goods_id" form:"skill_goods_id"`
	ProductId    uint   `json:"product_id" form:"product_id"`
	BossId       uint   `json:"boss_id" form:"boss_id"`
	AddressId    uint   `json:"address_id" form:"address_id"`
	Key          string `json:"key" form:"key"`
}

func (service *SkillGoodsImport) Import(c context.Context, file multipart.File) serializer.Response {
	xlFile, err := xlsx.OpenReader(file)
	if err != nil {
		logging.Info(err)
	}
	code := e.SUCCESS
	rows := xlFile.GetRows("Sheet1")
	length := len(rows[1:])
	skillGoods := make([]*model.SkillGoods, length, length)
	for index, colCell := range rows {
		if index == 0 {
			continue
		}
		pId, _ := strconv.Atoi(colCell[0])
		bId, _ := strconv.Atoi(colCell[1])
		num, _ := strconv.Atoi(colCell[3])
		money, _ := strconv.ParseFloat(colCell[4], 64)
		skillGood := &model.SkillGoods{
			ProductId: uint(pId),
			BossId:    uint(bId),
			Title:     colCell[2],
			Money:     money,
			Num:       num,
		}
		skillGoods[index-1] = skillGood
	}
	err = dao.NewSkillGoodsDao(c).CreateByList(skillGoods)
	if err != nil {
		code = e.ERROR
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Data:   "上传失败",
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// 初始化秒杀商品， 将MySQL信息存入Redis
func (service *SkillGoodsService) InitSkillGoods(c context.Context, uId uint) error {
	skillGoods, _ := dao.NewSkillGoodsDao(c).ListSkillGoods()
	r := cache.RedisClient

	// 存入Redis
	for i := range skillGoods {
		fmt.Println(*skillGoods[i])
		r.HSet("SK"+strconv.Itoa(int(skillGoods[i].Id)), "num", skillGoods[i].Num)
		r.HSet("SK"+strconv.Itoa(int(skillGoods[i].Id)), "money", skillGoods[i].Money)
	}
	return nil
}

func getUuid(gid string) string {
	codeLen := 8
	// 1.定义原始字符串
	rawStr := "hklajhdsfkhkdjhfak02304_"
	// 2.定义一个buf，将buf交给bytes往里面写数据
	buf := make([]byte, 0, codeLen)
	b := bytes.NewBuffer(buf)
	// 随机从中获取
	rand.Seed(time.Now().UnixNano())
	for rawStrLen := len(rawStr); codeLen > 0; codeLen-- {
		randNum := rand.Intn(rawStrLen)
		b.WriteByte(rawStr[randNum])
	}
	return b.String() + gid
}

// 加锁
func RedissonSecKillGoods(sk *model.SkillGood2MQ) error {
	p := strconv.Itoa(int(sk.ProductId))
	uuid := getUuid(p)
	_, err := cache.RedisClient.Del(p).Result()
	lockSuccess, err := cache.RedisClient.SetNX(p, uuid, time.Second*3).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail", err)
		return errors.New("get lock fail")
	} else {
		fmt.Println("get lock success")
	}
	_ = SendSecKillGoodsToMQ(sk)
	value, _ := cache.RedisClient.Get(p).Result()
	if value == uuid {
		_, err := cache.RedisClient.Del(p).Result()
		if err != nil {
			fmt.Println("unlock fail")
		} else {
			fmt.Println("unlock success")
		}
	}
	return nil
}

// 传送到MQ
func SendSecKillGoodsToMQ(sk *model.SkillGood2MQ) error {
	ch, err := mq.RabbitMQ.Channel()
	if err != nil {
		err = errors.New("rebbitMQ err:" + err.Error())
		return err
	}
	q, err := ch.QueueDeclare("skill_goods", true, false, false, false, nil)
	if err != nil {
		err = errors.New("rebbitMQ err:" + err.Error())
		return err
	}

	body, _ := json.Marshal(sk)
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})
	if err != nil {
		err = errors.New("rebbitMQ err:" + err.Error())
		return err
	}
	log.Printf("Sent %s", body)
	return nil
}

func (service *SkillGoodsService) SkillGoods(ctx context.Context, uId uint) serializer.Response {
	mo, _ := cache.RedisClient.HGet("SK"+strconv.Itoa(int(service.SkillGoodsId)), "money").Float64()
	sk := &model.SkillGood2MQ{
		ProductId:   service.ProductId,
		BossId:      service.BossId,
		UserId:      uId,
		AddressId:   service.AddressId,
		Key:         service.Key,
		Money:       mo,
		SkillGoodId: service.SkillGoodsId,
	}
	err := RedissonSecKillGoods(sk)
	if err != nil {
		return serializer.Response{}
	}
	return serializer.Response{}
}
