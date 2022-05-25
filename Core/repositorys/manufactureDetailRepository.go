package repositorys

import (
	"github.com/flabio/Core/entitys"
	"gorm.io/gorm"
)

type ManufactureDetailRepository interface {
	SetCreateManufactureDetail(manufactura entitys.ManufactureDetail) (entitys.ManufactureDetail, error)
	GetAllManufactureDetail() ([]entitys.ManufactureDetail, error)
	SetRemoveManufactureDetail(manufactura entitys.ManufactureDetail) (bool, error)
	GetFindManufacturegDetailById(Id uint) (entitys.ManufactureDetail, error)
}
type manufactureDetailConnection struct {
	connection *gorm.DB
}

func NewManufactureDetailRepository() ManufactureDetailRepository {
	var db *gorm.DB = entitys.DatabaseConnection()
	return &manufactureDetailConnection{
		connection: db,
	}
}

func (db *manufactureDetailConnection) SetCreateManufactureDetail(manufactura entitys.ManufactureDetail) (entitys.ManufactureDetail, error) {
	err := db.connection.Save(&manufactura).Error
	defer entitys.Closedb()
	return manufactura, err
}

func (db *manufactureDetailConnection) SetRemoveManufactureDetail(manufactura entitys.ManufactureDetail) (bool, error) {

	err := db.connection.Delete(&manufactura).Error
	defer entitys.Closedb()
	if err == nil {
		return true, err
	}
	return false, err
}

func (db *manufactureDetailConnection) GetFindManufacturegDetailById(Id uint) (entitys.ManufactureDetail, error) {

	var manufactura entitys.ManufactureDetail
	err := db.connection.Find(&manufactura, Id).Error
	defer entitys.Closedb()
	return manufactura, err
}

func (db *manufactureDetailConnection) GetAllManufactureDetail() ([]entitys.ManufactureDetail, error) {

	var manufacturas []entitys.ManufactureDetail
	err := db.connection.Find(&manufacturas).Error
	defer entitys.Closedb()
	return manufacturas, err
}
