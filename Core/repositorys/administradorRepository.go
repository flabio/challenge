package repositorys

import (
	"github.com/flabio/Core/entitys"
	"gorm.io/gorm"
)

type AdministradorRepository interface {
	SetCreateAdministrador(admin entitys.Administrador) (entitys.Administrador, error)
	GetAllAdministrador() ([]entitys.Administrador, error)
	SetRemoveAdministrador(admin entitys.Administrador) (bool, error)
	GetFindAdministradorById(Id uint) (entitys.Administrador, error)
	VerifyCredential(Name string, Owner string) interface{}
}
type administradorConnection struct {
	connection *gorm.DB
}

func NewAdministradorRepository() AdministradorRepository {
	var db *gorm.DB = entitys.DatabaseConnection()
	return &administradorConnection{
		connection: db,
	}
}

/*
@param admin, is a struct of Administrador
*/
func (db *administradorConnection) SetCreateAdministrador(admin entitys.Administrador) (entitys.Administrador, error) {
	err := db.connection.Save(&admin).Error
	defer entitys.Closedb()
	return admin, err
}

/*
@param admin, is a struct of Administrador
*/
func (db *administradorConnection) SetRemoveAdministrador(admin entitys.Administrador) (bool, error) {

	err := db.connection.Delete(&admin).Error
	defer entitys.Closedb()
	if err == nil {
		return true, err
	}
	return false, err
}

/*
@param Id, is a uint of Administrador
*/
func (db *administradorConnection) GetFindAdministradorById(Id uint) (entitys.Administrador, error) {

	var admin entitys.Administrador
	err := db.connection.Find(&admin, Id).Error
	defer entitys.Closedb()
	return admin, err
}

func (db *administradorConnection) GetAllAdministrador() ([]entitys.Administrador, error) {

	var admins []entitys.Administrador
	err := db.connection.Find(&admins).Error
	defer entitys.Closedb()
	return admins, err
}

/*
@param Name, de tipo string
@param Owner,de tipo string
*/
func (db *administradorConnection) VerifyCredential(Name string, Owner string) interface{} {
	var user entitys.Administrador
	err := db.connection.Where("name = ?", Name).Where("owner = ?", Owner).Find(&user).Error
	defer entitys.Closedb()

	if err == nil {
		return user
	}
	return nil
}
