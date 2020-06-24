package edche

import "strings"

// HasInverse checks if a multivector (i.e: key) has inverse.
// If so, the decryption can occour.
func HasInverse(k Multivector) bool {
	k1 := CloneMultivector(k)
	ki := Inverse(k1)
	gp := GeometricProduct(k, ki)

	return (strings.Compare(gp.E0.RatString(), "1") == 0 &&
		strings.Compare(gp.E1.RatString(), "0") == 0 &&
		strings.Compare(gp.E2.RatString(), "0") == 0 &&
		strings.Compare(gp.E3.RatString(), "0") == 0 &&
		strings.Compare(gp.E12.RatString(), "0") == 0 &&
		strings.Compare(gp.E13.RatString(), "0") == 0 &&
		strings.Compare(gp.E23.RatString(), "0") == 0 &&
		strings.Compare(gp.E123.RatString(), "0") == 0)
}
