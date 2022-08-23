package greet

import "fmt"

var Human = "Jert"

type Octopus struct {
	Name  string
	Color string
}

func (o Octopus) String() string {
	return fmt.Sprintf("The octopus name is %q and is the color %s", o.Name, o.Color)
}

func Hello() {
	fmt.Println("Hello")
}
