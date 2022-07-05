package services

import (
	"testing"
)

func TestPermitter(t *testing.T) {
	got := Permitter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
