package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(findFactorExact6(262144))
	fmt.Println(findFactorExact6(134217728))
}

func findFactorExact6(number float64) int{
	primes := getPrimeNumbers(number)

	var total int
	for index, prime1 := range primes[:] {
		if math.Pow(prime1, float64(5)) < number {
			total += 1
		}

		for _, prime2 := range primes[index+1:] {
			number1 := prime1 * prime2 * prime2;
			number2 := prime1 * prime1 * prime2;
			if number1 < number {
				total +=1
			}
			if number2 < number {
				total +=1
			}
			if number1 > number && number2 > number {
				break
			}
		}  
	}

	return total
} 

func getPrimeNumbers(number float64) []float64 {
	notPrimeMap := make(map[float64]bool)

	var primes []float64
	for candidate := float64(2); candidate <= number / 4; candidate++ {
		if notPrimeMap[candidate] {
			continue;
		}
		
		primes = append(primes, candidate)
		for notPrime := candidate * 2; notPrime <= number / 4; notPrime+= candidate {
			notPrimeMap[notPrime] = true
		}
	}
	
	return primes;
}