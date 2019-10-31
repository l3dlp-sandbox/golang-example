package inheritence

type  base struct{
}

type first interface {
	pay() int
	void()
}

type second interface {
	first
	dialogCall() string
	interBlockCall() int
}

func(b *base)dialogCall()string{
	return ""
}
func(b *base)interBlockCall()int{
	return 1
}
func(b *base)pay()int{
	return 1
}
func(b *base)void(){

}