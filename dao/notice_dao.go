package dao

import (
	"context"
	"go-mall/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

// NewNoticeDao 创建UserDao 通过上下文的方式创建
func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{NewDBClient(ctx)}
}

// NewNoticeDaoByDB 通过DB去创建UserDao
func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

// GetNoticeById 根据id查询通知notice模板信息
func (dao *NoticeDao) GetNoticeById(id uint) (notice *model.Notice, err error) {
	err = dao.DB.Model(&model.Notice{}).Where("id=?", id).First(&notice).Error
	return
}
