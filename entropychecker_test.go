package entropychecker

import (
	"testing"
	"time"
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

func TestFailure(t *testing.T) {
	Timeout = 200 * time.Millisecond
	MinimumEntropy = 100000
	err := WaitForEntropy()
	if err == nil {
		t.Error("Should get error when timeout waited too long")
	}
}
