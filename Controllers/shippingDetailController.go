package Controllers

import (
	"github.com/flabio/UseCases/services"
	"github.com/gin-gonic/gin"
)

type ShippingController interface {
	All(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
}

type shippingController struct {
	shipping services.ShippingDetailService
}

func NewShippingController() ShippingController {

	return &shippingController{
		shipping: services.NewShippingDetailService(),
	}
}

//get
func (c *shippingController) All(context *gin.Context) {
	c.shipping.GetshippingAllService(context)

}

//Post
func (c *shippingController) Create(context *gin.Context) {
	c.shipping.SetshippingCreateService(context)
}

//put
func (c *shippingController) Update(context *gin.Context) {
	c.shipping.SetshippingUpdateService(context)
}

// delete
func (c *shippingController) Remove(context *gin.Context) {
	c.shipping.SetshippingRemoveService(context)
}
