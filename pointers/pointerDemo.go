package pointers

import (
	"fmt"
	"reflect"
)

/**
pointer sote the memory address of another table
data store in memory at particular address memory address looks like(0xAFFFF)
to address data it's required to know the address
variable is just a convenient name given to memory location
pointer is also variable that store the memory address of another variable
 */
 type PointerDemo struct {

 }

 func(*PointerDemo) PointerDemo(){
 	var x =100
 	var p *int=&x
 	fmt.Println("Value stored in variable x=",x)
 	fmt.Println("Address of variable x=",&x)
 	fmt.Println("Value stored in variable p=",p)

 	//it's possible to use * operator to access the value stored in the variable
 	fmt.Println("dereferenced value of pointer p=",*p)

 	//changing value of variable using pointer reference
 	*p=200
 	fmt.Println("changed value x using pointer refernce:",x)
 }
func (*PointerDemo)PointerDemoNew()  {
	ptr:= new(int)
	*ptr=100
	fmt.Printf("Ptr = %#x, Ptr value = %d\n",ptr,*ptr)
}
func(*PointerDemo)PointerToPoinDemo(){
	var a =10
	var p = &a
	fmt.Println(reflect.TypeOf(p))
	fmt.Println("Address of variable a:",p)
	var pp = &p
	fmt.Println(reflect.TypeOf(pp))
	fmt.Println("Address of variable p:",pp)

	fmt.Println("dereferencing **pp",*pp)
	fmt.Println("dereferencing *p",**pp)
}
//go doesn't support pointer arithmetic