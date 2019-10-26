package vyut

import (
	"math"
	"testing"
)

func TestDeriv(t *testing.T) {
	const tolerance = .001
	x := NewDual(1.5, 1)

	y := simpleFunction(x) // f expected -2.05 at 1.5
	if math.Abs(GetDeriv(y)-(-2.0487)) > tolerance {
		t.Errorf("Got %v, expected: %v.", GetDeriv(y), -2.05)
	}

	sigma := sigmoid(x) // sigma expected 0.15 at 1.5
	if math.Abs(GetDeriv(sigma)-0.149) > tolerance {
		t.Errorf("Got %v, expected: %v.", GetDeriv(sigma), 0.15)
	}

	tanh := tanh(x) // tanh expected 0.18 at 1.5
	if math.Abs(GetDeriv(tanh)-0.180) > tolerance {
		t.Errorf("Got %v, expected: %v.", GetDeriv(tanh), 0.18)
	}

}

// f is a test function which we are going to calculate the derivative of
func simpleFunction(x *Dual) *Dual {
	return x.Mul(x.Mul(x).Sin()).Add(ConstDual(1))
}

// sigmoid function used in deep learning
func sigmoid(z *Dual) *Dual {
	return ConstDual(1).Div(ConstDual(1).Add(z.Neg().Exp()))
}

// tanh function used in deep learning
func tanh(z *Dual) *Dual {
	return z.Exp().Sub(z.Neg().Exp()).Div(z.Exp().Add(z.Neg().Exp()))
}
