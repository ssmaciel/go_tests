package main

import (
	"fmt"

	"github.com/ssmaciel/go_tests/fizzbuzz"
)

func main() {

	for numero := 1; numero <= 100; numero++ {
		fmt.Println(fizzbuzz.Fizzbuzz(numero))
		//Hello()
	}
}
