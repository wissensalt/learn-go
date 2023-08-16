package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "#1 Word Is Palindrome",
			input:    "Racecar",
			expected: true,
		},
		{
			name:     "#2 Word is not Palindrome",
			input:    "NotPalindrome",
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := IsPalindrome(test.input)
			if test.expected != actual {
				t.Errorf("Expected %v, But Actual is %v\n", test.expected, actual)
			}
		})
	}
}
