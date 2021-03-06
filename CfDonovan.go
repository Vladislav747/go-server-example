// Cf конвертирует числовой аргумен в температуру
// по Цельсию и по Фаренгейту,
package main

import (
	"fmt"
	"go-server-example/storage/popcount"
	"go-server-example/storage/tempconv"
	"os"
	"strconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
		popcount.Init()
		popcountVal := popcount.PopCount(40)
		fmt.Println(popcountVal)
	}
}
