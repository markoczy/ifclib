package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func tokenize(s string) ([]string, error) {
	reg, err := regexp.Compile("(?msU)TYPE(.*)END_TYPE;")
	if err != nil {
		return nil, err
	}
	return reg.FindAllString(s, -1), nil
}

func main() {
	filename := "data/IFC4x3_RC2.exp"
	data, err := ioutil.ReadFile(filename)
	check(err)

	tokens, err := tokenize(string(data))
	check(err)

	// fmt.Println(tokens)
	for _, v := range tokens {
		fmt.Println("*** TOKEN: " + v)
	}
}
