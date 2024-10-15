package fibonacci

import (
	"math/big"
	"sync"
)

type Matrix struct {
	a, b, c, d *big.Int
}

func (m *Matrix) Multiply(n *Matrix) *Matrix {
	var wg sync.WaitGroup
	wg.Add(4)

	newA, newB, newC, newD := new(big.Int), new(big.Int), new(big.Int), new(big.Int)

	go func() {
		defer wg.Done()
		newA.Add(new(big.Int).Mul(m.a, n.a), new(big.Int).Mul(m.b, n.c))
	}()

	go func() {
		defer wg.Done()
		newB.Add(new(big.Int).Mul(m.a, n.b), new(big.Int).Mul(m.b, n.d))
	}()

	go func() {
		defer wg.Done()
		newC.Add(new(big.Int).Mul(m.c, n.a), new(big.Int).Mul(m.d, n.c))
	}()

	go func() {
		defer wg.Done()
		newD.Add(new(big.Int).Mul(m.c, n.b), new(big.Int).Mul(m.d, n.d))
	}()

	wg.Wait()

	m.a.Set(newA)
	m.b.Set(newB)
	m.c.Set(newC)
	m.d.Set(newD)

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
