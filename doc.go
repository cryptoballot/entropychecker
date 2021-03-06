// Package entropychecker is a handy golang library for Linux that will ensure you have sufficient entropy available before doing important cryptographic operations.
//
// It only works on Linux. Mac, BSD and Windows lack the ability to check entropy levels.
//
// usage example:
//
//   package main
//
//   import (
//   	"github.com/cryptoballot/entropychecker"
//   	"log"
//   )
//
//   func main() {
//   	// Wait for sufficient entropy to be available
//   	err := entropychecker.WaitForEntropy()
//   	if err != nil {
//   		log.Fatal(err)
//   	}
//
//   	// Now it's safe to do important cryptographic stuff
//   }
//
// There are two configuration variables provided:
//
//
//   // By default we wait for 128 bits, but if you need more or less you can change it here
//   entropychecker.MinimumEntropy = 128
//
//   // By default we will wait 10 seconds before timing out, but we can set it differently.
//   // Set it to 0 to never time out
//   entropychecker.Timeout = 10 * time.Second
//
//
package entropychecker
