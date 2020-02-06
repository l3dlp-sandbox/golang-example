package generic

import "fmt"

type stringList []string
type intList []int

type list interface {
	size()
	add(interface{})
}

func (l *stringList) size() {
	fmt.Println("string array size:", len(*l))
}
func (l stringList) add(arr interface{}) {

}

func (l *intList) size() {
	fmt.Println("int array size:", len(*l))
}
func (l intList) add(interface{}) {

}

func GenericDemo() {
	s := []string{"a", "b"}
	stringList(s).size()
	i := []int{1, 5}
	intList(i).size()
	fmt.Print()
}

//**************************************************

type myinterface interface {
	method()
}

type mystruct struct{}

func (ms mystruct) method() {

}
func (ms mystruct) anothermethod() {

}

type mystruct1 struct{}

func (ms mystruct1) method() {

}
func (ms mystruct1) anothermethod() {

}

func GenericDemoTwo() {
	ms := mystruct{}
	testfunction(ms)
	testfunctiontwo(ms)
}

func testfunction(mi myinterface) {
	mi.(mystruct).method() //type assertion use to extract concrete type from its parent interface
}

//type switching can use to identify type
func testfunctiontwo(mi interface{}) {
	switch _ := mi.(type) {
	case string:
		fmt.Print("")
	}
}

//since empty interface{} doesn't include any methods any go type implicitly child of empty interface

//in allmost all cases when we make use of the empty interface in go, we have to convert the value back to
//the original type at runtime in order to obtain original value(type assertion or type switch) it has some performance overhead

//generics in other hand doesn't work in this way for generics, type checking happens at compile time
//no visible overhead at runtime
