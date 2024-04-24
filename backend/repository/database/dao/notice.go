/*
 * @Author: Ziguan Jin 18917950960@163.com
 * @Date: 2024-04-07 11:14:58
 * @LastEditors: Ziguan Jin 18917950960@163.com
 * @LastEditTime: 2024-04-16 10:32:47
 * @FilePath: /goMall/backend/repository/database/dao/notice.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package dao

import (
	"context"

	"gorm.io/gorm"

	"goMall/backend/repository/database/model"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// GetNoticeById 通过id获取notice
func (dao *NoticeDao) GetNoticeById(id uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}

// CreateNotice 创建notice
func (dao *NoticeDao) CreateNotice(notice *model.Notice) error {
	return dao.DB.Model(&model.Notice{}).Create(&notice).Error
}

// 获取一条notice
func (dao *NoticeDao) GetOneNotice() (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).First(&notice).Error
	return
}
