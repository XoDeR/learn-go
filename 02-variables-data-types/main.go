package main

import "fmt"

// decl without init
var grind string

// decl with init
var grinding = "Learning is fun"

// mult declarations
var play, work string = "Play", "Work"

// or
var (
	watch   string = "Watch"
	observe string = "Observe"
)

// inferred type
var untyped = "What's the type?"

func main() {
	// shorthand decl
	// works only inside function bodies
	shorthand := "Shorthand!"

	fmt.Println(shorthand)

	// constant
	const constant = "This is a constant"

	// only constants can be assigned to other constants
	const a = 10
	const b = a

	var aa = 10
	// invalid constant
	// const bb = aa
	fmt.Println(aa)

	// string
	var name string = "Name" // double quotes
	var surname string = `van
	Bossum` // backticks, can span multiple lines

	fmt.Println(name)
	fmt.Println(surname)

	// bool
	var isItTrue bool = true
	fmt.Println(isItTrue)
	// true or false

	// ops with bool:
	// logical: && || !
	// equality: == !=

	fmt.Println("Output from 02-var-data-types.")
}
