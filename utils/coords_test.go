package utils

import "testing"

func TestToHumanReadable(t *testing.T) {
	if ToHumanReadable(10) != "A2" {
		t.Errorf("10 Should be A2, got: %s", ToHumanReadable(10))
	}
	if ToHumanReadable(75) != "F8" {
		t.Errorf("75 Should be F8, got: %s", ToHumanReadable(75))
	}
	if ToHumanReadable(82) != "C9" {
		t.Errorf("82 Should be C9, got: %s", ToHumanReadable(82))
	}
	if ToHumanReadable(49) != "J5" {
		t.Errorf("49 Should be J5, got: %s", ToHumanReadable(49))
	}
}

func TestFromHumanReadable(t *testing.T) {
	if FromHumanReadable("A2") != 10 {
		t.Errorf("A2 Should be 10, got: %d", FromHumanReadable("A2"))
	}
	if FromHumanReadable("F8") != 75 {
		t.Errorf("F8 Should be 75, got: %d", FromHumanReadable("F8"))
	}
	if FromHumanReadable("C9") != 82 {
		t.Errorf("C9 Should be 82, got: %d", FromHumanReadable("C9"))
	}
	if FromHumanReadable("J5") != 49 {
		t.Errorf("J5 Should be 49, got: %d", FromHumanReadable("J5"))
	}
}
