package main

import (
	"fmt"
	"math/big"
)

func isPrime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(0)
}

func isPrimeStr(n string) (bool, error) {
	z := new(big.Int)
	_, err := fmt.Sscan(n, z)
	if err != nil {
		return false, err
	}
	return z.ProbablyPrime(20), nil
}
