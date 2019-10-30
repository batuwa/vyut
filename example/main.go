package main

import (
	"fmt"

	. "github.com/batuwa/vyut"
)

func main() {
	fmt.Println("\nBasic operations:")

	// Creating dual numbers with both parts
	d1 := NewDual(1, 1)
	d2 := NewDual(2, -4)

	// Creating dual numbers with only real component
	r := FromReal(2.5)
	fmt.Printf("2.5      = %s\n", r)

	// Basic math functions for two dual numbers
	fmt.Printf("d1 == d2 = %t\n", d1.IsEqual(d2))
	fmt.Printf("d1 + d2  = %s\n", d1.Add(d2))
	fmt.Printf("d1 - d2  = %s\n", d1.Sub(d2))
	fmt.Printf("d1 * d2  = %s\n", d1.Mul(d2))
	fmt.Printf("d1 / d2  = %s\n", d1.Div(d2))

	// Unary operations in dual numbers
	fmt.Printf("-d       = %s\n", d2.Neg())
	fmt.Printf("d**2     = %s\n", d2.Pow(4))
	fmt.Printf("sin(d)   = %s\n", d2.Sin())
	fmt.Printf("cos(d)   = %s\n", d2.Cos())
	fmt.Printf("exp(d)   = %s\n", d2.Exp())
	fmt.Printf("log(d)   = %s\n", d2.Log())

	fmt.Println("\n\nCalculate derivate:")

	// Creating variables to calculate derivatives with the dual part 1
	x := NewDual(1.5, 1)

	// Testing out on some basic functions
	y := simpleFunction(x)
	fmt.Println("f(x) = x * sin(x*x) + 1")
	fmt.Printf("f(1.5) = %.2f, f'(1.5) = %.2f\n", GetReal(y), GetGrad(y))

	sigma := sigmoid(x)
	fmt.Println("\nσ(x) = 1 / (1 + e^-z)")
	fmt.Printf("σ(1.5) = %.2f, σ'(1.5) = %.2f\n", GetReal(sigma), GetGrad(sigma))

	tanh := tanh(x)
	fmt.Println("\ntanh(x) = (e^z - e^-z) / (e^z + e^-z)")
	fmt.Printf("tanh(1.5) = %.2f, tanh'(1.5) = %.2f\n", GetReal(tanh), GetGrad(tanh))
}

// f is a test function which we are going to calculate the derivative of
func simpleFunction(x *Dual) *Dual {
	return x.Mul(x.Mul(x).Sin()).Add(FromReal(1))
}

// sigmoid function used in deep learning
func sigmoid(z *Dual) *Dual {
	return FromReal(1).Div(FromReal(1).Add(z.Neg().Exp()))
}

// tanh function used in deep learning
func tanh(z *Dual) *Dual {
	return z.Exp().Sub(z.Neg().Exp()).Div(z.Exp().Add(z.Neg().Exp()))
}
