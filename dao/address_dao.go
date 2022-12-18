package dao

import (
	"context"
	"go-mall/model"
	_ "go-mall/model"
	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

// NewAddressDao 创建AddressDao 通过上下文的方式创建
func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{NewDBClient(ctx)}
}

// NewAddressDaoByDB 通过DB去创建AddressDao
func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{db}
}

func (dao *AddressDao) CreateAddress(address *model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(&address).Error
}

func (dao *AddressDao) GetAddressById(id uint) (address *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("id = ?", id).First(&address).Error
	return
}

func (dao *AddressDao) ListAddressByUseId(uId uint) (address []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where("user_id", uId).Find(&address).Error
	return
}

func (dao *AddressDao) UpdateAddress(aId uint, address *model.Address) error {
	return dao.DB.Where(&model.Address{}).Where("id = ?", aId).Updates(&address).Error
}

func (dao *AddressDao) DeleteAddressByAddressId(aId, uId uint) error {
	return dao.DB.Where(&model.Address{}).Where("id = ? AND user_id = ?", aId, uId).Delete(&model.Address{}).Error
}
