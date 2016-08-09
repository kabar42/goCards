package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()

	d := NewStdDeck()
	allHands := GenAllHands(&d)

	generationDuration := time.Since(startTime)
	generationTime := time.Now()

	typeCounts := CountHandTypes(&allHands)

	runTime := time.Since(generationTime)

	fmt.Println("New deck: ", d.String())
	fmt.Println("Number of hands: ", len(allHands))
	fmt.Println("Counts: ", typeCounts)
	fmt.Println("Generation time: ", generationDuration)
	fmt.Println("Classification time: ", runTime)
	fmt.Println("Runtime: ", generationDuration+runTime)
}
