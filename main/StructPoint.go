package main

import "fmt"

type Person struct {
	Name  string
	Likes []string
}

var m = map[string]Person{
"p1" : Person{
	"vivek", []string{"foo", "bar", "baz"},
},
"p2" : Person{
	"vick", []string{"a", "b", "c"},
},
}
func main() {
	//m = make(map[string]Person)
	m["v"] = Person{"vicky", []string{"n", "a", "m","e"}}
	m1, ok := m[""]
	fmt.Println(m1, ok)
	fmt.Println(m)
	for _,v := range m {
		fmt.Println(v.Likes)
		for n,l := range v.Likes {
			fmt.Println(n,l)
		}
	}
	//var people []*Person cities:= []node{node{}}
	people := []*Person{}
	// people = Person{"Vivek", []string{"foo", "bar", "baz"}}

	peep := new(Person)
	//peep3 := new(Person)
	peep.Name = "vivek"
	peep.Likes = []string{"foo", "bar", "baz"}
	/*	likes := map[string][]string{
			peep.Name : {"foo", "bar", "baz"},
		}*/
	people = append(people, peep)
	fmt.Println(len(people))

	for _, p := range people {
		fmt.Println(p.Likes)
		for k, l := range p.Likes {
			//likes[l] = append(likes[l], p)
			fmt.Println(k, l)
		}
	}
}
