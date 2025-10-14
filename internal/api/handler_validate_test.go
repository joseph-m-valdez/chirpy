package api

import (
	"testing"
)

func TestNoNaughtyWords(t *testing.T) {
	blocked := "****"
	cases := []struct {
		in, want string
	}{
		{"kerfuffle", blocked },
		{"clean words", "clean words"},
		{"lorem ipsum kerfuffle", "lorem ipsum " + blocked,},
	}

	for _, tc := range cases {
		got := noNaughtyWords(tc.in)
		if got != tc.want {
			t.Errorf(`noNaughtyWords(%#q): got=%q want:%q`, tc.in, got, tc.want)
		}
	}
}

