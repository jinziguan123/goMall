/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 20:51:14
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-10 09:51:26
 * @FilePath: /goMall/backend/repository/database/dao/migration.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"fmt"
	model2 "goMall/backend/repository/database/model"
	"os"
)

func Migration() {
	// 数据迁移
	err := _db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(
		&model2.User{},
		&model2.Product{},
		&model2.Carousel{},
		&model2.Category{},
		&model2.Favorite{},
		&model2.ProductImg{},
		&model2.Order{},
		&model2.Cart{},
		&model2.Admin{},
		&model2.Address{},
		&model2.Notice{},
		&model2.SkillGoods{},
	)
	if err != nil {
		fmt.Println("register table fail")
		os.Exit(0)
	}
	fmt.Println("register table success")
}
