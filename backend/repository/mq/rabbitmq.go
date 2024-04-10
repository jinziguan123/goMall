/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-08 14:32:50
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-10 00:34:18
 * @FilePath: /goMall/backend/repository/mq/rabbitmq.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mq

import (
	"goMall/backend/config"
	"strings"

	"github.com/streadway/amqp"
)

// RabbitMQ rabbitMQ链接单例
var RabbitMQ *amqp.Connection

// InitRabbitMQ 在中间件中初始化rabbitMQ链接
func InitRabbitMQ() {
	pathRabbitMQ := strings.Join([]string{config.RabbitMQ, "://", config.RabbitMQUser, ":", config.RabbitMQPassWord, "@", config.RabbitMQHost, ":", config.RabbitMQPort, "/"}, "")
	conn, err := amqp.Dial(pathRabbitMQ)
	if err != nil {
		panic(err)
	}
	RabbitMQ = conn
}
