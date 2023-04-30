package main

import (
	"fmt"
	"testing"
)

func TestTimedSum(t *testing.T) {
	sum1 := timedSumFunc(Sum1)
	sum2 := timedSumFunc(Sum2)

	fmt.Printf("%d, %d\n", sum1(-10000, 1000000), sum2(-10000, 100000))
}
