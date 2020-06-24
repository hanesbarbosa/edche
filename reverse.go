package edche

import "math/big"

// Reverse of a multivector.
func Reverse(m Multivector) Multivector {
	// pf = positive factor
	pf := new(big.Rat)
	pf.SetString("1/1")

	// nf = negative factor
	nf := new(big.Rat)
	nf.SetString("-1/1")

	m.E0.Mul(pf, m.E0)
	m.E1.Mul(nf, m.E1)
	m.E2.Mul(nf, m.E2)
	m.E3.Mul(nf, m.E3)
	m.E12.Mul(pf, m.E12)
	m.E13.Mul(pf, m.E13)
	m.E23.Mul(pf, m.E23)
	m.E123.Mul(nf, m.E123)

	return m
}
