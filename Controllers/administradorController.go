package Controllers

import (
	"net/http"

	"github.com/flabio/Infeastructure/middleware"
	"github.com/flabio/UseCases/services"
	"github.com/flabio/UseCases/utilities"
	"github.com/gin-gonic/gin"
)

type AdminController interface {
	All(context *gin.Context)
	FindAdmin(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
}

type adminController struct {
	admin services.AdministradorService
	jwt   services.JWTService
}

func NewAdministradorController() AdminController {

	return &adminController{
		admin: services.NewAdministradorService(),
		jwt:   services.NewJWTService(),
	}
}

//get
func (c *adminController) All(context *gin.Context) {
	claim := middleware.GetAuthOwner(c.jwt, context)
	if len(claim.Criticidad) > 0 {
		c.admin.GetAllService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())

}

// get/:id
func (c *adminController) FindAdmin(context *gin.Context) {
	claim := middleware.GetAuthOwner(c.jwt, context)
	if claim.Criticidad == "high" || claim.Criticidad == "medium" {
		c.admin.GetFindByIdService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//Post
func (c *adminController) Create(context *gin.Context) {
	claim := middleware.GetAuthOwner(c.jwt, context)
	if claim.Criticidad == "high" || claim.Criticidad == "medium" {
		c.admin.SetCreateService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

//put
func (c *adminController) Update(context *gin.Context) {
	claim := middleware.GetAuthOwner(c.jwt, context)
	if claim.Criticidad == "high" || claim.Criticidad == "medium" {
		c.admin.SetUpdateService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}

// delete
func (c *adminController) Remove(context *gin.Context) {
	claim := middleware.GetAuthOwner(c.jwt, context)
	if claim.Criticidad == "high" {
		c.admin.SetRemoveService(context)
		return
	}
	context.JSON(http.StatusBadRequest, utilities.BuildDanedResponse())
}
