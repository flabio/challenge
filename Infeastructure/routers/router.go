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
	//jwtService services.JWTService = services.NewJWTService()

	adminController Controllers.AdminController = Controllers.NewAdministradorController()
)

func NewRouter() {

	r := gin.New()
	docs.SwaggerInfo.BasePath = "/api/v1"
	// Set up CORS middleware options
	/*
		config := cors.Config{
			Origins:         "*",
			RequestHeaders:  "Authorization",
			Methods:         "GET, POST, PUT,DELETE",
			Credentials:     true,
			ValidateHeaders: false,
			MaxAge:          1 * time.Minute,
		}

		// Apply the middleware to the router (works on groups too)
		r.Use(cors.Middleware(config))
	*/
	//	rolRoutes := r.Group("api/v1/rol", middleware.AuthorizeJWT(jwtService))
	adminRoutes := r.Group("api/v1/administrador")
	{
		adminRoutes.GET("/test", adminController.Alls)
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
