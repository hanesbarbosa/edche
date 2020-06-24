package edche

import (
	"strings"
	"testing"
)

// Tests if the cloned multivector keeps separated values.
func TestCloneMultivector(t *testing.T) {
	c := []string{"40", "-29", "29", "-25", "33", "-29", "32", "-28"}

	m := New(c)
	mc := CloneMultivector(m)

	if strings.Compare(m.E0.RatString(), mc.E0.RatString()) != 0 ||
		strings.Compare(m.E1.RatString(), mc.E1.RatString()) != 0 ||
		strings.Compare(m.E2.RatString(), mc.E2.RatString()) != 0 ||
		strings.Compare(m.E3.RatString(), mc.E3.RatString()) != 0 ||
		strings.Compare(m.E12.RatString(), mc.E12.RatString()) != 0 ||
		strings.Compare(m.E13.RatString(), mc.E13.RatString()) != 0 ||
		strings.Compare(m.E23.RatString(), mc.E23.RatString()) != 0 ||
		strings.Compare(m.E123.RatString(), mc.E123.RatString()) != 0 {
		t.Errorf("Different results for cloned multivector when it should be equal.")
	}

	m.E0.SetString("1/1")
	m.E1.SetString("1/1")
	m.E2.SetString("1/1")
	m.E3.SetString("1/1")
	m.E12.SetString("1/1")
	m.E13.SetString("1/1")
	m.E23.SetString("1/1")
	m.E123.SetString("1/1")

	if strings.Compare(m.E0.RatString(), mc.E0.RatString()) == 0 ||
		strings.Compare(m.E1.RatString(), mc.E1.RatString()) == 0 ||
		strings.Compare(m.E2.RatString(), mc.E2.RatString()) == 0 ||
		strings.Compare(m.E3.RatString(), mc.E3.RatString()) == 0 ||
		strings.Compare(m.E12.RatString(), mc.E12.RatString()) == 0 ||
		strings.Compare(m.E13.RatString(), mc.E13.RatString()) == 0 ||
		strings.Compare(m.E23.RatString(), mc.E23.RatString()) == 0 ||
		strings.Compare(m.E123.RatString(), mc.E123.RatString()) == 0 {
		t.Errorf("Equal results for cloned multivector when it should be different.")
	}
}
