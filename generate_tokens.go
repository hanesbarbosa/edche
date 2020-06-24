package edche

// GenerateTokens creates an array of two tokens: [t1, t2]
func GenerateTokens(s1, s2, s3, s4 Multivector) [2]Multivector {
	var tk [2]Multivector
	var s1r, s2r Multivector

	s1r = Inverse(s1)
	s2r = Inverse(s2)

	tk[0] = GeometricProduct(s3, s1r)
	tk[1] = GeometricProduct(s2r, s4)

	return tk
}
