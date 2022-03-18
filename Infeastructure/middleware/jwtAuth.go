package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/flabio/UseCases/services"
	"github.com/flabio/UseCases/utilities"
	"github.com/gin-gonic/gin"
)

type GetVariableSession interface {
}
type VariableSession struct {
	Id         uint
	Name       string
	Owner      string
	Criticidad string
}

func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := utilities.BuildErrorResponse("No se pudo procesar la solicitud", "No se encontró ningún token", nil)
			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token != nil {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[id]: ", claims["id"])
			log.Println("Claim[issuer] :", claims["issuer"])
			log.Println("Claim[Name] :", claims["name"])
			log.Println("Claim[owner] :", claims["owner"])
			log.Println("Claim[criticidad] :", claims["criticidad"])
			log.Println("claims[exp]", claims["exp"])
		} else {
			response := utilities.BuildErrorResponse("El token no es válido", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}

	}
}

func GetAuthOwner(jwtService services.JWTService, context *gin.Context) VariableSession {
	authHeader := context.GetHeader("Authorization")
	token, _ := jwtService.ValidateToken(authHeader)
	claims := token.Claims.(jwt.MapClaims)
	id, _ := strconv.ParseUint(fmt.Sprintf("%v", claims["id"]), 0, 0)
	name := fmt.Sprintf("%v", claims["name"])
	owner := fmt.Sprintf("%v", claims["owner"])
	criticidad := fmt.Sprintf("%v", claims["criticidad"])
	u := VariableSession{
		Id:         uint(id),
		Name:       name,
		Owner:      owner,
		Criticidad: criticidad,
	}

	return u
}
