package main

import (
	"fmt"
	"math"
)

// TESTE ...
func main() {
	fmt.Println("IR")
	ir := IR(1.3, 1, 36, 100000, 0)
	fmt.Println(ir)
}

// IPMT ...
func IPMT(rate float64, period int32, periods int32, present float64, future float64, tipo float64) float64 {

	tipo = 0.0

	future = 0.0

	payment := PMT(rate, periods, present, future, tipo)

	var interest float64

	if period == 1 {
		if tipo == 1 {
			interest = 0
		} else {
			interest = -present
		}
	} else {
		if tipo == 1 {
			interest = FV(rate, period-2, payment, present, 1) - payment
		} else {
			interest = FV(rate, period-1, payment, present, 0)
		}
	}

	return interest * rate
}

//PMT ...
func PMT(rate float64, periods int32, present float64, future float64, tipo float64) float64 {
	// Credits: algorithm inspired by Apache OpenOffice

	// Initialize tipo
	tipo = 0.0

	// Initialize future
	future = 0.0

	// Evaluate rate (TODO: replace with secure expression evaluator)
	// rate = eval(rate)

	// Return payment
	var result float64
	if rate == 0 {
		result = (present + future) / float64(periods)
	} else {
		var term = math.Pow(1+rate, float64(periods))
		if tipo == 1 {
			result = (future*rate/(term-1) + present*rate/(1-1/term)) / (1 + rate)
		} else {
			result = future*rate/(term-1) + present*rate/(1-1/term)
		}
	}

	// result = Math.round(result * 100) / 100;

	return -result
}

//FV ...
func FV(rate float64, periods int32, payment float64, value float64, tipo int32) float64 {
	var result float64

	if rate == 0 {
		result = value + payment*float64(periods)
	} else {
		term := math.Pow(1+rate, float64(periods))

		if tipo == 1 {
			result = value*term + payment*(1+rate)*(term-1.0)/rate
		} else {
			result = value*term + payment*(term-1)/rate
		}
	}

	return -result
}

//IR ...
func IR(retorno float64, fatura int32, prazo int32, valor float64, juros float64) float64 {
	var aliqIR float64
	var ipmt float64

	ipmt = -IPMT(retorno/100, fatura, prazo, valor, 0, 0)
	ipmt = math.Round(ipmt*100) / 100

	if fatura >= 1 && fatura < 6 {
		aliqIR = 0.225
	} else if fatura >= 6 && fatura < 12 {
		aliqIR = 0.2
	} else if fatura >= 12 && fatura < 24 {
		aliqIR = 0.175
	} else {
		aliqIR = 0.15
	}

	return (ipmt + juros) * aliqIR

}
