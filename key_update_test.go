package edche

import (
	"strings"
	"testing"
)

// Tests KeyUpdate function
func TestKeyUpdate(t *testing.T) {
	// Old cyphertext
	c1 := New([]string{"-54434609", "54747116", "24365035", "35788815", "28642198", "35931560", "20972281", "-23594743"})

	// Old keys
	// s1 230e0 + -179e1 + 156e2 + -110e3 + 238e12 + -192e13 + 246e23 + -200e123
	s1 := New([]string{"230", "-179", "156", "-110", "238", "-192", "246", "-200"})
	// s2 262e0 + -203e1 + 225e2 + -173e3 + 212e12 + -160e13 + 261e23 + -209e123
	s2 := New([]string{"262", "-203", "225", "-173", "212", "-160", "261", "-209"})

	// New keys
	// s3 407e0 + -332e1 + 511e2 + -439e3 + 361e12 + -289e13 + 526e23 + -454e123
	s3 := New([]string{"407", "-332", "511", "-439", "361", "-289", "526", "-454"})
	// s4 495e0 + -408e1 + 384e2 + -304e3 + 305e12 + -225e13 + 335e23 + -255e123
	s4 := New([]string{"495", "-408", "384", "-304", "305", "-225", "335", "-255"})

	c2 := KeyUpdate(c1, s1, s2, s3, s4)

	if strings.Compare(c2.E0.RatString(), "-98947311") != 0 ||
		strings.Compare(c2.E1.RatString(), "118667702") != 0 ||
		strings.Compare(c2.E2.RatString(), "107948892") != 0 ||
		strings.Compare(c2.E3.RatString(), "96969460") != 0 ||
		strings.Compare(c2.E12.RatString(), "137727185") != 0 ||
		strings.Compare(c2.E13.RatString(), "118794655") != 0 ||
		strings.Compare(c2.E23.RatString(), "120326237") != 0 ||
		strings.Compare(c2.E123.RatString(), "-149318681") != 0 {
		t.Errorf("Wrong updated cyphertext.")
	}
}
