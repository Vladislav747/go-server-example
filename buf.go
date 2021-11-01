package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var phrase = "Go Buf"
	fmt.Println(phrase)

	//counts := make(map[string]int)
	input := bufio.NewReader(os.Stdin)

	
	for {
		line, _ := input.ReadString('\n')
		fmt.Print(line)
  	}

	// for input.Scan() {
	// 	counts[input.Text()]++
	// }

	// for line, n := range counts {
	// 	if n > 1 {
	// 		fmt.Printf("%d\t%s\n", n, line)
	// 	}
	// }

}
