package utils

import (
	"testing"
)

func TestSample(t *testing.T) {
	expected := 1
	if expected != 1 {
		t.Errorf("expected %d", expected)
	}
}
