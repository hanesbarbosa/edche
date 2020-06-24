package edche

// AmplitudeSquared gives the amplitude square of a multivector.
func AmplitudeSquared(m Multivector) Multivector {
	// Clone the multivector
	mc := CloneMultivector(m)
	ccm := CliffordConjugation(m)
	return GeometricProduct(mc, ccm)
}
