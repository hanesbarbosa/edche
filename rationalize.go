package edche

// Rationalize is a function that rationalizes a multivector.
func Rationalize(m Multivector) Multivector {
	as := AmplitudeSquared(m)
	asc := CloneMultivector(as)

	asr := Reverse(as)
	return GeometricProduct(asc, asr)
}
