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

type ShippingDetailService interface {
	SetshippingCreateService(context *gin.Context)
	SetshippingUpdateService(context *gin.Context)

	SetshippingRemoveService(context *gin.Context)
	GetshippingAllService(context *gin.Context)
}

type shippingDetailService struct {
	shippingDetailRepository repositorys.ShippingDetailRepository
}

func NewShippingDetailService() ShippingDetailService {

	return &shippingDetailService{
		shippingDetailRepository: repositorys.NewShippingDetailRepository(),
	}
}

//service of create
func (service *shippingDetailService) SetshippingCreateService(context *gin.Context) {

	shipping := entitys.ShippingDetail{}
	var shippingDto dto.ShippingDetailDTO
	context.ShouldBind(&shippingDto)
	smapping.FillStruct(&shipping, smapping.MapFields(&shippingDto))

	data, err := service.shippingDetailRepository.SetCreateShippingDetail(shipping)

	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildCreateResponse(http.StatusText(http.StatusCreated), data)
	context.JSON(http.StatusOK, res)
}

//service of update
func (service *shippingDetailService) SetshippingUpdateService(context *gin.Context) {

	shipping := entitys.ShippingDetail{}
	var shippingDto dto.ShippingDetailDTO

	context.ShouldBind(&shippingDto)
	err := smapping.FillStruct(&shipping, smapping.MapFields(&shippingDto))
	checkError(err)
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	findById, _ := service.shippingDetailRepository.GetFindShippingDetailById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}

	data, err := service.shippingDetailRepository.SetCreateShippingDetail(shipping)
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildUpdateResponse(http.StatusText(http.StatusCreated), data)
	context.JSON(http.StatusOK, res)

}

//service of all
func (service *shippingDetailService) GetshippingAllService(context *gin.Context) {

	var shipping, err = service.shippingDetailRepository.GetAllShippingDetail()
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(http.StatusText(http.StatusOK), http.StatusText(http.StatusOK), shipping)
	context.JSON(http.StatusOK, res)
}

func (service *shippingDetailService) SetshippingRemoveService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	findById, _ := service.shippingDetailRepository.GetFindShippingDetailById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	if err != nil {
		validadErrors(err, context)
		return
	}
	result, err := service.shippingDetailRepository.SetRemoveShippingDetail(findById)

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
