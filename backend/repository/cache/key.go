/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-05 19:38:28
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-05 19:41:57
 * @FilePath: /goMall/backend/repository/cache/key.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package cache

import (
	"fmt"
	"strconv"
)

const (
	// 每日排名
	RankKey = "rank"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("view: product: %s", strconv.Itoa(int(id)))
}
