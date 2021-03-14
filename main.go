package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/markoczy/ifclib/parser"
	"github.com/markoczy/ifclib/xp/types"
)

var rxNormalize, _ = regexp.Compile("(\n|\\s)+")

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func tokenize(s string) ([]string, error) {
	ret := []string{}
	reg, err := regexp.Compile("(?msU)TYPE(.*)END_TYPE;")
	if err != nil {
		return nil, err
	}
	for _, v := range reg.FindAllString(s, -1) {
		ret = append(ret, normalize(v))
	}
	return ret, nil
}

var matcher = regexp.MustCompile(`\w+|'|\=|,|;|\(|\)|<=|>=|<|>|{|}|\[|\]|\:`)

func tokenize2(s string) ([][]string, error) {
	ret := [][]string{}
	reg, err := regexp.Compile("(?msU)TYPE(.*)END_TYPE;")
	if err != nil {
		return nil, err
	}
	for _, v := range reg.FindAllString(s, -1) {

		// tokens := strings.Split(normalize(v), " ")
		// tokens := ma
		tokens := matcher.FindAllString(v, -1)
		ret = append(ret, tokens)
	}
	return ret, nil
}

func normalize(s string) string {
	return rxNormalize.ReplaceAllString(s, " ")
}

func main() {
	// testCreateTypes()
	// testTokenize()
	testParse()
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
	ifcActionRequestTypeEnum := types.NewDerived("IfcActionRequestTypeEnum", types.NewEnumeration([]string{"EMAIL", "FAX", "PHONE", "POST", "VERBAL", "USERDEFINED", "NOTDEFINED"}))
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

	tokens, err := tokenize2(string(data))
	check(err)

	// fmt.Println(tokens)
	for _, v := range tokens {
		formatted := `"` + strings.Join(v, `", "`) + `"`
		fmt.Println("*** TOKENS:", formatted)
	}
}

func testParse() {
	input := `TYPE IfcContextDependentMeasure = REAL;
END_TYPE;

TYPE IfcCountMeasure = NUMBER;
END_TYPE;

TYPE IfcGloballyUniqueId = STRING(22) FIXED;
END_TYPE;

TYPE IfcComplexNumber = ARRAY [1:2] OF REAL;
END_TYPE;

TYPE IfcCompoundPlaneAngleMeasure = LIST [3:4] OF INTEGER;
END_TYPE;

TYPE IfcArcIndex = LIST [3:3] OF IfcPositiveInteger;
END_TYPE;

TYPE IfcPositiveInteger = IfcInteger;
END_TYPE;

TYPE IfcInteger = INTEGER;
END_TYPE;

TYPE IfcMooringDeviceTypeEnum = ENUMERATION OF
	(LINETENSIONER
	,MAGNETICDEVICE
	,MOORINGHOOKS
	,VACUUMDEVICE
	,BOLLARD
	,USERDEFINED
	,NOTDEFINED);
END_TYPE;`

	mp := parser.InitElementMap(input)
	tokens := parser.TokenizeTypeDefinitions(input)

	fmt.Println(mp)
	for _, v := range tokens {
		fmt.Println("*** Tokens:", v)
		tp := parser.ParseType(v, mp)
		mp.Assign(tp.Name(), tp)
	}

	fmt.Println("TypeMap:", mp)
}

func noop(i ...interface{}) {}
