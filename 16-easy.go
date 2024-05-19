package main

import (
	"fmt"
	"math/rand"
)

func easy16bit() int64 {
	n := 0
	for {
		n++
		i := rand.Int63()
		if isPrime(i) {
			fmt.Println("found prime after attempts", n, i)
			return i
		}
	}
}
