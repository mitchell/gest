package gest

import "testing"

// TestFunc is the Go standard test function shape.
type TestFunc func(t *testing.T)

// TestFuncGenerator is a function that takes in arbitrary arguments and returns a test function
// with access to them.
type TestFuncGenerator func(args ...interface{}) TestFunc

// Subtest represents a case in a test suite.
type Subtest struct {
	Name      string
	Generator TestFuncGenerator

	message string
}

// Run runs the Subtest.
func (s Subtest) Run(t *testing.T, args ...interface{}) {
	t.Run(s.Name, s.Generator(args...))
}

// Case creates a new test case.
func Case(s *[]Subtest, name string, message string, generator TestFuncGenerator) {
	*s = append(*s, Subtest{Name: name, Generator: generator, message: message})
}

// Test creates a new test suite and returns the cases.
func Test(cases func(s *[]Subtest)) (output []Subtest) {
	cases(&output)

	for idx, subtest := range output {
		var s Subtest
		test := subtest

		s.Generator = func(args ...interface{}) TestFunc {
			return func(t *testing.T) {
				t.Log(test.message)
				test.Generator(args...)(t)
			}
		}
		s.Name = test.Name

		output[idx] = s
	}

	return output
}
