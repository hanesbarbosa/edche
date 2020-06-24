package edche

import "math/big"

// GeometricProduct yields the geometric product of two multivectors
func GeometricProduct(m1, m2 Multivector) Multivector {
	// Initialized coefficients
	var m0 = []string{"0", "0", "0", "0", "0", "0", "0", "0"}

	m := New(m0)

	c := [8]*big.Rat{
		new(big.Rat),
		new(big.Rat),
		new(big.Rat),
		new(big.Rat),
		new(big.Rat),
		new(big.Rat),
		new(big.Rat),
		new(big.Rat),
	}

	// 1st Coefficient
	//m.E0 = (m1.E0 * m2.E0) + (m1.E1 * m2.E1) + (m1.E2 * m2.E2) + (m1.E3 * m2.E3) - (m1.E12 * m2.E12) - (m1.E13 * m2.E13) - (m1.E23 * m2.E23) - (m1.E123 * m2.E123)

	//(m1.E0 * m2.E0)
	c[0].Mul(m1.E0, m2.E0)
	//(m1.E1 * m2.E1)
	c[1].Mul(m1.E1, m2.E1)
	//(m1.E2 * m2.E2)
	c[2].Mul(m1.E2, m2.E2)
	//(m1.E3 * m2.E3)
	c[3].Mul(m1.E3, m2.E3)
	//(m1.E12 * m2.E12)
	c[4].Mul(m1.E12, m2.E12)
	//(m1.E13 * m2.E13)
	c[5].Mul(m1.E13, m2.E13)
	//(m1.E23 * m2.E23)
	c[6].Mul(m1.E23, m2.E23)
	//(m1.E123 * m2.E123)
	c[7].Mul(m1.E123, m2.E123)

	//m.E0 = c[0] + c[1] + c[2] + c[3] - c[4] - c[5] - c[6] - c[7]
	m.E0.Add(c[0], c[1])
	m.E0.Add(m.E0, c[2])
	m.E0.Add(m.E0, c[3])
	m.E0.Sub(m.E0, c[4])
	m.E0.Sub(m.E0, c[5])
	m.E0.Sub(m.E0, c[6])
	m.E0.Sub(m.E0, c[7])

	// 2nd Coefficient
	//m.E1 = (m1.E0 * m2.E1) + (m1.E1 * m2.E0) - (m1.E2 * m2.E12) - (m1.E3 * m2.E13) + (m1.E12 * m2.E2) + (m1.E13 * m2.E3) - (m1.E23 * m2.E123) - (m1.E123 * m2.E23)

	//(m1.E0 * m2.E1)
	c[0].Mul(m1.E0, m2.E1)
	//(m1.E1 * m2.E0)
	c[1].Mul(m1.E1, m2.E0)
	//(m1.E2 * m2.E12)
	c[2].Mul(m1.E2, m2.E12)
	//(m1.E3 * m2.E13)
	c[3].Mul(m1.E3, m2.E13)
	//(m1.E12 * m2.E2)
	c[4].Mul(m1.E12, m2.E2)
	//(m1.E13 * m2.E3)
	c[5].Mul(m1.E13, m2.E3)
	//(m1.E23 * m2.E123)
	c[6].Mul(m1.E23, m2.E123)
	//(m1.E123 * m2.E23)
	c[7].Mul(m1.E123, m2.E23)

	//m.E1 = + - - + + - -
	m.E1.Add(c[0], c[1])
	m.E1.Sub(m.E1, c[2])
	m.E1.Sub(m.E1, c[3])
	m.E1.Add(m.E1, c[4])
	m.E1.Add(m.E1, c[5])
	m.E1.Sub(m.E1, c[6])
	m.E1.Sub(m.E1, c[7])

	// 3rd Coefficient
	//m.E2 = (m1.E0 * m2.E2) + (m1.E1 * m2.E12) + (m1.E2 * m2.E0) - (m1.E3 * m2.E23) - (m1.E12 * m2.E1) + (m1.E13 * m2.E123) + (m1.E23 * m2.E3) + (m1.E123 * m2.E13)

	//(m1.E0 * m2.E2)
	c[0].Mul(m1.E0, m2.E2)
	//(m1.E1 * m2.E12)
	c[1].Mul(m1.E1, m2.E12)
	//(m1.E2 * m2.E0)
	c[2].Mul(m1.E2, m2.E0)
	//(m1.E3 * m2.E23)
	c[3].Mul(m1.E3, m2.E23)
	//(m1.E12 * m2.E1)
	c[4].Mul(m1.E12, m2.E1)
	//(m1.E13 * m2.E123)
	c[5].Mul(m1.E13, m2.E123)
	//(m1.E23 * m2.E3)
	c[6].Mul(m1.E23, m2.E3)
	//(m1.E123 * m2.E13)
	c[7].Mul(m1.E123, m2.E13)

	//m.E2 = + + - - + + +
	m.E2.Add(c[0], c[1])
	m.E2.Add(m.E2, c[2])
	m.E2.Sub(m.E2, c[3])
	m.E2.Sub(m.E2, c[4])
	m.E2.Add(m.E2, c[5])
	m.E2.Add(m.E2, c[6])
	m.E2.Add(m.E2, c[7])

	// 4th Coefficient
	//m.E3 = (m1.E0 * m2.E3) + (m1.E1 * m2.E13) + (m1.E2 * m2.E23) + (m1.E3 * m2.E0) - (m1.E12 * m2.E123) - (m1.E13 * m2.E1) - (m1.E23 * m2.E2) - (m1.E123 * m2.E12)

	//(m1.E0 * m2.E3)
	c[0].Mul(m1.E0, m2.E3)
	//(m1.E1 * m2.E13)
	c[1].Mul(m1.E1, m2.E13)
	//(m1.E2 * m2.E23)
	c[2].Mul(m1.E2, m2.E23)
	//(m1.E3 * m2.E0)
	c[3].Mul(m1.E3, m2.E0)
	//(m1.E12 * m2.E123)
	c[4].Mul(m1.E12, m2.E123)
	//(m1.E13 * m2.E1)
	c[5].Mul(m1.E13, m2.E1)
	//(m1.E23 * m2.E2)
	c[6].Mul(m1.E23, m2.E2)
	//(m1.E123 * m2.E12)
	c[7].Mul(m1.E123, m2.E12)

	//m.E3 = + + + - - - -
	m.E3.Add(c[0], c[1])
	m.E3.Add(m.E3, c[2])
	m.E3.Add(m.E3, c[3])
	m.E3.Sub(m.E3, c[4])
	m.E3.Sub(m.E3, c[5])
	m.E3.Sub(m.E3, c[6])
	m.E3.Sub(m.E3, c[7])

	// 5th Coefficient
	//m.E12 = (m1.E0 * m2.E12) + (m1.E1 * m2.E2) - (m1.E2 * m2.E1) + (m1.E3 * m2.E123) + (m1.E12 * m2.E0) - (m1.E13 * m2.E23) + (m1.E23 * m2.E13) + (m1.E123 * m2.E3)

	//(m1.E0 * m2.E12)
	c[0].Mul(m1.E0, m2.E12)
	//(m1.E1 * m2.E2)
	c[1].Mul(m1.E1, m2.E2)
	//(m1.E2 * m2.E1)
	c[2].Mul(m1.E2, m2.E1)
	//(m1.E3 * m2.E123)
	c[3].Mul(m1.E3, m2.E123)
	//(m1.E12 * m2.E0)
	c[4].Mul(m1.E12, m2.E0)
	//(m1.E13 * m2.E23)
	c[5].Mul(m1.E13, m2.E23)
	//(m1.E23 * m2.E13)
	c[6].Mul(m1.E23, m2.E13)
	//(m1.E123 * m2.E3)
	c[7].Mul(m1.E123, m2.E3)

	//m.E12 = + - + + - + +
	m.E12.Add(c[0], c[1])
	m.E12.Sub(m.E12, c[2])
	m.E12.Add(m.E12, c[3])
	m.E12.Add(m.E12, c[4])
	m.E12.Sub(m.E12, c[5])
	m.E12.Add(m.E12, c[6])
	m.E12.Add(m.E12, c[7])

	// 6th Coefficient
	//m.E13 = (m1.E0 * m2.E13) + (m1.E1 * m2.E3) - (m1.E2 * m2.E123) - (m1.E3 * m2.E1) + (m1.E12 * m2.E23) + (m1.E13 * m2.E0) - (m1.E23 * m2.E12) - (m1.E123 * m2.E2)

	//(m1.E0 * m2.E13)
	c[0].Mul(m1.E0, m2.E13)
	//(m1.E1 * m2.E3)
	c[1].Mul(m1.E1, m2.E3)
	//(m1.E2 * m2.E123)
	c[2].Mul(m1.E2, m2.E123)
	//(m1.E3 * m2.E1)
	c[3].Mul(m1.E3, m2.E1)
	//(m1.E12 * m2.E23)
	c[4].Mul(m1.E12, m2.E23)
	//(m1.E13 * m2.E0)
	c[5].Mul(m1.E13, m2.E0)
	//(m1.E23 * m2.E12)
	c[6].Mul(m1.E23, m2.E12)
	//(m1.E123 * m2.E2)
	c[7].Mul(m1.E123, m2.E2)

	//m.E13 = + - - + + - -
	m.E13.Add(c[0], c[1])
	m.E13.Sub(m.E13, c[2])
	m.E13.Sub(m.E13, c[3])
	m.E13.Add(m.E13, c[4])
	m.E13.Add(m.E13, c[5])
	m.E13.Sub(m.E13, c[6])
	m.E13.Sub(m.E13, c[7])

	// 7th Coefficient
	//m.E23 = (m1.E0 * m2.E23) + (m1.E1 * m2.E123) + (m1.E2 * m2.E3) - (m1.E3 * m2.E2) - (m1.E12 * m2.E13) + (m1.E13 * m2.E12) + (m1.E23 * m2.E0) + (m1.E123 * m2.E1)

	//(m1.E0 * m2.E23)
	c[0].Mul(m1.E0, m2.E23)
	//(m1.E1 * m2.E123)
	c[1].Mul(m1.E1, m2.E123)
	//(m1.E2 * m2.E3)
	c[2].Mul(m1.E2, m2.E3)
	//(m1.E3 * m2.E2)
	c[3].Mul(m1.E3, m2.E2)
	//(m1.E12 * m2.E13)
	c[4].Mul(m1.E12, m2.E13)
	//(m1.E13 * m2.E12)
	c[5].Mul(m1.E13, m2.E12)
	//(m1.E23 * m2.E0)
	c[6].Mul(m1.E23, m2.E0)
	//(m1.E123 * m2.E1)
	c[7].Mul(m1.E123, m2.E1)

	//m.E23 = + + - - + + +
	m.E23.Add(c[0], c[1])
	m.E23.Add(m.E23, c[2])
	m.E23.Sub(m.E23, c[3])
	m.E23.Sub(m.E23, c[4])
	m.E23.Add(m.E23, c[5])
	m.E23.Add(m.E23, c[6])
	m.E23.Add(m.E23, c[7])

	// 8th Coefficient
	//m.E123 = (m1.E0 * m2.E123) + (m1.E1 * m2.E23) - (m1.E2 * m2.E13) + (m1.E3 * m2.E12) + (m1.E12 * m2.E3) - (m1.E13 * m2.E2) + (m1.E23 * m2.E1) + (m1.E123 * m2.E0)

	//(m1.E0 * m2.E123)
	c[0].Mul(m1.E0, m2.E123)
	//(m1.E1 * m2.E23)
	c[1].Mul(m1.E1, m2.E23)
	//(m1.E2 * m2.E13)
	c[2].Mul(m1.E2, m2.E13)
	//(m1.E3 * m2.E12)
	c[3].Mul(m1.E3, m2.E12)
	//(m1.E12 * m2.E3)
	c[4].Mul(m1.E12, m2.E3)
	//(m1.E13 * m2.E2)
	c[5].Mul(m1.E13, m2.E2)
	//(m1.E23 * m2.E1)
	c[6].Mul(m1.E23, m2.E1)
	//(m1.E123 * m2.E0)
	c[7].Mul(m1.E123, m2.E0)

	//m.E123 = + - + + - + +
	m.E123.Add(c[0], c[1])
	m.E123.Sub(m.E123, c[2])
	m.E123.Add(m.E123, c[3])
	m.E123.Add(m.E123, c[4])
	m.E123.Sub(m.E123, c[5])
	m.E123.Add(m.E123, c[6])
	m.E123.Add(m.E123, c[7])

	return m
}
