package utils

import "testing"

func TestContainsBit(t *testing.T) {
	if !ContainsBit(5, 1) {
		t.Error("5 contains 1")
	}

	if !ContainsBit(3, 1) {
		t.Error("3 contains 1")
	}

	if !ContainsBit(3, 2) {
		t.Error("3 contains 2")
	}
}
