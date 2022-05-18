package main

import (
	"flag"
	"fmt"
)

/* Default значение World */
var name = flag.String("name", "World", "A name to say hello to")

var spanish bool

func init() {
	flag.BoolVar(&spanish, "spanish", false, "Use Spanish language")
	flag.BoolVar(&spanish, "s", false, "Use Spanish language")
}

func main() {
	flag.Parse()
	if spanish == true {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hola %s!\n", *name)
	}
}
