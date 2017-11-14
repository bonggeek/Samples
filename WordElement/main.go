package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	//    "io"
	//    "io/ioutil"
	"os"
)

type element struct {
	symbol   string
	atNumber uint
	name     string
}

func ReadElement(elementStr string) (element, error) {
	splits := strings.Split(elementStr, ",")
	atNum, err := strconv.ParseUint(splits[0], 10, 32)
	if err != nil {
		return element{}, err
	}

	el := element{atNumber: uint(atNum), symbol: splits[1], name: splits[2]}

	return el, nil
}

func main() {
	//elements := make(map[string]element)
	f, err := os.Open("./elementlist.csv")
	if err != nil {
		panic("failed to open file")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		element, err := ReadElement(line)
		if err != nil {
			fmt.Println("Failed to parse element", err)
			panic("Error!!")
		}
		fmt.Println(element)
	}
}
