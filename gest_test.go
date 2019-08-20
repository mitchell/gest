package gest

import "testing"

func TestDescribe(t *testing.T) {
	var ss []Subtest

	Describe(&ss, "DescribeName", "Describe message", func(args ...interface{}) TestFunc {
		return func(t *testing.T) {
			assert(t, args[0] == "zero", "args[0] should equal 'zero'")
			assert(t, args[1] == 1, "args[1] should equal 1")
		}
	})

	s := ss[0]

	s.Run(t, "zero", 1)

	assert(t, s.Name == "DescribeName", "s.Name should equal 'DescribeName'")
	assert(t, s.message == "Describe message", "s.message should equal 'Describe message'")
}

func TestTest(t *testing.T) {
	ss := Test("TestName", func(s *[]Subtest) {
		Describe(s, "Describe1", "describe 1", func(args ...interface{}) TestFunc {
			return func(*testing.T) {
				t.Log("Describe1 run")
			}
		})

		Describe(s, "Describe2", "describe 2", func(args ...interface{}) TestFunc {
			return func(*testing.T) {
				t.Log("Describe2 run")
			}
		})
	})

	for _, test := range ss {
		test.Run(t)
	}

	assert(t, ss[0].Name == "TestName/Describe1", "first position test should be Desribe1")
	assert(t, ss[1].Name == "TestName/Describe2", "second position test should be Desribe2")
	assert(t, len(ss) == 2, "length of ss should equal 2")
}

func assert(t *testing.T, ok bool, message string) {
	if ok {
		t.Log(message)
	} else {
		t.Error("FAILED: " + message)
	}
}
