package repositorys

import (
	"github.com/flabio/Core/entitys"
	"gorm.io/gorm"
)

type AttributeRepository interface {
	SetCreateAttribute(attribute entitys.Attribute) (entitys.Attribute, error)
	GetAllAttribute() ([]entitys.Attribute, error)
	SetRemoveAttribute(attribute entitys.Attribute) (bool, error)
	GetFindAttributeById(Id uint) (entitys.Attribute, error)
}
type attributeConnection struct {
	connection *gorm.DB
}

func NewAttributeRepository() AttributeRepository {
	var db *gorm.DB = entitys.DatabaseConnection()
	return &attributeConnection{
		connection: db,
	}
}

func (db *attributeConnection) SetCreateAttribute(attribute entitys.Attribute) (entitys.Attribute, error) {
	err := db.connection.Save(&attribute).Error
	defer entitys.Closedb()
	return attribute, err
}

func (db *attributeConnection) SetRemoveAttribute(attribute entitys.Attribute) (bool, error) {

	err := db.connection.Delete(&attribute).Error
	defer entitys.Closedb()
	if err == nil {
		return true, err
	}
	return false, err
}

func (db *attributeConnection) GetFindAttributeById(Id uint) (entitys.Attribute, error) {

	var attribute entitys.Attribute
	err := db.connection.Find(&attribute, Id).Error
	defer entitys.Closedb()
	return attribute, err
}

func (db *attributeConnection) GetAllAttribute() ([]entitys.Attribute, error) {

	var attributes []entitys.Attribute
	err := db.connection.Find(&attributes).Error
	defer entitys.Closedb()
	return attributes, err
}
