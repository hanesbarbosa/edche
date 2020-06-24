package edche

import "math/big"

// MultivectorToNumber transforms multivector back into a number.
func MultivectorToNumber(m Multivector) *big.Int {
	sum := big.NewRat(0, 1)
	sum.Add(m.E0, m.E1)
	sum.Add(sum, m.E2)
	sum.Add(sum, m.E3)
	sum.Add(sum, m.E12)
	sum.Add(sum, m.E13)
	sum.Add(sum, m.E23)
	sum.Add(sum, m.E123)
	return sum.Num()
}
