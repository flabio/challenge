package Controllers

import (
	"github.com/flabio/UseCases/services"
	"github.com/gin-gonic/gin"
)

type ProductController interface {
	All(context *gin.Context)
	Create(context *gin.Context)
	Update(context *gin.Context)
	Remove(context *gin.Context)
}

type productController struct {
	product services.ProductService
}

func NewProductController() ProductController {

	return &productController{
		product: services.NewProductService(),
	}
}

//get
func (c *productController) All(context *gin.Context) {
	c.product.GetProductAllService(context)

}

//Post
func (c *productController) Create(context *gin.Context) {
	c.product.SetProductCreateService(context)
}

//put
func (c *productController) Update(context *gin.Context) {
	c.product.SetProductUpdateService(context)
}

// delete
func (c *productController) Remove(context *gin.Context) {
	c.product.SetProductRemoveService(context)
}
