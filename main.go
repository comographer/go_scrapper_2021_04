package main

import "fmt"

func canIDrink(age int) bool {
	if age < 18 {
		return false
	}
	return true
}

func main() {
	fmt.Println(canIDrink(16))
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
