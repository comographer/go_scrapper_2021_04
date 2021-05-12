package main

import "fmt"

func main() {
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn")
	fmt.Println(names)
}

// fmt is package for formatting

// functions in Go start with uppercase in order to export

// when creating variable, you need to tell type

// when within func, you can create variable like {varaible_name} := {variable}

// when making function, you need to tell Go what types the argument and result will be

// Go can return multiple values from a single function

// If you want only one value out of multiple value function, use _ <= this is ignore value

// defer executes after func is done

// for is the only way to loop in GO

// with GO, you can create variable creation in if statement (Variable Expression)

// switch can be used to avoid too much if else

// adding & shows memory location * see through of memory
