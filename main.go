package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	d := NewStdDeck()
	allHands := GenAllHands(&d)

	typeCounts := CountHandTypes(&allHands)

	runTime := time.Since(startTime)

	fmt.Println("New deck: ", d.String())
	fmt.Println("Number of hands: ", len(allHands))
	fmt.Println("Counts: ", typeCounts)
	fmt.Println("Ran in ", runTime)
}
