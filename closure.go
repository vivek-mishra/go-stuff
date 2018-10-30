package main

import (
	"fmt"
)
type People struct {
	Name  string
	Likes []string
}
type names string
// function that can receive a type other than struct, here per function receives names type of string and can be invoked by an instance of names type eg. n.per()
func (n names) per() {
	fmt.Println(n)
}

// anonymous/closure function example, func()
func str() {
	message := "Hello golangcode.com readers!"
	func (message string) {
		fmt.Println(message)
	}(message)

}
// normal function with People as receiver to per returning string
func  per(p People) string {
	fmt.Println(p.Name)
	return p.Name
}

// per function that receives a type of People and is a method to struct type and invoked by an instance of People say p1.per()
func (p People) per() string {
	fmt.Println(p.Name)
	return p.Name
}

// here person is a method, methods with pointer receivers take either a value or a pointer as the receiver when they are called: idea is pointer changes the master struct name variable value to vm, without a *People will just change the local copy of name without affecting the struct name variable value
//Go interprets the statement p2.person("x") as (&p2).person("x") since the person method has a pointer receiver.
func (p *People) person() string  {
	p.Name = "vm"
	return  p.Name
}
// function pers is a function and not a method with a type pointer and pers can be invoked by &p1 that passes the values to the pointer in pers method
func pers(p *People) string {
	//p.Name = "point"
	return p.Name
}

func main() {
	var i int = 1
	j := &i //point to i
	fmt.Println("j by ref", *j) // read i through the pointer
	*j = 21 // set i to 21 through a pointer
	fmt.Println(j)
	fmt.Println(i)

	p1 := People{"name", []string{"name"}} // has type People
	p2 := &p1
	p3 := &People{"vi", []string{"name", "vi"}}  // has type *People
	fmt.Println("p3", *p3)
	fmt.Println("read p1 through p2", *p2)
	per(p1)
	//inderection - function that receives arguments as value can be passed value or pointer
	per(*p2)
	fmt.Println((*p2).per())

	fmt.Println(p1.per())
	fmt.Println("printing add of p1", *p2)
	fmt.Println((&p1).per())
	fmt.Println(pers(&p1))

	fmt.Println(p2.person())
	fmt.Println((&p1).person())

	str()
	n := names("vivek")
	n.per()
	}
//In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.
