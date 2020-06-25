package edche

// KeyUpdate substitutes the key pair s1 and s2 by s3 and s4 through tokens t1 and t2.
func KeyUpdate(c, t1, t2 Multivector) Multivector {
	ct2 := GeometricProduct(c, t2)
	ct1 := GeometricProduct(t1, ct2)

	return ct1
}
