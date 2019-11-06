package vyut

import (
	"fmt"
	"math"
)

// Dual is a dual number of the form a + bε
type Dual struct {
	real, grad float64
}

// Package functions

// NewDual function creates new dual numbers from the real and dual values
func NewDual(r, d float64) *Dual {
	return &Dual{real: r, grad: d}
}

// FromReal function creates creates a dual number from real numeric types like floats and ints
func FromReal(r float64) *Dual {
	return &Dual{real: r, grad: 0}
}

// GetVal returns the real part of the dual number
func GetVal(d *Dual) float64 {
	return d.real
}

// Gradient returns the derivative from the dual part of the dual number
func Gradient(d *Dual) float64 {
	return d.grad
}

// Methods

// IsEqual compares two dual numbers
func (d *Dual) IsEqual(other *Dual) bool {
	return d.real == other.real && d.grad == other.grad
}

// Add will add two dual numbers
func (d *Dual) Add(other *Dual) *Dual {
	return &Dual{d.real + other.real, d.grad + other.grad}
}

// Sub will subtract one dual number from another
func (d *Dual) Sub(other *Dual) *Dual {
	return &Dual{d.real - other.real, d.grad - other.grad}
}

// Mul will multiply two dual numbers
func (d *Dual) Mul(other *Dual) *Dual {
	return &Dual{d.real * other.real, d.real*other.grad + d.grad*other.real}
}

// Div will divide one dual number from another
func (d *Dual) Div(other *Dual) *Dual {
	return &Dual{d.real / other.real, (d.grad*other.real - d.real*other.grad) / (other.real * other.real)}
}

// Neg will create the negative of the dual number
func (d *Dual) Neg() *Dual {
	return &Dual{-d.real, -d.grad}
}

// Pow will raise the dual number to some power n
func (d *Dual) Pow(n float64) *Dual {
	return &Dual{math.Pow(d.real, n), float64(n) * d.grad * math.Pow(d.real, n-1)}
}

// Sin will raise the dual number to some power n
func (d *Dual) Sin() *Dual {
	return &Dual{math.Sin(d.real), d.grad * math.Cos(d.real)}
}

// Cos will raise the dual number to some power n
func (d *Dual) Cos() *Dual {
	return &Dual{math.Cos(d.real), -d.grad * math.Sin(d.real)}
}

// Exp will raise the dual number to some power n
func (d *Dual) Exp() *Dual {
	return &Dual{math.Exp(d.real), d.grad * math.Exp(d.real)}
}

// Log will raise the dual number to some power n
func (d *Dual) Log() *Dual {
	return &Dual{math.Log(d.real), d.grad / d.real}
}

// String format for the type
func (d *Dual) String() string {
	var str string

	if d.grad < 0 {
		str = fmt.Sprintf("%.2f - %.2fε", d.real, -d.grad)
	} else {
		str = fmt.Sprintf("%.2f + %.2fε", d.real, d.grad)
	}

	return str
}
