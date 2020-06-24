package edche

import "math/big"

// NumberToMultivector converts a number, represented as a string, to multivector.
func NumberToMultivector(s string) Multivector {
	var c [8]*big.Int
	q, r, n := new(big.Int), new(big.Int), new(big.Int)
	d := big.NewInt(8)

	for i := 0; i <= 7; i++ {
		c[i] = new(big.Int)
	}

	n.SetString(s, 10)
	q.Div(n, d)
	r.Mod(n, d)

	c[0].Add(q, r)
	for i := 1; i <= 7; i++ {
		c[i] = q
	}

	values := []string{
		c[0].String(),
		c[1].String(),
		c[2].String(),
		c[3].String(),
		c[4].String(),
		c[5].String(),
		c[6].String(),
		c[7].String(),
	}

	return New(values)
}
