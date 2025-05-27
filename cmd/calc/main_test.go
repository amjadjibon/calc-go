package main

import (
	"os"
	"strings"
	"testing"
)

func TestMainArgsDirect(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()
	tests := []struct {
		args     []string
		expected string
	}{
		{[]string{"calc", "1+2"}, "3"},
		{[]string{"calc", "sin(pi/2)"}, "1"},
		{[]string{"calc", "lcm(3,4)"}, "12"},
		{[]string{"calc", "1.3+1.4"}, "2.7"},
		{[]string{"calc", "factorial(5)"}, "120"},
		{[]string{"calc", "gcd(12,18)"}, "6"},
		{[]string{"calc", "2^3"}, "8"},
		{[]string{"calc", "2*(3+4)"}, "14"},
		{[]string{"calc", "2+3*4"}, "14"},
		{[]string{"calc", "2+3*4/2"}, "8"},
		{[]string{"calc", "2+3*4-5"}, "9"},
	}
	for _, test := range tests {
		os.Args = test.args
		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// call the main function which should read from os.Args
		main()

		_ = w.Close()
		os.Stdout = oldStdout
		out := make([]byte, 1024)
		n, _ := r.Read(out)
		result := string(out[:n])
		if !strings.Contains(result, test.expected) {
			t.Errorf("args %v: got %q, want %q", test.args, result, test.expected)
		}
	}
}
