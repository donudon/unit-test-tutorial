package animal_prefix

import "testing"

func TestGetAnimal(t *testing.T) {
	testCases := []struct {
		name        string
		params      string
		expected    string
		expectedErr error
	}{
		{
			name:        "1. Dog Return Success Case",
			params:      "d",
			expected:    "Dog",
			expectedErr: nil,
		},
	}
	for _, tc := range testCases {
		got, err := GetAnimal(tc.params)
		if err != tc.expectedErr {
			t.Errorf("[TestGetAnimal] Function failed at %s case, expected error %s, but got %s", tc.name, tc.expectedErr, err)
		}

		if got != tc.expected {
			t.Errorf("[TestGetAnimal] Function failed at %s case, expected return %s, but got %s", tc.name, tc.expected, got)
		}
	}
}
