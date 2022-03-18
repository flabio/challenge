package utilities

import (
	"strings"
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
func BuildErrorResponse(status string, message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  status,
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
func BuildDanedResponse(status string) Response {
	res := Response{
		Status:  status,
		Message: "permission denied",
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
		Message: "Create successfully",
		Error:   nil,
		Data:    data,
	}
	return res
}

//BuildUpdateResponse method is to inject dasta value to dynamic success response
func BuildUpdateResponse(status string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: "Update successfully",
		Error:   nil,
		Data:    data,
	}
	return res
}

//BuildErrorResponse method is to inject dasta value to dynamic failed response
func BuildDeteleteResponse(status string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: "It was successfully removed",
		Data:    data,
	}
	return res
}

//"Data not found", "No data with given id",
//BuildErrorResponse method is to inject dasta value to dynamic failed response
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
