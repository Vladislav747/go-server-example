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
		Fib(10)
	}

}

func BenchmarkFibParallel10(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Fib(10)
		}
	})
}
