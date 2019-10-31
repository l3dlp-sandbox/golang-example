package pass_by_value_reference

import "fmt"

type Person struct {
	firstName string
	lastName string
}

func changeName(p *Person)  {

	p.firstName = "maljinee"

}
func StartPass()  {
	person := Person{
		firstName:"Damitha",
		lastName:"Dayananda",
	}
	fmt.Println(person.firstName)
	changeName(&person)
	fmt.Println(person.firstName)
}
