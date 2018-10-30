package main

import "fmt"

func caller(f func(string) string) {
	result := f("David")
	fmt.Println(result)
}
type stru struct {
	a string
	b int
}
func main() {
	f := func(name string) string {
		return "Hello " + name
	}
	caller(f)
	// structs equality
	p1 := stru{"left", 5}
	p2 := stru{a: "left", b: 4}
	if p1 == p2 {
		fmt.Println("Same position")
	}

}
