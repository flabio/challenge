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

type ProductService interface {
	SetProductCreateService(context *gin.Context)
	SetProductUpdateService(context *gin.Context)
	SetProductRemoveService(context *gin.Context)
	GetProductAllService(context *gin.Context)
}

type productService struct {
	productRepository repositorys.ProductRepository
}

func NewProductService() ProductService {

	return &productService{
		productRepository: repositorys.NewProductRepository(),
	}
}

//service of create
func (service *productService) SetProductCreateService(context *gin.Context) {

	product := entitys.Product{}
	var productDto dto.ProductDTO
	context.ShouldBind(&productDto)
	smapping.FillStruct(&product, smapping.MapFields(&productDto))

	data, err := service.productRepository.SetCreateProduct(product)

	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildCreateResponse(http.StatusText(http.StatusCreated), data)
	context.JSON(http.StatusOK, res)
}

//service of update
func (service *productService) SetProductUpdateService(context *gin.Context) {

	product := entitys.Product{}
	var productDto dto.ProductDTO

	context.ShouldBind(&productDto)
	err := smapping.FillStruct(&product, smapping.MapFields(&productDto))
	checkError(err)
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	findById, _ := service.productRepository.GetFindProductById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}

	data, err := service.productRepository.SetCreateProduct(product)
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildUpdateResponse(http.StatusText(http.StatusCreated), data)
	context.JSON(http.StatusOK, res)

}

//service of all
func (service *productService) GetProductAllService(context *gin.Context) {

	var Products, err = service.productRepository.GetAllProduct()
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(http.StatusText(http.StatusOK), http.StatusText(http.StatusOK), Products)
	context.JSON(http.StatusOK, res)
}

func (service *productService) SetProductRemoveService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	findById, _ := service.productRepository.GetFindProductById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	if err != nil {
		validadErrors(err, context)
		return
	}
	result, err := service.productRepository.SetRemoveProduct(findById)

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
