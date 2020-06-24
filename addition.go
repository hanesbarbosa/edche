package edche

// Addition sums two multivectors
func Addition(m1, m2 Multivector) Multivector {
	var m0 = []string{"0", "0", "0", "0", "0", "0", "0", "0"}

	m := New(m0)

	m.E0.Add(m1.E0, m2.E0)
	m.E1.Add(m1.E1, m2.E1)
	m.E2.Add(m1.E2, m2.E2)
	m.E3.Add(m1.E3, m2.E3)
	m.E12.Add(m1.E12, m2.E12)
	m.E13.Add(m1.E13, m2.E13)
	m.E23.Add(m1.E23, m2.E23)
	m.E123.Add(m1.E123, m2.E123)

	return m
}
