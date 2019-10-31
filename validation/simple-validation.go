package validation

import (
	"gopkg.in/go-playground/validator.v9"
	"fmt"
)

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

var validate *validator.Validate

func simpleValidator()  {
	validate =validator.New()
}

func validateStruct()  {
	address:=&Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	err:=validate.Struct(address)

	if err!=nil{
		for _,err:= range err.(validator.ValidationErrors){
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
	}
}

func validateVariable()  {
	myGmail:="damithadayananda@gmail.com"
	err:=validate.Var(myGmail, "required,email")

	if err !=nil{
		fmt.Println(err)
	}
}
