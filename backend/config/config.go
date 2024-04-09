/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 16:39:07
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 15:21:44
 * @FilePath: /goMall/backend/config/config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppMode     string `mapstructure:"APP_MODE"`
	HttpPort    string `mapstructure:"HTTP_PORT"`
	UploadModel string `mapstructure:"UPLOAD_MODEL"`

	MySQLHost     string `mapstructure:"MYSQL_HOST"`
	MySQLPort     string `mapstructure:"MYSQL_PORT"`
	MySQLUserName string `mapstructure:"MYSQL_USER_NAME"`
	MySQLPassword string `mapstructure:"MYSQL_PASSWORD"`
	MySQLDbName   string `mapstructure:"MYSQL_DB_NAME"`

	RedisDb     string `mapstructure:"REDIS_DB"`
	RedisAddr   string `mapstructure:"REDIS_ADDR"`
	RedisPw     string `mapstructure:"REDIS_PASSWORD"`
	RedisDbName string `mapstructure:"REDIS_DB_NAME"`

	AccessKey    string `mapstructure:"ACCESS_KEY"`
	SecretKey    string `mapstructure:"SECRET_KEY"`
	Bucket       string `mapstructure:"BUCKET"`
	HuaweiServer string `mapstructure:"HUAWEI_SERVER"`

	ValidEmail string `mapstructure:"VALID_EMAIL"`
	SmtpHost   string `mapstructure:"SMTP_HOST"`
	SmtpEmail  string `mapstructure:"SMTP_EMAIL"`
	SmtpPass   string `mapstructure:"SMTP_PASS"`

	PhotoHost        string `mapstructure:"PHOTO_HOST"`
	ProductPhotoHost string `mapstructure:"PRODUCT_PHOTO_HOST"`
	AvatarPath       string `mapstructure:"AVATAR_PATH"`

	EsHost  string
	EsPort  string
	EsIndex string

	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassWord string
	RabbitMQHost     string
	RabbitMQPort     string
}

func NewEnv() *Env {
	env := Env{}

	vip := viper.New()
	vip.SetConfigFile("config.yaml")
	vip.SetConfigType("yaml")

	err := vip.ReadInConfig()
	if err != nil {
		log.Fatal("can't find the config file", err)
	}

	err = vip.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded", err)
	}

	return &env
}
