package calc

import (
	"testing"
)

func TestBasicArithmetic(t *testing.T) {
	tests := []struct {
		expr     string
		expected float64
	}{
		{"1+2", 3},
		{"2*3+4", 10},
		{"2*(3+4)", 14},
		{"2^3", 8},
		{"10/2", 5},
		{"10-3", 7},
	}
	for _, test := range tests {
		root := NewParser(test.expr).Parse()
		result := root.Eval()
		if result != test.expected {
			t.Errorf("%s: got %v, want %v", test.expr, result, test.expected)
		}
	}
}

func TestBrackets(t *testing.T) {
	tests := []struct {
		expr     string
		expected float64
	}{
		{"(2+3)*4", 20},
		{"[2+3]*4", 20},
		{"{2+3}*4", 20},
		{"(2+[3*4])", 14},
		{"{2+[3*(4+1)]}", 17},
	}
	for _, test := range tests {
		root := NewParser(test.expr).Parse()
		result := root.Eval()
		if result != test.expected {
			t.Errorf("%s: got %v, want %v", test.expr, result, test.expected)
		}
	}
}

func TestFunctionsAndConstants(t *testing.T) {
	tests := []struct {
		expr     string
		expected float64
	}{
		{"sin(pi/2)", 1},
		{"cos(0)", 1},
		{"ln(e)", 1},
		{"log(100)", 2},
		{"sqrt(9)", 3},
		{"max(2,5)", 5},
		{"min(2,5)", 2},
		{"gcd(12,18)", 6},
		{"lcm(3,4)", 12},
		{"factorial(5)", 120},
	}
	for _, test := range tests {
		root := NewParser(test.expr).Parse()
		result := root.Eval()
		if asFloat(result) != test.expected {
			t.Errorf("%s: got %v, want %v", test.expr, result, test.expected)
		}
	}
}

func TestFloatParsing(t *testing.T) {
	root := NewParser("1.3+1.4").Parse()
	result := root.Eval()
	if asFloat(result) != 2.7 {
		t.Errorf("1.3+1.4: got %v, want 2.7", result)
	}
}
