package Controllers

import (
	"github.com/flabio/UseCases/services"
	"github.com/gin-gonic/gin"
)

type ManufactureController interface {
	All(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
}

type manufactureController struct {
	manufactura services.ManufactureDetailService
}

func NewManufactureController() ManufactureController {

	return &manufactureController{
		manufactura: services.NewManufactureDetailService(),
	}
}

//get
func (c *manufactureController) All(context *gin.Context) {
	c.manufactura.GetManufactureAllService(context)

}

//Post
func (c *manufactureController) Create(context *gin.Context) {
	c.manufactura.SetManufactureCreateService(context)
}

//put
func (c *manufactureController) Update(context *gin.Context) {
	c.manufactura.SetManufactureUpdateService(context)
}

// delete
func (c *manufactureController) Remove(context *gin.Context) {
	c.manufactura.SetManufactureRemoveService(context)
}
