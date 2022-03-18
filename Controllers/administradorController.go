package Controllers

import (
	"log"

	"github.com/flabio/UseCases/services"
	"github.com/gin-gonic/gin"
)

//RolController is a..

type AdminController interface {
	Alls(context *gin.Context)
	All(context *gin.Context)
	FindAdmin(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
}

type adminController struct {
	admin services.AdministradorService
	//jwt services.JWTService
}

func NewAdministradorController() AdminController {

	return &adminController{
		admin: services.NewAdministradorService(),
		//jwt: services.NewJWTService(),
	}
}

//GET /rols
// get list of rol
func (c *adminController) All(context *gin.Context) {
	//claim := middleware.GetRol(c.jwt, context)
	/*if claim.Rol == 1 {
		c.rol.GetAllService(context)
		return
	}*/
	c.admin.GetAllService(context)
	//context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//GET /rols
// get list of rol
func (c *adminController) Alls(context *gin.Context) {
	//claim := middleware.GetRol(c.jwt, context)
	/*if claim.Rol == 1 {
		c.rol.GetAllService(context)
		return
	}*/
	log.Println("Error : something terrible happen -> ")

	//context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

// get
func (c *adminController) FindAdmin(context *gin.Context) {
	/*
		claim := middleware.GetRol(c.jwt, context)
		if claim.Rol == 1 {
			fmt.Println("aqui controller")
			c.rol.GetFindByIdService(context)
			return
		}*/
	c.admin.GetFindByIdService(context)
	//context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}

//create rol metho post
func (c *adminController) Create(context *gin.Context) {
	/*
		claim := middleware.GetRol(c.jwt, context)
		if claim.Rol == 1 {
			c.rol.SetCreateService(context)
			return
		}*/
	c.admin.SetCreateService(context)
	//context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}

//update rol method put
func (c *adminController) Update(context *gin.Context) {
	/*claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.rol.SetUpdateService(context)
		return
	}*/
	c.admin.SetUpdateService(context)
	//context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}

// delete rol
func (c *adminController) Remove(context *gin.Context) {
	/*claim := middleware.GetRol(c.jwt, context)
	if claim.Rol == 1 {
		c.rol.SetRemoveService(context)
		return
	}*/
	c.admin.SetRemoveService(context)
	//context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}
