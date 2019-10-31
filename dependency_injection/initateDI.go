package dependency_injection

import (
	"database/sql"
	"test/test/dependency_injection/infrastructure"
	"test/test/dependency_injection/usercase"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"test/test/dependency_injection/entity"
)


func InitiateDI(){
	db,err:=sql.Open(`mysql`,"root:@tcp(localhost:3306)/driver_disconnect")
	if err!=nil{
		fmt.Println("error happened")
	}
	userRepository:=infrastructure.NewUserRepository(db)
	usercase.NewUserRegisterUseCase(userRepository)
	x:=usercase.UserRegisterUseCase{}
	dy:=entity.User{
		Email:"damithadayananda@gmail.com",
		Password:"damitha",
		Name:"Damitha",
		DateOfBirth:"1991.9.1",
	}

	x.RegisterUser(dy)
}
