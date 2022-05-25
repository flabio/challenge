package repositorys

import (
	"github.com/flabio/Core/entitys"
	"gorm.io/gorm"
)

type ProductRepository interface {
	SetCreateProduct(product entitys.Product) (entitys.Product, error)
	GetAllProduct() ([]entitys.Product, error)
	SetRemoveProduct(product entitys.Product) (bool, error)
	GetFindProductById(Id uint) (entitys.Product, error)
}
type productConnection struct {
	connection *gorm.DB
}

func NewProductRepository() ProductRepository {
	var db *gorm.DB = entitys.DatabaseConnection()
	return &productConnection{
		connection: db,
	}
}

func (db *productConnection) SetCreateProduct(product entitys.Product) (entitys.Product, error) {
	err := db.connection.Save(&product).Error
	defer entitys.Closedb()
	return product, err
}

func (db *productConnection) SetRemoveProduct(product entitys.Product) (bool, error) {

	err := db.connection.Delete(&product).Error
	defer entitys.Closedb()
	if err == nil {
		return true, err
	}
	return false, err
}

func (db *productConnection) GetFindProductById(Id uint) (entitys.Product, error) {

	var product entitys.Product
	err := db.connection.Find(&product, Id).Error
	defer entitys.Closedb()
	return product, err
}

func (db *productConnection) GetAllProduct() ([]entitys.Product, error) {

	var products []entitys.Product
	err := db.connection.Preload("ManufactureDetail").Preload("ShippingDetail").Preload("Attribute").Find(&products).Error
	defer entitys.Closedb()
	return products, err
}
