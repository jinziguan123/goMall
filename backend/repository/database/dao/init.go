/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 20:08:26
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-05 20:54:39
 * @FilePath: /goMall/backend/repository/database/dao/init.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"
	"goMall/backend/config"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var (
	_db *gorm.DB
)

func InitMySQL() {
	env := config.NewEnv()
	pathRead := strings.Join([]string{env.MySQLUserName, ":", env.MySQLPassword, "@tcp(", env.MySQLHost, ")/", env.MySQLDbName, "?charset=utf8&parseTime=True&loc=Local"}, "")
	pathWrite := strings.Join([]string{env.MySQLUserName, ":", env.MySQLPassword, "@tcp(", env.MySQLHost, ")/", env.MySQLDbName, "?charset=utf8&parseTime=True&loc=Local"}, "")

	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       pathRead, // DSN data source name
		DefaultStringSize:         256,      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,    // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 设置连接池
	sqlDB.SetMaxOpenConns(100) // 打开
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(pathRead)},
		Replicas: []gorm.Dialector{mysql.Open(pathWrite), mysql.Open(pathWrite)},
		Policy:   dbresolver.RandomPolicy{},
	}))
	Migration()
}

func NewDBClient(c context.Context) *gorm.DB {
	db := _db
	return db.WithContext(c)
}
