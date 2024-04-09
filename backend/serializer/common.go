/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 16:16:57
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-07 16:17:00
 * @FilePath: /goMall/backend/serializer/common.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

// DataList 带有总数的Data结构
type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

// TokenData 带有token的Data结构
type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

// TrackedErrorResponse 有追踪信息的错误反应
type TrackedErrorResponse struct {
	Response
	TrackID string `json:"track_id"`
}

// BuildListResponse 带有总数的列表构建器
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Item:  items,
			Total: total,
		},
		Msg: "ok",
	}
}
