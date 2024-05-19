package main

import (
	"fmt"
)

func easy16bit() uint16 {
	n := 0
	for {
		n++
		i := randomU16() | 0b1000000000000001
		if isPrime(int64(i)) {
			fmt.Println("found prime after attempts", n, i)
			return i
		}
		fmt.Println("Not a prime", n, i)
	}
}
