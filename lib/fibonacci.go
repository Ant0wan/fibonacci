package fiblib

import (
	"math/big"
)

type Matrix struct {
	a, b, c, d *big.Int
}

func (m *Matrix) Multiply(n *Matrix) *Matrix {
	a := new(big.Int).Add(new(big.Int).Mul(m.a, n.a), new(big.Int).Mul(m.b, n.c))
	b := new(big.Int).Add(new(big.Int).Mul(m.a, n.b), new(big.Int).Mul(m.b, n.d))
	c := new(big.Int).Add(new(big.Int).Mul(m.c, n.a), new(big.Int).Mul(m.d, n.c))
	d := new(big.Int).Add(new(big.Int).Mul(m.c, n.b), new(big.Int).Mul(m.d, n.d))

	m.a.Set(a)
	m.b.Set(b)
	m.c.Set(c)
	m.d.Set(d)

	return m
}

func (m *Matrix) Power(n *big.Int) *Matrix {
	result := &Matrix{
		a: big.NewInt(1),
		b: big.NewInt(0),
		c: big.NewInt(0),
		d: big.NewInt(1),
	}
	base := &Matrix{
		a: new(big.Int).Set(m.a),
		b: new(big.Int).Set(m.b),
		c: new(big.Int).Set(m.c),
		d: new(big.Int).Set(m.d),
	}

	zero := big.NewInt(0)
	one := big.NewInt(1)

	for n.Cmp(zero) > 0 {
		if new(big.Int).And(n, one).Cmp(one) == 0 {
			result = result.Multiply(base)
		}
		base = base.Multiply(base)
		n.Rsh(n, 1)
	}

	return result
}


func FibonacciMatrix(n *big.Int) *big.Int {
	if n.Cmp(big.NewInt(0)) == 0 {
		return big.NewInt(0)
	}
	if n.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(1)
	}

	baseMatrix := &Matrix{
		a: big.NewInt(1),
		b: big.NewInt(1),
		c: big.NewInt(1),
		d: big.NewInt(0),
	}

	resultMatrix := baseMatrix.Power(new(big.Int).Sub(n, big.NewInt(1)))

	return resultMatrix.a
}

