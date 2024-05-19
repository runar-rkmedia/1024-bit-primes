package main

import (
	"fmt"
	"math/big"
)

func isPrime(n int64) bool {
	probablyCorrect := big.NewInt(n).ProbablyPrime(0)
	myTruth := trialDivisionSimple(uint16(n))
	if probablyCorrect != myTruth {
		panic(fmt.Sprintf("Mytruth does not agree with math %f != %f", myTruth, probablyCorrect))
	}
	return myTruth
}

func isPrimeStr(n string) (bool, error) {
	z := new(big.Int)
	_, err := fmt.Sscan(n, z)
	if err != nil {
		return false, err
	}
	return z.ProbablyPrime(20), nil
}
