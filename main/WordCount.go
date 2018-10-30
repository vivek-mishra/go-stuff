package main

import (
	"fmt"
	"strings"
)

func main()  {
	m := wordcount("counting words")
	fmt.Println( m)
}
func wordcount(s string) map[string]int {
	var	j int
	if len(s) > 0 {
		strs := strings.Split(s, " ")
		//for i,wc := range strs {
		for i:= 0; i < len(strs); i++ {
		fmt.Println(strs[i])
			j++
		}
	} else {
		fmt.Println("can't count words")
	}
	countMap := map[string]int{s : j}
	return countMap
}


