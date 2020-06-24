package edche

import (
	"strings"
	"testing"
)

// Tests GenerateTokens function
func TestGenerateTokens(t *testing.T) {
	// Keys
	// s1 230e0 + -179e1 + 156e2 + -110e3 + 238e12 + -192e13 + 246e23 + -200e123
	s1 := New([]string{"230", "-179", "156", "-110", "238", "-192", "246", "-200"})
	// s2 262e0 + -203e1 + 225e2 + -173e3 + 212e12 + -160e13 + 261e23 + -209e123
	s2 := New([]string{"262", "-203", "225", "-173", "212", "-160", "261", "-209"})

	// New keys
	// s3 407e0 + -332e1 + 511e2 + -439e3 + 361e12 + -289e13 + 526e23 + -454e123
	s3 := New([]string{"407", "-332", "511", "-439", "361", "-289", "526", "-454"})
	// s4 495e0 + -408e1 + 384e2 + -304e3 + 305e12 + -225e13 + 335e23 + -255e123
	s4 := New([]string{"495", "-408", "384", "-304", "305", "-225", "335", "-255"})

	// Generated tokens
	// t1
	// 8587226614/9823510385e0 + 164486879/151130929e1 + 304740719/9823510385e2 +
	// -2941114845/1964702077e3 + -58559965/1964702077e12 + 12087626263/9823510385e13 +
	// 8716848/151130929e23 + 103152622/9823510385e123
	// t2
	// 1765938562/1757509429e0 + -81135545/135193033e1 + -14662422/135193033e2 +
	// 2548676726/1757509429e3 + -467852232/1757509429e12 + 196424304/135193033e13 +
	// -19921704/135193033e23 + 486323692/1757509429e123
	tk := GenerateTokens(s1, s2, s3, s4)

	if strings.Compare(tk[0].E0.RatString(), "8587226614/9823510385") != 0 ||
		strings.Compare(tk[0].E1.RatString(), "164486879/151130929") != 0 ||
		strings.Compare(tk[0].E2.RatString(), "304740719/9823510385") != 0 ||
		strings.Compare(tk[0].E3.RatString(), "-2941114845/1964702077") != 0 ||
		strings.Compare(tk[0].E12.RatString(), "-58559965/1964702077") != 0 ||
		strings.Compare(tk[0].E13.RatString(), "12087626263/9823510385") != 0 ||
		strings.Compare(tk[0].E23.RatString(), "8716848/151130929") != 0 ||
		strings.Compare(tk[0].E123.RatString(), "103152622/9823510385") != 0 {
		t.Errorf("Wrong results for t1.")
	}

	if strings.Compare(tk[1].E0.RatString(), "1765938562/1757509429") != 0 ||
		strings.Compare(tk[1].E1.RatString(), "-81135545/135193033") != 0 ||
		strings.Compare(tk[1].E2.RatString(), "-14662422/135193033") != 0 ||
		strings.Compare(tk[1].E3.RatString(), "2548676726/1757509429") != 0 ||
		strings.Compare(tk[1].E12.RatString(), "-467852232/1757509429") != 0 ||
		strings.Compare(tk[1].E13.RatString(), "196424304/135193033") != 0 ||
		strings.Compare(tk[1].E23.RatString(), "-19921704/135193033") != 0 ||
		strings.Compare(tk[1].E123.RatString(), "486323692/1757509429") != 0 {
		t.Errorf("Wrong results for t2.")
	}
}
