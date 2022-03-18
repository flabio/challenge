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

var errChanAdministrador = make(chan error, 1)

/*
@param admin, is a struct of Administrador
*/
func (db *administradorConnection) SetCreateAdministrador(admin entitys.Administrador) (entitys.Administrador, error) {

	go func() {
		err := db.connection.Save(&admin).Error
		defer entitys.Closedb()
		errChanAdministrador <- err
	}()
	err := <-errChanAdministrador

	return admin, err
}

/*
@param admin, is a struct of Administrador
*/
func (db *administradorConnection) SetRemoveAdministrador(admin entitys.Administrador) (bool, error) {

	go func() {
		err := db.connection.Delete(&admin).Error
		defer entitys.Closedb()
		errChanAdministrador <- err
	}()
	err := <-errChanAdministrador
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
	go func() {
		err := db.connection.Find(&admin, Id).Error
		defer entitys.Closedb()
		errChanAdministrador <- err
	}()
	err := <-errChanAdministrador
	return admin, err
}

func (db *administradorConnection) GetAllAdministrador() ([]entitys.Administrador, error) {

	var admins []entitys.Administrador
	go func() {
		err := db.connection.Find(&admins).Error
		defer entitys.Closedb()
		errChanAdministrador <- err
	}()
	err := <-errChanAdministrador
	return admins, err
}
