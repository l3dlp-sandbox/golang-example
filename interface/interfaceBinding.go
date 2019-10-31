package _interface

type Check interface {
	Check1(a string)
	Check2(b string)
}

type first struct {

}
func (*first)Check1(a string){

}
func (*first)Check2(b string){

}





