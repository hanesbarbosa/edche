package edche

// ScalarDivision makes the division between a multivector and a scalar.
func ScalarDivision(m Multivector, d string) Multivector {
	m.E0.SetString(m.E0.Num().String() + "/" + d)
	m.E1.SetString(m.E1.Num().String() + "/" + d)
	m.E2.SetString(m.E2.Num().String() + "/" + d)
	m.E3.SetString(m.E3.Num().String() + "/" + d)
	m.E12.SetString(m.E12.Num().String() + "/" + d)
	m.E13.SetString(m.E13.Num().String() + "/" + d)
	m.E23.SetString(m.E23.Num().String() + "/" + d)
	m.E123.SetString(m.E123.Num().String() + "/" + d)

	return m
}
