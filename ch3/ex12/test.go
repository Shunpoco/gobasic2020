package main

import "testing"

func TestIsAnagram(t *testing.T) {
	for _, tc := range [...]struct {
		input []string
		want  bool
	}{
		{input: []string{"hoge", "hoge"}, want: true},
	} {
		result := isAnagram(tc.input[0], tc.input[1])
		if result != tc.want {
			t.Errorf("error")
		}
	}
}
