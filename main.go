package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const maxChildrenPerGeneration int = 1000
const parentAMask Individual = 0xF0F0F0F0
const parentBMask Individual = 0x0F0F0F0F

func main() {
	//
	// Seed the random number generator so that we don't always get the same "random" numbers.
	//
	rand.Seed(time.Now().UTC().UnixNano())

	//
	// Randomly generate the God and the first generation of parents.
	//
	var parentA Individual = Individual(rand.Uint32())
	var parentB Individual = Individual(rand.Uint32())
	var god Individual = Individual(rand.Uint32())
	var godScore int = god.Score(god)

	printSeparator()
	fmt.Printf("Adam  = %08X (score = %d)\n", parentA, parentA.Score(god))
	fmt.Printf("Eve   = %08X (score = %d)\n", parentB, parentB.Score(god))
	fmt.Printf("God   = %08X (score = %d)\n", god, godScore)

	//
	// Breed generations until one evolves to God.
	//
	for {
		parentA, parentB = breedGeneration(parentA, parentB, god, godScore)

		time.Sleep(1 * time.Second)
	}
}

func breedGeneration(parentA Individual, parentB Individual, god Individual, godScore int) (newParentA Individual, newParentB Individual) {
	var childCount int = (2 + rand.Intn(maxChildrenPerGeneration-2))
	var childChannel chan Individual = make(chan Individual, childCount)

	printSeparator()
	fmt.Printf("Breeding next generation (of %d children)...\n", childCount)

	for i := 0; i < childCount; i++ {
		go breedChild(childChannel, parentA, parentB)
	}

	var bestChild Individual
	var bestChildScore int = 0
	var secondBestChild Individual
	var secondBestChildScore int = 0

	for i := 0; i < childCount; i++ {
		var child Individual = <-childChannel
		var childScore int = child.Score(god)

		if childScore > bestChildScore {
			if bestChildScore > secondBestChildScore {
				secondBestChild = bestChild
				secondBestChildScore = bestChildScore

				fmt.Printf("Auto New Second Best Child = %08X (score = %d)\n", secondBestChild, secondBestChildScore)
			}

			bestChild = child
			bestChildScore = childScore

			fmt.Printf("New Best Child            = %08X (score = %d)\n", bestChild, bestChildScore)
		} else if childScore > secondBestChildScore {
			secondBestChild = child
			secondBestChildScore = childScore

			fmt.Printf("New Second Best Child     = %08X (score = %d)\n", secondBestChild, secondBestChildScore)
		}
	}

	if bestChildScore == godScore {
		fmt.Printf("...Best Child has Evolved to God\n")
		os.Exit(0)
	} else if secondBestChildScore == godScore {
		fmt.Printf("...Second Best Child has Evolved to God\n")
		os.Exit(0)
	}

	fmt.Printf("...Finished (scores = %d, %d)\n", bestChildScore, secondBestChildScore)

	return bestChild, secondBestChild
}

func breedChild(c chan<- Individual, parentA Individual, parentB Individual) {
	var child Individual = parentA.Breed(parentB)

	child.Mutate()

	c <- child
}

func printSeparator() {
	fmt.Printf("--------------------------------------------------\n")
}
