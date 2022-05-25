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

type ManufactureDetailService interface {
	SetManufactureCreateService(context *gin.Context)
	SetManufactureUpdateService(context *gin.Context)
	SetManufactureRemoveService(context *gin.Context)
	GetManufactureAllService(context *gin.Context)
}

type manufactureDetailService struct {
	manufactureDetailRepository repositorys.ManufactureDetailRepository
}

func NewManufactureDetailService() ManufactureDetailService {

	return &manufactureDetailService{
		manufactureDetailRepository: repositorys.NewManufactureDetailRepository(),
	}
}

//service of create
func (service *manufactureDetailService) SetManufactureCreateService(context *gin.Context) {

	manufacture := entitys.ManufactureDetail{}
	var manufactureDto dto.ManufactureDetailDTO
	context.ShouldBind(&manufactureDto)
	smapping.FillStruct(&manufacture, smapping.MapFields(&manufactureDto))

	data, err := service.manufactureDetailRepository.SetCreateManufactureDetail(manufacture)

	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildCreateResponse(http.StatusText(http.StatusCreated), data)
	context.JSON(http.StatusOK, res)
}

//service of update
func (service *manufactureDetailService) SetManufactureUpdateService(context *gin.Context) {

	manufacture := entitys.ManufactureDetail{}
	var manufactureDto dto.ManufactureDetailDTO

	context.ShouldBind(&manufactureDto)
	err := smapping.FillStruct(&manufacture, smapping.MapFields(&manufactureDto))
	checkError(err)
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	findById, _ := service.manufactureDetailRepository.GetFindManufacturegDetailById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}

	data, err := service.manufactureDetailRepository.SetCreateManufactureDetail(manufacture)
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildUpdateResponse(http.StatusText(http.StatusCreated), data)
	context.JSON(http.StatusOK, res)

}

//service of all
func (service *manufactureDetailService) GetManufactureAllService(context *gin.Context) {

	var manufactures, err = service.manufactureDetailRepository.GetAllManufactureDetail()
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(http.StatusText(http.StatusOK), http.StatusText(http.StatusOK), manufactures)
	context.JSON(http.StatusOK, res)
}

func (service *manufactureDetailService) SetManufactureRemoveService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	findById, _ := service.manufactureDetailRepository.GetFindManufacturegDetailById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	if err != nil {
		validadErrors(err, context)
		return
	}
	result, err := service.manufactureDetailRepository.SetRemoveManufactureDetail(findById)

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
