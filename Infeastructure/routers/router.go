package routers

import (
	"github.com/flabio/Controllers"
	docs "github.com/flabio/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//controller
var (
	shippingController    Controllers.ShippingController    = Controllers.NewShippingController()
	manufacturaController Controllers.ManufactureController = Controllers.NewManufactureController()
	attributeController   Controllers.AttributeController   = Controllers.NewAttributeController()
	producController      Controllers.ProductController     = Controllers.NewProductController()
)

func NewRouter() {

	r := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"

	shippingRoutes := r.Group("api/v1/shipping")
	{

		shippingRoutes.GET("/", shippingController.All)
		shippingRoutes.POST("/create", shippingController.Create)
		shippingRoutes.PUT("/update/:id", shippingController.Update)
		shippingRoutes.DELETE("/:id", shippingController.Remove)
	}
	manufacturaRoutes := r.Group("api/v1/manufactura")
	{

		manufacturaRoutes.GET("/", manufacturaController.All)
		manufacturaRoutes.POST("/create", manufacturaController.Create)
		manufacturaRoutes.PUT("/update/:id", manufacturaController.Update)
		manufacturaRoutes.DELETE("/:id", manufacturaController.Remove)
	}
	attributeRoutes := r.Group("api/v1/attribute")
	{

		attributeRoutes.GET("/", attributeController.All)
		attributeRoutes.POST("/create", attributeController.Create)
		attributeRoutes.PUT("/update/:id", attributeController.Update)
		attributeRoutes.DELETE("/:id", attributeController.Remove)
	}
	productRoutes := r.Group("api/v1/product")
	{

		productRoutes.GET("/", producController.All)
		productRoutes.POST("/create", producController.Create)
		productRoutes.PUT("/update/:id", producController.Update)
		productRoutes.DELETE("/:id", producController.Remove)
	}
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	r.Run(":8080")

	return
}
