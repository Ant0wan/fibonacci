package fibgmp

import (
	"github.com/ncw/gmp"
)

// Matrix structure with GMP integers
type Matrix struct {
	a, b, c, d *gmp.Int
}

// Multiply two 2x2 matrices using GMP big integers
func (m *Matrix) Multiply(n *Matrix) *Matrix {
	a := new(gmp.Int).Add(new(gmp.Int).Mul(m.a, n.a), new(gmp.Int).Mul(m.b, n.c))
	b := new(gmp.Int).Add(new(gmp.Int).Mul(m.a, n.b), new(gmp.Int).Mul(m.b, n.d))
	c := new(gmp.Int).Add(new(gmp.Int).Mul(m.c, n.a), new(gmp.Int).Mul(m.d, n.c))
	d := new(gmp.Int).Add(new(gmp.Int).Mul(m.c, n.b), new(gmp.Int).Mul(m.d, n.d))

	m.a.Set(a)
	m.b.Set(b)
	m.c.Set(c)
	m.d.Set(d)

	return m
}

// Exponentiate the matrix using GMP big integers
func (m *Matrix) Power(n *gmp.Int) *Matrix {
	one := gmp.NewInt(1)
	zero := gmp.NewInt(0)

	// Identity matrix
	result := &Matrix{
		a: gmp.NewInt(1),
		b: gmp.NewInt(0),
		c: gmp.NewInt(0),
		d: gmp.NewInt(1),
	}
	base := &Matrix{
		a: new(gmp.Int).Set(m.a),
		b: new(gmp.Int).Set(m.b),
		c: new(gmp.Int).Set(m.c),
		d: new(gmp.Int).Set(m.d),
	}

	for n.Cmp(zero) > 0 {
		if new(gmp.Int).And(n, one).Cmp(one) == 0 {
			result = result.Multiply(base)
		}
		base = base.Multiply(base)
		n.Rsh(n, 1)
	}

	return result
}

// Fibonacci function using matrix exponentiation with GMP
func FibonacciGMP(n int64) *gmp.Int {
	if n == 0 {
		return gmp.NewInt(0)
	}
	if n == 1 {
		return gmp.NewInt(1)
	}

	// Base Fibonacci transformation matrix
	baseMatrix := &Matrix{
		a: gmp.NewInt(1),
		b: gmp.NewInt(1),
		c: gmp.NewInt(1),
		d: gmp.NewInt(0),
	}

	// Matrix exponentiation to get Fibonacci(n)
	resultMatrix := baseMatrix.Power(gmp.NewInt(n - 1))

	return resultMatrix.a
}

