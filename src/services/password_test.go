package services_test

import (
	"testing"

	"github.com/psycomentis/psycofolio++/src/services"
)

func TestVerifyPassword(t *testing.T) {
	pass := "pass@1234"
	p := services.HashPassword(pass)
	err := services.VerifyPassword(p, pass)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("Valid Password")
	}
}
