package main

import "testing"


func TestCleanInput(t *testing.T) {
	// test cases
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: " hello world ",
			expected: []string {"hello", "world"},
		},
		{
			input: "    HEllO WoRlD    ",
			expected: []string {"hello", "world"},
		},
		{
			input: "Someone said hello World",
			expected: []string {"someone", "said", "hello", "world"},
		},
		{
			input: "",
			expected: []string {},
		},
	}

	// loop over the tests
	for _, test := range cases {
		actual := cleanInput(test.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(test.expected) {
			t.Errorf("Wrong output length: %v", len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := test.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Unexpected word: %v", word)
				continue
			}
		}
	}
}