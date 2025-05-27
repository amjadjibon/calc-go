package calc

import "testing"

func TestGCD(t *testing.T) {
	tests := []struct {
		a, b   int
		expect int
	}{
		{12, 18, 6},
		{18, 12, 6},
		{7, 13, 1},
		{0, 5, 5},
		{5, 0, 5},
		{0, 0, 0},
		{-12, 18, 6},
		{12, -18, 6},
	}
	for _, test := range tests {
		got := GCD(test.a, test.b)
		if got != test.expect {
			t.Errorf("GCD(%d, %d) = %d, want %d", test.a, test.b, got, test.expect)
		}
	}
}

func TestLCM(t *testing.T) {
	tests := []struct {
		a, b   int
		expect int
	}{
		{3, 4, 12},
		{0, 5, 0},
		{5, 0, 0},
		{7, 13, 91},
		{12, 18, 36},
	}
	for _, test := range tests {
		got := LCM(test.a, test.b)
		if got != test.expect {
			t.Errorf("LCM(%d, %d) = %d, want %d", test.a, test.b, got, test.expect)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		n      int
		expect int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
	}
	for _, test := range tests {
		got := Factorial(test.n)
		if got != test.expect {
			t.Errorf("Factorial(%d) = %d, want %d", test.n, got, test.expect)
		}
	}
}
