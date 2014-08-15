package entropychecker

import (
	"testing"
)

func TestEntropyChecker(t *testing.T) {
	err := WaitForEntropy()
	if err != nil {
		t.Error(err)
		return
	}

	entropy, err := GetEntropy()
	if err != nil {
		t.Error(err)
		return
	}

	if entropy < MinimumEntropy {
		t.Error("Insufficient entropy not properly detected")
		return
	}

	Timeout = 0
	err = WaitForEntropy()
	if err != nil {
		t.Error(err)
		return
	}
}
