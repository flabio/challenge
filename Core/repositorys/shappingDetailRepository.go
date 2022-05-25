package repositorys

import (
	"github.com/flabio/Core/entitys"
	"gorm.io/gorm"
)

type ShippingDetailRepository interface {
	SetCreateShippingDetail(shipping entitys.ShippingDetail) (entitys.ShippingDetail, error)
	GetAllShippingDetail() ([]entitys.ShippingDetail, error)
	SetRemoveShippingDetail(shipping entitys.ShippingDetail) (bool, error)
	GetFindShippingDetailById(Id uint) (entitys.ShippingDetail, error)
}
type shippingDetailConnection struct {
	connection *gorm.DB
}

func NewShippingDetailRepository() ShippingDetailRepository {
	var db *gorm.DB = entitys.DatabaseConnection()
	return &shippingDetailConnection{
		connection: db,
	}
}

func (db *shippingDetailConnection) SetCreateShippingDetail(shipping entitys.ShippingDetail) (entitys.ShippingDetail, error) {
	err := db.connection.Save(&shipping).Error
	defer entitys.Closedb()
	return shipping, err
}

func (db *shippingDetailConnection) SetRemoveShippingDetail(shipping entitys.ShippingDetail) (bool, error) {

	err := db.connection.Delete(&shipping).Error
	defer entitys.Closedb()
	if err == nil {
		return true, err
	}
	return false, err
}

func (db *shippingDetailConnection) GetFindShippingDetailById(Id uint) (entitys.ShippingDetail, error) {

	var shipping entitys.ShippingDetail
	err := db.connection.Find(&shipping, Id).Error
	defer entitys.Closedb()
	return shipping, err
}

func (db *shippingDetailConnection) GetAllShippingDetail() ([]entitys.ShippingDetail, error) {

	var shippings []entitys.ShippingDetail
	err := db.connection.Find(&shippings).Error
	defer entitys.Closedb()
	return shippings, err
}
