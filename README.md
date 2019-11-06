# Vyut

 Go library for Auto Differentiation (experimental)

> Adjective (english translation: _derivative_)
>
> Vyutpannna (व्युत्पन्न) - Sanskrit word meaning **derivative**, derived or descended. 

## Description

_Vyut_ is a [Automatic Differentiation (AD)](https://en.wikipedia.org/wiki/Automatic_differentiation) library written in the Go programming language. It provides a simple API to calculate derivatives of mathematical functions, which is widely used in Deep Learning, Robotics and other fields of numeric computing.

Note: So far, a simple [Forward Differentiaion](https://en.wikipedia.org/wiki/Automatic_differentiation#Automatic_differentiation_using_dual_numbers) algorithm using dual numbers has been implemented.

## Usage

### Import the package

```go
go get github.com/batuwa/vyut

import . "github.com/batuwa/vyut"
```

### Code example

```go

// Dual number instantiation
d1 := NewDual(5, -2)
d2 := NewDual(-2, 1.5)

fmt.Println(d1.Add(d2))
fmt.Println(d1.Sin())

// Function to differentiate f(x) = 1 + e^x
func f(x *Dual) *Dual {
    return FromReal(1.0).Add(x.Exp())
}

// Input scalar variable in dual form for differentiation
x := NewDual(3, 1)

// Variable to store the funcation value at x, and store its gradient 
y := f(x)

// Obtain the derivative of y wrt x
df = Gradient(y)

// The Gradient is available via the function GetGrad(y)
fmt.Printf("Derivative of f(x) at x=3 is f'(3) = %.2f", df)
```

Output is:

```bash
3.00 - 0.50ε
-0.91 - 0.62ε

Derivative of f(x) at x=3 is df(3) = 20.09
```

Run the main.go in `examples` folder for a more comprehensive example.
