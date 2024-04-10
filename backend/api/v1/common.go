/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-10 00:47:02
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-10 09:05:59
 * @FilePath: /goMall/backend/api/v1/common.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"encoding/json"
	"fmt"
	"goMall/backend/config"
	"goMall/backend/consts"
	"goMall/backend/serializer"

	"github.com/go-playground/validator/v10"
)

func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := config.T(fmt.Sprintf("Field.%s", e.Field()))
			tag := config.T(fmt.Sprintf("Tag.Valid.%s", e.Tag()))
			return serializer.Response{
				Status: consts.IlleageRequest,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: consts.IlleageRequest,
			Msg:    "JSON类型不匹配",
			Error:  fmt.Sprint(err),
		}
	}
	return serializer.Response{
		Status: consts.IlleageRequest,
		Msg:    "参数错误",
		Error:  fmt.Sprint(err),
	}
}
