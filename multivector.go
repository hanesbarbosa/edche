package edche

import (
	"math/big"
	"regexp"
)

// Multivector object
type Multivector struct {
	E0   *big.Rat
	E1   *big.Rat
	E2   *big.Rat
	E3   *big.Rat
	E12  *big.Rat
	E13  *big.Rat
	E23  *big.Rat
	E123 *big.Rat
}

// New creates a new multivector from string coefficients
func New(s []string) Multivector {
	var c [8]*big.Rat
	var success bool

	for i := 0; i <= 7; i++ {
		c[i] = new(big.Rat)
		_, success = c[i].SetString(s[i])

		if !success {
			panic("Error converting string to big.Rat.")
		}
	}

	return Multivector{
		E0:   c[0],
		E1:   c[1],
		E2:   c[2],
		E3:   c[3],
		E12:  c[4],
		E13:  c[5],
		E23:  c[6],
		E123: c[7],
	}
}

// ToString converts multivector to a string
func (m *Multivector) ToString() string {
	return m.E0.RatString() + "e0" + "+" +
		m.E1.RatString() + "e1" + "+" +
		m.E2.RatString() + "e2" + "+" +
		m.E3.RatString() + "e3" + "+" +
		m.E12.RatString() + "e12" + "+" +
		m.E13.RatString() + "e13" + "+" +
		m.E23.RatString() + "e23" + "+" +
		m.E123.RatString() + "e123"
}

// StringToMultivector converts strings into multivectors
func StringToMultivector(s string) Multivector {
	c := regexp.MustCompile("e[0-3]+(\\+)?").Split(s, -1)
	return New(c)
}
