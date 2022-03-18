package routers

import (
	"github.com/flabio/Controllers"
	"github.com/flabio/Infeastructure/middleware"
	"github.com/flabio/UseCases/services"
	docs "github.com/flabio/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//controller
var (
	jwtService services.JWTService = services.NewJWTService()

	authController Controllers.AuthController = Controllers.NewAuthController()

	adminController Controllers.AdminController = Controllers.NewAdministradorController()
)

func NewRouter() {

	r := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"

	authRoutes := r.Group("api/v1/login")
	{
		authRoutes.POST("/", authController.Login)

	}

	adminRoutes := r.Group("api/v1/administrador", middleware.AuthorizeJWT(jwtService))
	{

		adminRoutes.GET("/", adminController.All)
		adminRoutes.POST("/create", adminController.Create)
		adminRoutes.PUT("/update/:id", adminController.Update)
		adminRoutes.DELETE("/:id", adminController.Remove)
		adminRoutes.GET("/:id", adminController.FindAdmin)
	}

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	r.Run(":8080")

	return
}
