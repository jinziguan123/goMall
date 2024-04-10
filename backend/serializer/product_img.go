/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 16:44:22
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-10 00:36:04
 * @FilePath: /goMall/backend/serializer/product_img.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package serializer

import (
	"goMall/backend/config"
	"goMall/backend/consts"
	"goMall/backend/repository/database/model"
)

type ProductImg struct {
	ProductID uint   `json:"product_id" form:"product_id"`
	ImgPath   string `json:"img_path" form:"img_path"`
}

func BuildProductImg(item *model.ProductImg) ProductImg {
	pImg := ProductImg{
		ProductID: item.ProductID,
		ImgPath:   config.PhotoHost + config.HttpPort + config.ProductPhotoPath + item.ImgPath,
	}
	if config.UploadModel == consts.UploadModelOss {
		pImg.ImgPath = item.ImgPath
	}

	return pImg
}

func BuildProductImgs(items []*model.ProductImg) (productImgs []ProductImg) {
	for _, item := range items {
		product := BuildProductImg(item)
		productImgs = append(productImgs, product)
	}
	return productImgs
}
