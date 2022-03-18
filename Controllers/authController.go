package Controllers

import (
	"net/http"
	"strconv"

	"github.com/flabio/Core/entitys"
	"github.com/flabio/UseCases/dto"
	"github.com/flabio/UseCases/services"
	"github.com/flabio/UseCases/utilities"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
	jwtService  services.JWTService
}

func NewAuthController() AuthController {
	jwtService := services.NewJWTService()
	authService := services.NewAuthService()
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {

	var loginDTO dto.LoginDTO

	ctx.ShouldBind(&loginDTO)

	if len(loginDTO.Name) == 0 && len(loginDTO.Owner) == 0 {
		response := utilities.BuildNameOwnerIncorrectResponse()
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}
	authResult := c.authService.VerifyCredential(loginDTO.Name, loginDTO.Owner)

	if v, ok := authResult.(entitys.Administrador); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(uint64(v.Id), 10), v.Criticidad)
		v.Token = generatedToken
		response := utilities.BuildResponse("200", "OK!", v)
		ctx.JSON(http.StatusOK, response)
		return
	}

	response := utilities.BuildErrorResponse("por favor verifique nuevamente su credencial", "Credencial Invalidad", utilities.EmptObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
}
