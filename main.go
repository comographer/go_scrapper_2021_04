package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, upperName := lenAndUpper("como")
	fmt.Println(totalLength, upperName)
}

// fmt is package for formatting

// functions in Go start with uppercase in order to export

// when creating variable, you need to tell type

// when within func, you can create variable like {varaible_name} := {variable}

// when making function, you need to tell Go what types the argument and result will be
