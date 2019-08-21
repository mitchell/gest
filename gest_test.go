package gest

import "testing"

func TestCase(t *testing.T) {
	var ss []Subtest

	Case(&ss, "CaseName", "Case message", func(args ...interface{}) TestFunc {
		return func(t *testing.T) {
			assert(t, args[0] == "zero", "args[0] should equal 'zero'")
			assert(t, args[1] == 1, "args[1] should equal 1")
		}
	})

	s := ss[0]

	s.Run(t, "zero", 1)

	assert(t, s.Name == "CaseName", "s.Name should equal 'CaseName'")
	assert(t, s.message == "Case message", "s.message should equal 'Case message'")
}

func TestTest(t *testing.T) {
	ss := Test(func(s *[]Subtest) {
		Case(s, "Case1", "describe 1", func(args ...interface{}) TestFunc {
			return func(*testing.T) {
				t.Log("Case1 run")
				assert(t, args[0] == "zero", "args[0] should equal 'zero'")
				assert(t, args[1] == 1, "args[1] should equal 1")
			}
		})

		Case(s, "Case2", "describe 2", func(args ...interface{}) TestFunc {
			return func(*testing.T) {
				t.Log("Case2 run")
				assert(t, args[0] == "zero", "args[0] should equal 'zero'")
				assert(t, args[1] == 1, "args[1] should equal 1")
			}
		})
	})

	for _, test := range ss {
		test.Run(t, "zero", 1)
	}

	assert(t, ss[0].Name == "Case1", "first position test should be Case1")
	assert(t, ss[1].Name == "Case2", "second position test should be Case2")
	assert(t, len(ss) == 2, "length of ss should equal 2")
}

func assert(t *testing.T, ok bool, message string) {
	if ok {
		t.Log(message)
	} else {
		t.Error("FAILED: " + message)
	}
}
