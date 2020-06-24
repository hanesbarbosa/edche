package edche

import "testing"

// Tests if the HasInverse function checks properly that the resulting multivector has
// 1 in the first coefficient and all other coefficients are zero,
// after calculating the geometric product with its own inverse.
func TestHasInverse(t *testing.T) {
	k := New([]string{"29", "-18", "23", "-19", "33", "-29", "32", "-28"})
	hi := HasInverse(k)

	if !hi {
		t.Errorf("Multivector should have inverse, but the function reported otherwise.")
	}

	k = New([]string{"100", "100", "100", "100", "100", "100", "100", "100"})
	hi = HasInverse(k)

	if hi {
		t.Errorf("Multivector should not have inverse, but the function reported otherwise.")
	}
}
