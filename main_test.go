package main

import (
	"fmt"
	"testing"
	"time"
)

// testConcurrencySpeed test if my concurrency implementation
// made a difference. (IT DID!!)
// ignoring test because it has http requirements on domain
// I don't want to bother
func testConcurrencySpeed(t *testing.T) {

	// time non-concurrent
	start := time.Now()
	displayAptsNoChannels()
	elapsed := time.Since(start)

	// time concurrent
	start = time.Now()
	displayAptsChannels()
	elapsedWithCh := time.Since(start)

	// output results
	fmt.Printf("no chan: %s\nchan: %s\n", elapsed, elapsedWithCh)
}
