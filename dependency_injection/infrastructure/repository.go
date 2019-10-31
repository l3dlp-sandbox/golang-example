package infrastructure

import (
	"database/sql"
	"test/test/dependency_injection/entity"
	"fmt"
)

type UserRepository struct {
	sql *sql.DB
}

func NewUserRepository(sql *sql.DB)*UserRepository  {
	return &UserRepository{sql}
}

func(u UserRepository) Save(user entity.User)error {
	fmt.Println("successfully saved")
	err:=u.sql.Ping()
	if err !=nil{
		fmt.Println(err)
	}
	return  nil
}

