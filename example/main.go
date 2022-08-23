package main

import (
	"fmt"
	"github.com/packageslearn/greet"
)

func main() {
	greet.Hello()

	fmt.Println(greet.Human)

	oct := greet.Octopus{
		Name:  "Jar",
		Color: "Abob",
	}
	fmt.Println(oct.String())

}
