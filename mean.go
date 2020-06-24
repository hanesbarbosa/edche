package edche

import "strconv"

// Mean calculates the average amount between multivectors
func Mean(ms []Multivector) Multivector {
	n := len(ms)
	d := strconv.Itoa(n)
	var m0 = []string{"0", "0", "0", "0", "0", "0", "0", "0"}

	tm := New(m0)

	for i := 0; i < n; i++ {
		tm = Addition(tm, ms[i])
	}

	return ScalarDivision(tm, d)
}
