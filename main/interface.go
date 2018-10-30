package main

import "fmt"

type Peeps interface {
	 peep() string
}

type personame struct {
	name string
}

type namestring string

func (fname namestring) peepname() {
	fmt.Println(fname)
}

func (n personame) peep() string {
	n.name = "vick"
	return n.name
}

func main()  {

	var p Peeps
	p = personame{"vivek"}
	fmt.Println(p.peep())

	var pe namestring
	//pe := namestring("viv")
	pe.peepname()

	ages1 := map[string]int{"John": 30}
	ages2 := map[string]int{"Mary": 28}

	allAges := mergeMaps(ages1, ages2)
	fmt.Println(allAges)
}
func mergeMaps(ages1, ages2 map[string]int) map[string]int {
	allAges := make(map[string]int, len(ages1) + len(ages2))
	fmt.Println(allAges, len(allAges))
	for k, v := range ages1 {
		allAges[k] = v
		fmt.Println(allAges)
	}
	for k, v := range ages2 {
		allAges[k] = v
		fmt.Println(allAges)
	}
	return allAges
}