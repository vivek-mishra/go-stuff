package main

import "fmt"


func main() {
	 type Person struct {
		Name  string
		Likes []string
	}
	//var people []*Person cities:= []node{node{}}
 people := []*Person{}
	// people = Person{"Vivek", []string{"foo", "bar", "baz"}}

	peep := new(Person)
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