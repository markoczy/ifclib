package parser

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/markoczy/ifclib/xp/elems"
)

var matcher = regexp.MustCompile(`\w+|'|\=|,|;|\(|\)|<=|>=|<|>|{|}|\[|\]|\:`)

// func TokenizeTypeDefinitions(s string) [][]string {
// 	return tokenize(s, "(?msU)TYPE(.*)END_TYPE;")
// }

// func TokenizeEntityDefinitions(s string) [][]string {
// 	return tokenize(s, "(?msU)ENTITY(.*)END_ENTITY;")
// }

// func tokenize(s, regex string) [][]string {
// 	ret := [][]string{}
// 	reg := regexp.MustCompile(regex)
// 	for _, v := range reg.FindAllString(s, -1) {
// 		tokens := matcher.FindAllString(v, -1)
// 		ret = append(ret, tokens)
// 	}
// 	return ret
// }

func InitElementMap(s string) elems.Map {
	rx := regexp.MustCompile(`^\s*(TYPE|ENTITY)\s+(?P<name>\w+).*`)
	names := []string{}
	lines := strings.Split(s, "\n")
	fmt.Println(strings.Join(lines, ":\n"))
	idx := rx.SubexpIndex("name")
	for _, line := range lines {
		submatch := rx.FindStringSubmatch(line)
		fmt.Println("Submatch", submatch)
		if len(submatch) > idx {
			names = append(names, submatch[idx])
		}
	}
	return elems.NewMap(names)
}

func assert(b bool, s string) {
	if !b {
		panic(fmt.Errorf(s))
	}
}

func popAndAssertEquals(expected string, queue *tokenQueue) {
	token := queue.Pop()
	assert(token.Content() == expected, fmt.Sprintf("Expected '%s' found %s at line %d and position %d", expected, token.Content(), token.Line(), token.Begin()))
}

// func popAndAssert(queue *)

// TODO remove me
func noop(i ...interface{}) {}
