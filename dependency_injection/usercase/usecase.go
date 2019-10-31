package usercase

import (
	"test/test/dependency_injection/entity"
	"fmt"
)

type UserSaver interface {
	Save(user entity.User) error
}

type UserRegisterUseCase struct {
	userSaver UserSaver
}

func NewUserRegisterUseCase(userSaver UserSaver){
	y:= &UserRegisterUseCase{}
	y.userSaver = userSaver
}

func (u UserRegisterUseCase)RegisterUser(user entity.User)error  {
	err := u.userSaver.Save(user)
	fmt.Println("here")
	if err != nil{
		return  err
	}
	return nil
}

