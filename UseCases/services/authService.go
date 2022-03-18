package services

import (
	"github.com/flabio/Core/entitys"
	"github.com/flabio/Core/repositorys"
)

type AuthService interface {
	VerifyCredential(Name string, Owner string) interface{}
	GetOwner(Id uint) string
}

type authService struct {
	authRepository repositorys.AdministradorRepository
}

func NewAuthService() AuthService {

	return &authService{
		authRepository: repositorys.NewAdministradorRepository(),
	}
}

/*
@param Owner is of type string
*/
func (authService *authService) VerifyCredential(Name string, Owner string) interface{} {
	res := authService.authRepository.VerifyCredential(Name, Owner)

	if v, ok := res.(entitys.Administrador); ok {

		if v.Owner == Owner {
			return res
		}

	}
	return nil
}

func (authService *authService) GetOwner(Id uint) string {

	user, _ := authService.authRepository.GetFindAdministradorById(Id)
	return user.Criticidad
}
