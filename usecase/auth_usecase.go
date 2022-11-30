package usecase

import (
	"warung_nasi_padang/model"
	"warung_nasi_padang/utils/authenticator"
)

type AuthUseCase interface {
	UserAuth(user model.UserCredential)(token string, err error)
}

type authUseCase struct {
	tokenService authenticator.AccessToken
}

func (a *authUseCase) UserAuth(user model.UserCredential)(token string,err error) {
	if user.Username == "admin" && user.Password == "123" {
		token,err := a.tokenService.CreateAccessToken(&user)
		if err != nil {
			return "",err
		}
		return token, nil
	}else{
		return "",nil
	}
}

func NewAuthUseCase(service authenticator.AccessToken) AuthUseCase {
	authUseCase := new(authUseCase)
	authUseCase.tokenService = service
	return authUseCase
}