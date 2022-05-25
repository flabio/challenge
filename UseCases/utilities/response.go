package utilities

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//Response is used for static shape json return
type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

//EmptyObj object is used when data doesnt want to be null on json
type EmptObj struct{}

//BuildResponse method is to inject dasta value to dynamic success response
func BuildResponse(status string, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  "400",
		Message: message,
		Error:   splittedError,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildNotFoudResponse(status string) Response {
	res := Response{
		Status:  status,
		Message: "not found",
	}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildDanedResponse() Response {
	res := Response{
		Status:  "300",
		Message: "Permiso denegado",
		Error:   nil,
	}

	return res
}

//BuildErrorAllResponse method is to inject dasta value to dynamic failed response
func BuildErrorAllResponse(status string, err string) Response {
	res := Response{
		Status:  status,
		Message: err,
	}
	return res
}

//BuildErrorAllResponse method is to inject dasta value to dynamic failed response
func BuildErrorAllResponseMessage(status string, message string) Response {

	res := Response{
		Status:  status,
		Message: message,
	}
	return res
}

//BuildCreateResponse method is to inject dasta value to dynamic success response
func BuildCreateResponse(status string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: "Se creo éxitosamente",
		Error:   nil,
		Data:    data,
	}
	return res
}

//BuildUpdateResponse method is to inject dasta value to dynamic success response
func BuildUpdateResponse(status string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: "Se modifico éxitosamente",
		Error:   nil,
		Data:    data,
	}
	return res
}

func BuildDeteleteResponse(status string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: "Fue eliminado con éxito",
		Data:    data,
	}
	return res
}

func BuildErrorByIdResponse(status string) Response {

	res := Response{
		Status:  status,
		Message: "Data not found",
		Error:   "No data with given id"}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildCanNotDeteleteResponse(status string, data interface{}) Response {

	res := Response{
		Status:  status,
		Message: "The record cannot be deleted",
		Data:    data,
	}
	return res
}
func BuildNameOwnerIncorrectResponse() Response {
	res := Response{
		Status:  "400",
		Message: "El Owner es incorrecto",
		Error:   nil,
	}
	return res
}

func Pagination(c *gin.Context, limit int) (int, int) {
	p := c.Query("page")

	if p == "" {
		return 1, 0
	}
	page, _ := strconv.Atoi(c.Query("page"))
	if page < 1 {
		return 1, 0
	}

	begin := (limit * page) - limit
	return page, begin
}
