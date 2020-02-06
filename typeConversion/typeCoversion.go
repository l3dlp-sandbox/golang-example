package typeConversion

import (
	"fmt"
	"reflect"
	"unsafe"
)

//static type language like java/c++ support implicit type
//conversion
//private void demo(){
//	int i=5.5
//}
//will compile without error but golang
//doesn't support implicit type conversion

func TypeConversionDemo() {
	//Go type conversion will do it's best to
	//maintain the same value in the new data
	//type if possible. To do that it may transform
	//the underlying bit structure
	var x float64
	var y = 32
	x = float64(32)
	fmt.Println(x, "type of x is", reflect.TypeOf(x))
	fmt.Println("type of y is", reflect.TypeOf(y))
	//
	//casting in go doesn't change underlying data structure
	type Common struct {
		Gender int
		From   string
		To     string
	}

	type Foo struct {
		Id    string
		Name  string
		Extra Common
	}

	type Bar struct {
		Id    string
		Name  string
		Extra Common
	}
	foo := Foo{
		Id:   "123",
		Name: "damitha",
		Extra: struct {
			Gender int
			From   string
			To     string
		}{Gender: 1, From: "xx", To: "yy"},
	}
	bar := *(*Bar)(unsafe.Pointer(&foo))
	fmt.Printf("%+v\n", bar)
}
