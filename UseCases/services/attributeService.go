package services

import (
	"net/http"
	"strconv"

	"github.com/flabio/Core/entitys"
	"github.com/flabio/Core/repositorys"
	"github.com/flabio/UseCases/dto"
	"github.com/flabio/UseCases/utilities"
	"github.com/gin-gonic/gin"
	"github.com/mashingan/smapping"
)

type AttributeService interface {
	SetAttributeCreateService(context *gin.Context)
	SetAttributeUpdateService(context *gin.Context)
	SetAttributeRemoveService(context *gin.Context)
	GetAttributeAllService(context *gin.Context)
}

type attributeService struct {
	attributeRepository repositorys.AttributeRepository
}

func NewAttributeService() AttributeService {

	return &attributeService{
		attributeRepository: repositorys.NewAttributeRepository(),
	}
}

//service of create
func (service *attributeService) SetAttributeCreateService(context *gin.Context) {

	attribute := entitys.Attribute{}
	var attributeDto dto.AttributeDTO
	context.ShouldBind(&attributeDto)
	smapping.FillStruct(&attribute, smapping.MapFields(&attributeDto))

	data, err := service.attributeRepository.SetCreateAttribute(attribute)

	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildCreateResponse(http.StatusText(http.StatusCreated), data)
	context.JSON(http.StatusOK, res)
}

//service of update
func (service *attributeService) SetAttributeUpdateService(context *gin.Context) {

	attribute := entitys.Attribute{}
	var attributeDto dto.AttributeDTO

	context.ShouldBind(&attributeDto)
	err := smapping.FillStruct(&attribute, smapping.MapFields(&attributeDto))
	checkError(err)
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	findById, _ := service.attributeRepository.GetFindAttributeById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}

	data, err := service.attributeRepository.SetCreateAttribute(attribute)
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildUpdateResponse(http.StatusText(http.StatusCreated), data)
	context.JSON(http.StatusOK, res)

}

//service of all
func (service *attributeService) GetAttributeAllService(context *gin.Context) {

	var attributes, err = service.attributeRepository.GetAllAttribute()
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(http.StatusText(http.StatusOK), http.StatusText(http.StatusOK), attributes)
	context.JSON(http.StatusOK, res)
}

func (service *attributeService) SetAttributeRemoveService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	findById, _ := service.attributeRepository.GetFindAttributeById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	if err != nil {
		validadErrors(err, context)
		return
	}
	result, err := service.attributeRepository.SetRemoveAttribute(findById)

	if err != nil {
		validadErrorRemove(findById, context)
		return
	}
	if result {
		res := utilities.BuildDeteleteResponse(http.StatusText(http.StatusOK), findById)
		context.JSON(http.StatusOK, res)
	} else {
		res := utilities.BuildDeteleteResponse(http.StatusText(http.StatusBadRequest), findById)
		context.JSON(http.StatusOK, res)
	}

}
