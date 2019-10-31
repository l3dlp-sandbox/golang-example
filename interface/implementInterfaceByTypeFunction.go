package _interface

type TestInterface interface {
	Test(string,int64)
}

type TestFunc  func(string,int64)

func(t TestFunc)Test(s string,i int64){
	t(s,i)
}