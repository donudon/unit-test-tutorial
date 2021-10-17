package animal_prefix

import (
	"testing"
)

func TestGetFirstLetterAnimal(t *testing.T) {
	testCases := []struct {
		name        string
		params      string
		expected    string
		expectedErr error
	}{
		{
			name:        "1. D Return Success Case",
			params:      "Dog",
			expected:    "D",
			expectedErr: nil,
		},
		{
			name:        "2. D Return Error Case",
			params:      "",
			expected:    "",
			expectedErr: ErrCode1,
		},
	}
	for _, tc := range testCases {
		got, err := GetFirstLetterAnimal(tc.params)
		if err != tc.expectedErr {
			t.Errorf("[TestGetFirstLetterAnimal] Function failed at %s case, expected error %s, but got %s", tc.name, tc.expectedErr, err)
		}

		if got != tc.expected {
			t.Errorf("[TestGetFirstLetterAnimal] Function failed at %s case, expected return %s, but got %s", tc.name, tc.expected, got)
		}
	}
}
