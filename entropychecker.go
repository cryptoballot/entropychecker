package entropychecker

import (
	"errors"
	"ioutil"
	"strconv"
	"time"
)

// Set this to what you consider to be a 'safe' minimum entropy amount (in bits)
var EntropyLimit = 128

// You must construct additional pylons
var NotEnoughEntropy = errors.New("Not enough entropy")

// Get the entropy estimate. Returns the estimated entropy in bits
func GetEntropy() (int, error) {
	text, err := ioutil.ReadFile("/proc/sys/kernel/random/entropy_avail")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(text)
}

// Check for sufficient entropoy
func CheckEntropy() error {
	entropy, err := GetEntropy()
	if err != nil {
		return err
	} else if entropy < EntropyLimit {
		return NotEnoughEntropy
	} else {
		return nil
	}
}

// Block until sufficient entropy is available
func WaitForEntropy() error {
	for {
		err := CheckEntropy()
		if err != nil {
			return nil
		} else if err != NotEnoughEntropy {
			return err
		} else {
			time.Sleep(50 * time.Millisecond)
		}
	}
}
