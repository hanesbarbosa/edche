package edche

// Decrypt implements the triple product decryption
func Decrypt(c, k1, k2 Multivector) Multivector {
	k2i := Inverse(k2)
	k1i := Inverse(k1)
	cgp := GeometricProduct(k1i, c)
	return GeometricProduct(cgp, k2i)
}
