package edche

import (
	"strings"
	"testing"
)

// Tests if a number can be encrypted and decrypted using functions Encrypt and Decrypt.
func TestEncryptDecrypt(t *testing.T) {
	// Number to be emcrypted converted into multivector
	m := NumberToMultivector("35")
	// First key
	k1 := NumberToMultivector("101")
	// Second key
	k2 := NumberToMultivector("51")
	// Encryption by triple product
	em := Encrypt(m, k1, k2)
	// Decryption by triple product
	dm := Decrypt(em, k1, k2)
	// Conversion from multivector to number
	n := MultivectorToNumber(dm)

	if strings.Compare(n.String(), "35") != 0 {
		t.Errorf("The encryption and/or decryption is not correct!")
	}
}
