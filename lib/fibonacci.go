package fiblib

import (
	"math/big"
)

// Matrix structure to hold the 2x2 matrix of big.Int numbers
type Matrix struct {
	a, b, c, d *big.Int
}

//// Function to multiply two 2x2 matrices of big.Int
//func (m *Matrix) Multiply(n *Matrix) *Matrix {
//	return &Matrix{
//		a: new(big.Int).Add(new(big.Int).Mul(m.a, n.a), new(big.Int).Mul(m.b, n.c)),
//		b: new(big.Int).Add(new(big.Int).Mul(m.a, n.b), new(big.Int).Mul(m.b, n.d)),
//		c: new(big.Int).Add(new(big.Int).Mul(m.c, n.a), new(big.Int).Mul(m.d, n.c)),
//		d: new(big.Int).Add(new(big.Int).Mul(m.c, n.b), new(big.Int).Mul(m.d, n.d)),
//	}
//}
// Function to multiply two 2x2 matrices of big.Int (in-place)
func (m *Matrix) Multiply(n *Matrix) *Matrix {
	// Temporary variables to store the result to avoid allocating new big.Int objects
	a := new(big.Int).Add(new(big.Int).Mul(m.a, n.a), new(big.Int).Mul(m.b, n.c))
	b := new(big.Int).Add(new(big.Int).Mul(m.a, n.b), new(big.Int).Mul(m.b, n.d))
	c := new(big.Int).Add(new(big.Int).Mul(m.c, n.a), new(big.Int).Mul(m.d, n.c))
	d := new(big.Int).Add(new(big.Int).Mul(m.c, n.b), new(big.Int).Mul(m.d, n.d))

	// Reuse existing matrix fields (avoid allocating new objects)
	m.a.Set(a)
	m.b.Set(b)
	m.c.Set(c)
	m.d.Set(d)

	return m
}

// Function to compute the matrix raised to the power n using binary exponentiation
func (m *Matrix) Power(n *big.Int) *Matrix {
	// Identity matrix
	result := &Matrix{
		a: big.NewInt(1),
		b: big.NewInt(0),
		c: big.NewInt(0),
		d: big.NewInt(1),
	}
	base := m

	// Temporary variables
	zero := big.NewInt(0)
	one := big.NewInt(1)
	exponent := new(big.Int).Set(n)

	// Binary exponentiation
	for exponent.Cmp(zero) > 0 {
		if new(big.Int).And(exponent, one).Cmp(one) == 0 {
			result = result.Multiply(base)
		}
		base = base.Multiply(base)
		exponent.Rsh(exponent, 1)
	}

	return result
}

// Function to compute the nth Fibonacci number using matrix exponentiation
func FibonacciMatrix(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	if n.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}

	// Fibonacci transformation matrix
	baseMatrix := &Matrix{
		a: big.NewInt(1),
		b: big.NewInt(1),
		c: big.NewInt(1),
		d: big.NewInt(0),
	}

	// Compute baseMatrix^(n-1)
	resultMatrix := baseMatrix.Power(new(big.Int).Sub(n, big.NewInt(1)))

	// F(n) is stored in resultMatrix.a
	return resultMatrix.a
}

