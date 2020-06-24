package edche

// CloneMultivector clones a multivector.
func CloneMultivector(m Multivector) Multivector {
	mc := New([]string{
		m.E0.RatString(),
		m.E1.RatString(),
		m.E2.RatString(),
		m.E3.RatString(),
		m.E12.RatString(),
		m.E13.RatString(),
		m.E23.RatString(),
		m.E123.RatString(),
	})
	return mc
}
