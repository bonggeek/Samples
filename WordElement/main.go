package main

import (
	"fmt"
	element "github.com/bonggeek/element"
	//    "io"
	//    "io/ioutil"
)



func main() {
	var elements element.Elements
	elements.LoadPeriodicTable("./elementlist.csv")

	word1 := "basu"
	fmt.Println(word1)
	result2 := elements.GetElementsForWord(word1)
	for _, k := range result2 {
		fmt.Println(k)
	}


	word2 := "how"
	fmt.Println(word2)
	result1 := elements.GetElementsForWord(word2)
	for _, k := range result1 {
		fmt.Println(k)
	}
}
