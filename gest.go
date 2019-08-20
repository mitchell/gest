package gest

import "testing"

type TestFunc func(t *testing.T)

type TestFuncGenerator func(args ...interface{}) TestFunc

type Subtest struct {
	Name      string
	Generator TestFuncGenerator

	message string
}

func (s Subtest) Run(t *testing.T, args ...interface{}) {
	t.Run(s.Name, s.Generator(args...))
}

func Describe(s *[]Subtest, name string, message string, generator TestFuncGenerator) {
	*s = append(*s, Subtest{Name: name, Generator: generator, message: message})
}

func Test(name string, cases func(s *[]Subtest)) (output []Subtest) {
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

		s.Name = name + "/" + test.Name
		output[idx] = s
	}

	return output
}
