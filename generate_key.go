package edche

import (
	"crypto/rand"
	"math/big"
)

// GenerateKey function generates a prime and checks
// if its equivalent multivector has an inverse.
func GenerateKey(b int) string {
	p := generatePrime(b)
	m := NumberToMultivector(p.String())
	for !HasInverse(m) {
		m = NumberToMultivector(p.String())
	}
	return m.ToString()
}

func generatePrime(b int) *big.Int {
	p, err := rand.Prime(rand.Reader, b)
	if err != nil {
		panic(err)
	}
	return p
}
