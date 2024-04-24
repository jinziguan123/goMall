/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-16 10:22:13
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-16 10:40:12
 * @FilePath: /goMall/backend/serializer/notice.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import "goMall/backend/repository/database/model"

// Notice 公告序列化器
type Notice struct {
	ID        uint   `json:"id"`
	Text      string `json:"text"`
	CreatedAt int64  `json:"created_at"`
}

// BuildNotice 序列化公告
func BuildNotice(item *model.Notice) Notice {
	return Notice{
		ID:        item.ID,
		Text:      item.Text,
		CreatedAt: item.CreatedAt.Unix(),
	}
}
