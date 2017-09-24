Entropy Checker
===============
[![Build Status](https://scrutinizer-ci.com/g/cryptoballot/entropychecker/badges/build.png?b=master)](https://scrutinizer-ci.com/g/cryptoballot/entropychecker/build-status/master)
[![Build Status](https://travis-ci.org/cryptoballot/entropychecker.svg?branch=master)](https://travis-ci.org/cryptoballot/entropychecker)
[![Go Report Card](https://goreportcard.com/badge/github.com/cryptoballot/entropychecker)](https://goreportcard.com/report/github.com/cryptoballot/entropychecker)
[![Scrutinizer Issues](https://img.shields.io/badge/scrutinizer-issues-blue.svg)](https://scrutinizer-ci.com/g/cryptoballot/entropychecker/issues)
[![GoDoc](https://godoc.org/github.com/cryptoballot/entropychecker?status.svg)](https://godoc.org/github.com/cryptoballot/entropychecker)


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

There are two configuration variables provided:

```go

// By default we wait for 128 bits, but if you need more or less you can change it here
entropychecker.MinimumEntropy = 128

// By default we will wait 10 seconds before timing out, but we can set it differently.
// Set it to 0 to never time out
entropychecker.Timeout = 10 * time.Second
```
