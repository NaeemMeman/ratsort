package ratsort

import (
    "testing"
)

//Testing the Less function
func TestLess(t *testing.T) {
    testQ := [][2]string{
		{"", "a"},
		{"a", ""},
		{"a", "b"},
		{"b", "a"},
		{"a1", "a 1"},
		{"a 1", "a1"},
		{"a1", "b1"},
		{"b1", "a1"},
		{"a 1", "b 1"},
		{"b 1", "a 1"},
		{"a1", "a1.5"},
		{"a1.5", "a1"},
		{"a1.5", "a2"},
		{"a2", "a1.5"},
		{"a 1", "a 1.5"},
		{"a 1.5", "a 1"},
		{"a 1.5", "a 2"},
		{"a 2", "a 1.5"},
		{"1", "2"},
		{"2", "1"},
		{"1", "1a"},
		{"1a", "1"},
	}
	
	testA := []bool{
	    true,
	    false,
	    true,
	    false,
	    true,
	    false,
	    true,
	    false,
	    true,
	    false,
	    true,
	    false,
	    true,
	    false,
	    true,
	    false,
	    true,
	    false,
	    true,
	    false,
	    true,
	    false,
	}
	
	for i, q := range testQ {
	    a := testA[i]
		if Less(q[0], q[1]) != a {
		    t.Errorf("Less(%q, %q) did not return %t.", q[0], q[1], a)
		}
	}
}

//Testing the Sort function
func TestSort(t *testing.T) {
    testQ := []string{"a 3", "a2", "1b", "a3", "b 1", "1", "a 1", "a", "a1", "1a", "a 2", "a 2.1", "", "b1", "a2.2", "2", "a2.12", "a 2.12", "a 2.2", "a2.1", "b"}
    testA := []string{"","1", "1a", "1b", "2", "a", "a1", "a2", "a2.1", "a2.12", "a2.2", "a3", "a 1", "a 2", "a 2.1", "a 2.12", "a 2.2", "a 3", "b", "b1", "b 1"}
    Sort(testQ)
    
    for i, q := range testQ {
	    a := testA[i]
		if q != a {
		    t.Errorf("%q not equal to %q.", q, a)
		}
	}
}
