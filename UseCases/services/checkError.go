package services

import (
	"log"
	"net/http"

	"github.com/flabio/UseCases/utilities"

	"github.com/gin-gonic/gin"
)

/*
@param Err is of type error
*/
func checkError(Err error) bool {
	if Err != nil {
		log.Fatalf("Failed map %v", Err)
		return false

	}
	return true
}

//func of validation

/*
@param ErrDTO is of type error
*/
func validadErrors(ErrDTO error, context *gin.Context) {

	res := utilities.BuildErrorAllResponse(http.StatusText(http.StatusBadRequest), ErrDTO.Error())
	context.AbortWithStatusJSON(http.StatusBadRequest, res)
}

//validadRequiredMsg
/*
@param ErrDTO is of type error
*/
func validadRequiredMsg(Message string, context *gin.Context) {

	context.AbortWithStatusJSON(http.StatusBadRequest, utilities.BuildErrorAllResponseMessage(http.StatusText(http.StatusBadRequest), Message))
}
func validadBirdateRequiredMsg(message string, context *gin.Context) {

	res := utilities.BuildErrorAllResponseMessage(http.StatusText(http.StatusBadRequest), message)
	context.AbortWithStatusJSON(http.StatusBadRequest, res)
}
func validadErrorById(context *gin.Context) {
	res := utilities.BuildErrorByIdResponse(http.StatusText(http.StatusBadRequest))
	context.AbortWithStatusJSON(http.StatusBadRequest, res)
}

func validadErrorRemove(data interface{}, context *gin.Context) {
	response := utilities.BuildCanNotDeteleteResponse(http.StatusText(http.StatusBadRequest), data)
	context.JSON(http.StatusBadRequest, response)
}
