package edche

// Encrypt implements the triple product encryption
func Encrypt(m1, k1, k2 Multivector) Multivector {
	m1 = GeometricProduct(k1, m1)
	m1 = GeometricProduct(m1, k2)
	return m1
}
