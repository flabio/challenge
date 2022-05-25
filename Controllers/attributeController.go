package Controllers

import (
	"github.com/flabio/UseCases/services"
	"github.com/gin-gonic/gin"
)

type AttributeController interface {
	All(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
}

type attributeController struct {
	attribute services.AttributeService
}

func NewAttributeController() AttributeController {

	return &attributeController{
		attribute: services.NewAttributeService(),
	}
}

//get
func (c *attributeController) All(context *gin.Context) {
	c.attribute.GetAttributeAllService(context)

}

//Post
func (c *attributeController) Create(context *gin.Context) {
	c.attribute.SetAttributeCreateService(context)
}

//put
func (c *attributeController) Update(context *gin.Context) {
	c.attribute.SetAttributeUpdateService(context)
}

// delete
func (c *attributeController) Remove(context *gin.Context) {
	c.attribute.SetAttributeRemoveService(context)
}
