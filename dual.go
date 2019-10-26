package vyut

import (
	"fmt"
	"math"
)

// Dual is a dual number of the form a + bε
type Dual struct {
	real, dual float64
}

// Package functions

// NewDual function will allow you to create new dual numbers
func NewDual(r, d float64) *Dual {
	return &Dual{real: r, dual: d}
}

// ConstDual funnction will allow you to create new dual numbers
func ConstDual(r float64) *Dual {
	return &Dual{real: r, dual: 0}
}

// GetVal returns the real part of the dual number
func GetVal(d *Dual) float64 {
	return d.real
}

// GetDeriv returns the derivative from the dual part of the dual number
func GetDeriv(d *Dual) float64 {
	return d.dual
}

// Methods

// IsEqual compares two dual numbers
func (d *Dual) IsEqual(other *Dual) bool {
	return d.real == other.real && d.dual == other.dual
}

// Add will add two dual numbers
func (d *Dual) Add(other *Dual) *Dual {
	return &Dual{d.real + other.real, d.dual + other.dual}
}

// Sub will subtract one dual number from another
func (d *Dual) Sub(other *Dual) *Dual {
	return &Dual{d.real - other.real, d.dual - other.dual}
}

// Mul will multiply two dual numbers
func (d *Dual) Mul(other *Dual) *Dual {
	return &Dual{d.real * other.real, d.real*other.dual + d.dual*other.real}
}

// Div will divide one dual number from another
func (d *Dual) Div(other *Dual) *Dual {
	return &Dual{d.real / other.real, (d.dual*other.real - d.real*other.dual) / (other.real * other.real)}
}

// Neg will create the negative of the dual number
func (d *Dual) Neg() *Dual {
	return &Dual{-d.real, -d.dual}
}

// Pow will raise the dual number to some power n
func (d *Dual) Pow(n float64) *Dual {
	return &Dual{math.Pow(d.real, n), float64(n) * d.dual * math.Pow(d.real, n-1)}
}

// Sin will raise the dual number to some power n
func (d *Dual) Sin() *Dual {
	return &Dual{math.Sin(d.real), d.dual * math.Cos(d.real)}
}

// Cos will raise the dual number to some power n
func (d *Dual) Cos() *Dual {
	return &Dual{math.Cos(d.real), -d.dual * math.Sin(d.real)}
}

// Exp will raise the dual number to some power n
func (d *Dual) Exp() *Dual {
	return &Dual{math.Exp(d.real), d.dual * math.Exp(d.real)}
}

// Log will raise the dual number to some power n
func (d *Dual) Log() *Dual {
	return &Dual{math.Log(d.real), d.dual / d.real}
}

// String format for the type
func (d *Dual) String() string {
	var str string

	if d.dual < 0 {
		str = fmt.Sprintf("%.2f - %.2fε", d.real, -d.dual)
	} else {
		str = fmt.Sprintf("%.2f + %.2fε", d.real, d.dual)
	}

	return str
}
