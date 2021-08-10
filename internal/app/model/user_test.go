package model_test

import (
	"testing"

	"github.com/kiki-ki/go-test-example/internal/app/model"
)

func TestIsOverTwentyYearsOld(t *testing.T) {
	cases := map[string]struct {
		in   int
		want bool
	}{
		"eqaul": {in: 20, want: true},
		"above": {in: 21, want: true},
		"below": {in: 19, want: false},
	}
	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			u := new(model.User)
			u.Age = tt.in
			if got := u.IsOverTwentyYearsOld(); tt.want != got {
				t.Errorf("want: age(%d) = %v, got: %v", tt.in, tt.want, got)
			}
		})
	}
}
