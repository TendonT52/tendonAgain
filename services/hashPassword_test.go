package services

import "testing"

func TestHashAndValidatePassword(t *testing.T) {
	tc := []struct {
		password string
		check    string
		expect   bool
	}{
		// {
		// 	"test",
		// 	"test",
		// 	true,
		// },
		// {
		// 	"test",
		// 	"wrong",
		// 	false,
		// },
		{
			"IS10Ucv1mMJsptb",
			"$2a$14$5iRXiBBSbDteBnzgRbAfXOK/B6y0OO16wEZ1vKHS5VHyHNEyNLCBy",
			true,
		},
	}

	for _, tc := range tc {
		// hashPassword := HashPassword(tc.password)
		valid := CheckPasswordHash(tc.password, tc.check)
		if valid != tc.expect {
			t.Errorf("wrong validate password got %v, expect %v", valid, tc.expect)
			return
		}
	}
}
