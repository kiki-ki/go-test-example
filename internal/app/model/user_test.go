package model_test

import (
	"testing"

	"github.com/kiki-ki/go-test-example/internal/app/model"
)

func TestUser_IsOverTwentyYearsOld(t *testing.T) {
	cases := map[string]struct {
		in   model.User
		want bool
	}{
		"eqaul": {in: model.User{Age: 20}, want: true},
		"above": {in: model.User{Age: 21}, want: true},
		"below": {in: model.User{Age: 19}, want: false},
	}
	for k, tt := range cases {
		t.Run(k, func(t *testing.T) {
			got := tt.in.IsOverTwentyYearsOld()
			if tt.want != got {
				t.Errorf("want: age(%d) = %v, got: %v", tt.in.Age, tt.want, got)
			}
		})
	}
}
