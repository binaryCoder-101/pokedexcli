package main

import (
    "testing"
)

func TestCleanInput(t *testing.T) {
 cases := []struct {
	 input string
	 expected []string
 }{
   {
	   input: " hello world ",
	   expected: []string{"hello", "world"},
   },
   {
	   input: " i do not know",
	   expected: []string{"i", "do", "not", "know"},
   },
   {
	   input: "boot dot  dev",
	   expected: []string{"boot", "dot", "dev"},
   },
 }
 
 for _, c := range cases {
	 actualSlice := cleanInput(c.input)
	 expectedSlice := c.expected
	 if len(actualSlice) != len(expectedSlice){
	 	t.Errorf("wrong length")
		continue
	 }
     for i := range actualSlice {
	word := actualSlice[i]
	expectedWord := expectedSlice[i]
	if expectedWord != word {
		t.Errorf("Wrong word")
	}
     } 	 
 }

}
