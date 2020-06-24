package edche

// KeyUpdate substitutes the key pair s1 and s2 by s3 and s4 through tokens t1 and t2.
func KeyUpdate(c, s1, s2, s3, s4 Multivector) Multivector {
	tk := GenerateTokens(s1, s2, s3, s4)
	ct2 := GeometricProduct(c, tk[1])
	ct1 := GeometricProduct(tk[0], ct2)

	return ct1
}
