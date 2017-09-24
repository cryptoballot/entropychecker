package entropychecker

import (
	"errors"
	"io/ioutil"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// MinimumEntropy is the minimum amount of entropy that will be considered safe.
// Set this to what you consider to be a 'safe' minimum entropy amount (in bits)
var MinimumEntropy = 128

// Timeout sets the maximum amount of time to wait for entropy.
// Waiting for entropy will time out after this amount of time. Setting to zero will never time out.
var Timeout = time.Second * 10

// ErrTimeout is for when the system waits too long and gives up
var ErrTimeout = errors.New("entropychecker: Timed out waiting for sufficient entropy")

// ErrUnsupportedOS is for for an invalid OS that does not provide entropy estimates
var ErrUnsupportedOS = errors.New("entropychecker: Unsupported OS. Only Linux is supported")

// GetEntropy gets the entropy estimate. Returns the estimated entropy in bits
func GetEntropy() (int, error) {
	if runtime.GOOS != "linux" {
		return 0, ErrUnsupportedOS
	}

	text, err := ioutil.ReadFile("/proc/sys/kernel/random/entropy_avail")
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(strings.TrimSuffix(string(text), "\n"))
}

// WaitForEntropy blocks until sufficient entropy is available
func WaitForEntropy() error {
	if runtime.GOOS != "linux" {
		return ErrUnsupportedOS
	}

	// set up the timeout
	timeout := make(chan bool, 1)
	if Timeout != 0 {
		go func(timeoutDuration time.Duration) {
			time.Sleep(timeoutDuration)
			timeout <- true
		}(Timeout)
	}

	for {
		entropy, err := GetEntropy()

		switch {
		case err != nil:
			return err
		case entropy > MinimumEntropy:
			return nil
		default:
			select {
			case <-timeout:
				return ErrTimeout
			default:
				time.Sleep(50 * time.Millisecond)
			}
		}
	}
}
