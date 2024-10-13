package ch1_test

import (
	"testing"

	"github.com/ricksy/learn_go/ch1"
)

func TestMascot(t *testing.T) {
	want := "Gopher"
	if got := ch1.BestMascot(); got != want {
		t.Errorf("BestMascot() = %q, want %q", got, want)
	}
}
