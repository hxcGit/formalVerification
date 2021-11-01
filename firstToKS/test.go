package main

import (
	"fmt"
	"reflect"
	"strings"
)

type myobject struct {
	Name string
	Sex  int
	Age  int `json:"age"`
}

func main() {
	onevar := "something"
	other := "something else"

	var msg string
	sa := []string{onevar, other}
	for _, v := range sa {
		vName := reflect.TypeOf(v).Name()
		vName = fmt.Sprintf("{{%s}}", vName)
		msg = strings.Replace(msg, vName, v, -1)
	}
}
