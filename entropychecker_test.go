package entropychecker

import (
	"testing"
)

func TestEntropyChecker(t *testing.T) {
	err = WaitForEntropy()
	if err != nil {
		t.Error(err)
	}

	entropy, err := GetEntropy()
	if err != nil {
		t.Error(err)
	}

	if entropy < EntropyLimit {
		t.Error("Insufficient entropy not properly detected")
	}

	Timeout = 0
	err = WaitForEntropy()
	if err != nil {
		t.Error(err)
	}
}
