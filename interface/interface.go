package _interface

import "fmt"

//just like class
type User struct {
	FirstName string
	LastName string
}


type Name interface{
	Name() string
}

//User class has implemented Namer interface
//which means User is also Namer type itself
func(u *User) Name() string{
	return fmt.Sprintf("%s %s",u.FirstName,u.LastName)
}


//Great accepts namer types
//so User type also because
//User implement Namer interface
func Great(n Name) string{
	return fmt.Sprintf("Dear %s",n.Name())
}

func InterfaceTest(){
	u:=&User{
		"Matt",
		"Alice",
	}

	fmt.Println(Great(u))
}

type second struct {

}

