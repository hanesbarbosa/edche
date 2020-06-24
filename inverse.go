package edche

// Inverse gives the inverse of a multivector.
func Inverse(m Multivector) Multivector {
	mc := CloneMultivector(m)
	n := Numerator(m)
	d := Denominator(mc)
	return ScalarDivision(n, d)
}

// Numerator to be defined.
func Numerator(m Multivector) Multivector {
	mc := CloneMultivector(m)
	as := AmplitudeSquared(m)
	asr := Reverse(as)
	cc := CliffordConjugation(mc)

	return GeometricProduct(cc, asr)
}

// Denominator to be defined.
func Denominator(m Multivector) string {
	r := Rationalize(m)
	return r.E0.Num().String()
}
