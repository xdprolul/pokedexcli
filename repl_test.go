package main 

import "testing"

func TestCleanInput(t *testing.T) {
	cases:=[]struct {
		input string
		expected []string	
	}{
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input: "HELLO world",
			expected: []string{"hello", "world"},
		},
	}

	for _,c:=range cases {
		actual:=cleanInput(c.input)
		if len(actual)!=len(c.expected) {
			t.Errorf("test failed, lengths don't match")
		}
		for i:=range actual {
			word:=actual[i]
			expectedWord:=c.expected[i]
			if word!=expectedWord {
				t.Errorf("test failed, word: %v, expectedWord:%v don't match",word,expectedWord)
			}
		}
	}
}
