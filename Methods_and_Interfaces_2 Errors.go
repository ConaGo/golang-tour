package main

import (
	"fmt"
)
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g",float64(e)) 
}

// Sqrt approximates the square root of x using  Newton's method with a maximum of 10 iterations
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		err := ErrNegativeSqrt(x)
		return 0, err
	}
	z := 1.0
	_z:= 1.0
	for i:=0; i<10; i++ {
		_z -= (z*z - x) / (2*z)
		if -0.00000000000001<_z-z && _z-z<0.0000000000001 {
			return _z, nil
		}
	z =_z
	}	
  return z, nil
}
