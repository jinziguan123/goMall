/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 16:36:26
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-05 19:35:50
 * @FilePath: /goMall/backend/repository/cache/common.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package cache

import (
	"goMall/backend/config"
	"log"
	"strconv"

	"github.com/go-redis/redis"
)

// Redis单例
var RedisClient *redis.Client

// 初始化链接，防止循环导包
func InitCache() {
	Redis()
}

func Redis() {
	env := config.NewEnv()
	db, _ := strconv.ParseUint(env.RedisDbName, 10, 64)
	client := redis.NewClient(&redis.Options{
		Addr:     env.RedisAddr,
		Password: env.RedisPw,
		DB:       int(db),
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("can't connect to Redis", err)
	}
	RedisClient = client
}
