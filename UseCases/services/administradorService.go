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

//administradorService is a contract.....
type AdministradorService interface {
	SetCreateService(context *gin.Context)
	SetUpdateService(context *gin.Context)
	GetFindByIdService(context *gin.Context)
	SetRemoveService(context *gin.Context)
	GetAllService(context *gin.Context)
}

type administradorService struct {
	administradorRepository repositorys.AdministradorRepository
}

//NewAdministradorService creates a new instance of administradorServicess
func NewAdministradorService() AdministradorService {

	return &administradorService{
		administradorRepository: repositorys.NewAdministradorRepository(),
	}
}

//service of create
func (service *administradorService) SetCreateService(context *gin.Context) {

	admin := entitys.Administrador{}
	var adminDto dto.AdministradorDTO
	context.ShouldBind(&adminDto)
	if validarAdministrador(adminDto, context, utilities.OPTION_CREATE) {
		return
	}
	smapping.FillStruct(&admin, smapping.MapFields(&adminDto))

	data, err := service.administradorRepository.SetCreateAdministrador(admin)

	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildCreateResponse(http.StatusText(http.StatusCreated), data)
	context.JSON(http.StatusOK, res)
}

//service of update
func (service *administradorService) SetUpdateService(context *gin.Context) {

	admin := entitys.Administrador{}
	var adminDto dto.AdministradorDTO

	context.ShouldBind(&adminDto)
	if validarAdministrador(adminDto, context, utilities.OPTION_EDIT) {
		return
	}
	err := smapping.FillStruct(&admin, smapping.MapFields(&adminDto))
	checkError(err)
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	findById, _ := service.administradorRepository.GetFindAdministradorById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}

	data, err := service.administradorRepository.SetCreateAdministrador(admin)
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildUpdateResponse(http.StatusText(http.StatusCreated), data)
	context.JSON(http.StatusOK, res)

}

//service of all
func (service *administradorService) GetAllService(context *gin.Context) {

	var admins, err = service.administradorRepository.GetAllAdministrador()
	if err != nil {
		validadErrors(err, context)
		return
	}
	res := utilities.BuildResponse(http.StatusText(http.StatusOK), http.StatusText(http.StatusOK), admins)
	context.JSON(http.StatusOK, res)
}

func (service *administradorService) SetRemoveService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	findById, _ := service.administradorRepository.GetFindAdministradorById(uint(id))
	if findById.Id == 0 {
		validadErrorById(context)
		return
	}
	if err != nil {
		validadErrors(err, context)
		return
	}
	result, err := service.administradorRepository.SetRemoveAdministrador(findById)

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

func (service *administradorService) GetFindByIdService(context *gin.Context) {

	id, err := strconv.ParseUint(context.Param("id"), 0, 0)

	if err != nil {
		validadErrors(err, context)
		return
	}
	admin, err := service.administradorRepository.GetFindAdministradorById(uint(id))
	if err != nil {
		validadErrors(err, context)
		return
	}

	if (admin == entitys.Administrador{}) {
		validadErrorById(context)
		return
	}
	res := utilities.BuildResponse(http.StatusText(http.StatusOK), http.StatusText(http.StatusOK), admin)
	context.JSON(http.StatusOK, res)

}

//method private

//validarAdministrador

func validarAdministrador(r dto.AdministradorDTO, context *gin.Context, option int) bool {
	context.ShouldBind(&r)
	switch option {
	case 1:
		if len(r.Name) == 0 || r.Name == "" {
			msg := utilities.MessageRequired{}
			validadRequiredMsg(msg.RequiredName(), context)
			return true
		}
		if len(r.Criticidad) == 0 || r.Criticidad == "" {
			msg := utilities.MessageRequired{}
			validadBirdateRequiredMsg(msg.RequiredCriticidad(), context)
			return true
		}
		if len(r.Owner) == 0 || r.Owner == "" {
			msg := utilities.MessageRequired{}
			validadBirdateRequiredMsg(msg.RequiredOwner(), context)
			return true
		}
	case 2:
		if r.Id == 0 {
			msg := utilities.MessageRequired{}
			validadRequiredMsg(msg.RequiredId(), context)
			return true
		}
		if len(r.Name) == 0 {
			msg := utilities.MessageRequired{}
			validadRequiredMsg(msg.RequiredName(), context)
			return true
		}
		if len(r.Criticidad) == 0 || r.Criticidad == "" {
			msg := utilities.MessageRequired{}
			validadBirdateRequiredMsg(msg.RequiredCriticidad(), context)
			return true
		}
		if len(r.Owner) == 0 || r.Owner == "" {
			msg := utilities.MessageRequired{}
			validadBirdateRequiredMsg(msg.RequiredOwner(), context)
			return true
		}
	}
	return false
}
