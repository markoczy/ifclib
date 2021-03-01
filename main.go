package main

import (
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/markoczy/ifclib/xp/types"
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
	testCreateTypes()
}

func testCreateTypes() {
	// Simple String derived type
	ifcDate := types.NewDerived("IfcDate", types.String)
	fmt.Println("IfcDate: ", ifcDate)
	fmt.Println()

	// Fixed String derived type
	ifcGloballyUniqueId := types.NewDerived("IfcGloballyUniqueId", types.NewString(22, 22, true))
	fmt.Println("IfcGloballyUniqueId: ", ifcGloballyUniqueId)
	fmt.Println()

	// Enumeration derived
	ifcActionRequestTypeEnum := types.NewEnumeration("IfcActionRequestTypeEnum", []string{"EMAIL", "FAX", "PHONE", "POST", "VERBAL", "USERDEFINED", "NOTDEFINED"})
	fmt.Println("IfcActionRequestTypeEnum: ", ifcActionRequestTypeEnum)
	fmt.Println()

	// 3-Layer derived List type
	ifcInteger := types.NewDerived("IfcInteger", types.Integer)
	ifcPositiveInteger := types.NewDerived("IfcPositiveInteger", ifcInteger)
	ifcLineIndex := types.NewList(2, -1, ifcPositiveInteger)
	fmt.Println("IfcLineIndex: ", ifcLineIndex)
	fmt.Println()
}

func testTokenize() {
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
