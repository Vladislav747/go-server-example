package main

import (
	"fmt"
	"testing"
)

func main() {
	var phrase = "Go Производительность"
	fmt.Println(phrase)

}

func BenchmarkFib10(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(40)
	}
	//b.RunParallel(func(pb *testing.PB) {
	//	for pb.Next() {
	//		Fib(10)
	//	}
	//})
}
