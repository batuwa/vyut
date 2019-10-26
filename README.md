# Vyut

 Go library for Auto Differentiation (experimental)

> Adjective (english translation: _derivative_)
>
> Vyutpannna (व्युत्पन्न) - Sanskrit word meaning **derivative**, derived or descended. 

## Description

_Yyut_ is a [Automatic Differentiation (AD)](https://en.wikipedia.org/wiki/Automatic_differentiation) library written in the Go programming language. It provides a simple API to calculate derivatives of mathematical functions, which is widely used in Deep Learning, Robotics and other fields of numeric computing.

Note: So far, a simple [Forward Differentiaion](https://en.wikipedia.org/wiki/Automatic_differentiation#Automatic_differentiation_using_dual_numbers) algorithm using dual numbers has been implemented.

## Usage

### Import the package

```go
go get github.com/batuwa/vyut

import . "github.com/batuwa/vyut"
```

### Code example

```go

// Dual numbers example
d1 := NewDual(5, 1)
d2 := NewDual(-2, 1)

fmt.Println(d1.Add(d2))
fmt.Println(d1.Sin())

// Function differentiation
func f(x *Dual) *Dual {
    return ConstDual(1.0).Add(x.Exp())   // 1 + e^x
}

x := NewDual(3, 1)
y := f(x)

fmt.Println("Derivative of f(x) at x=3 f'(3) is", GetDeriv(y))


```

Run the main.go in `examples` folder for a more comprehensive example.
