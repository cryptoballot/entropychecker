Entropy Checker
===============

Entropy Checker is a handy golang library for Linux that will ensure you have sufficient entropy available before doing important cryptographic operations.

It only works on Linux. Mac, BSD and Windows lack the ability to check entropy levels.

usage example:

```go
package main

import (
	"github.com/cryptoballot/entropychecker"
	"log"
)

func main() {
	// Wait for sufficient entropy to be available
	err := entropychecker.WaitForEntropy()
	if err != nil {
		log.Fatal(err)
	}

	// Now it's safe to do important cryptographic stuff
}
```
