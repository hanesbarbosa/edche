package edche

import "math/big"

// ScalarMultiplication multiplies a multivector by a scalar.
func ScalarMultiplication(m Multivector, scalar *big.Rat) Multivector {
	m.E0.Mul(m.E0, scalar)
	m.E1.Mul(m.E1, scalar)
	m.E2.Mul(m.E2, scalar)
	m.E3.Mul(m.E3, scalar)
	m.E12.Mul(m.E12, scalar)
	m.E13.Mul(m.E13, scalar)
	m.E23.Mul(m.E23, scalar)
	m.E123.Mul(m.E123, scalar)

	return m
}
