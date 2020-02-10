package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/lukehollenback/genesis/creature"
)

const maxChildrenPerGeneration = 1000
const mutationProbability = 75

func main() {
	//
	// Seed the random number generator so that we don't always get the same "random" numbers.
	//
	rand.Seed(time.Now().UTC().UnixNano())

	//
	// Randomly generate the God (a.k.a. the Creature with all ideal traits) and the first generation
	// of parents.
	//
	var god, parentA, parentB *creature.Creature

	god = creature.CreateGodCreature()
	parentA = creature.CreateCreature(god)
	parentB = creature.CreateCreature(god)

	printSeparator()
	fmt.Printf("God  = %s\n", god)
	fmt.Printf("Adam = %s\n", parentA)
	fmt.Printf("Eve  = %s\n", parentB)

	//
	// Breed generations until one evolves to have the same traits as its god.
	//
	genCount := 0

	for {
		//
		// Determine the generation number and figure out how many children we are going to breed into
		// it.
		//
		genCount++
		childCount := (2 + rand.Intn(maxChildrenPerGeneration - 2))

		printSeparator()
		fmt.Printf("Breeding next generation (%d) (of %d children)...\n", genCount, childCount)

		//
		// Actually breed the generation's children.
		//
		parentA, parentB = breedGeneration(parentA, parentB, childCount)

		fmt.Printf("  ↝        Best Child = %s\n", parentA)
		fmt.Printf("  ↝ Second Best Child = %s\n", parentB)
		fmt.Printf("...Finished.\n")

		//
		// Assess if the best child has acquired the same traits as its god. Because of the sorting that
		// is performed inside of te actual breedGeneration(...) function, we do NOT need to check the
		// second best child – they will either be just as good or worse than the best child.
		//
		if parentA.CurScore == parentA.God.CurScore {
			printSeparator()
			fmt.Printf("Found best possible child after %d generations.\n", genCount)
			fmt.Printf("Best Possible Child = %s\n", parentA)

			break
		}

		//
		// Yield for a little while so that we don't hammer the CPU.
		//
		time.Sleep(1 * time.Second)
	}
}

//
// breedGeneration leverages goroutines to breed all of the necessary children from the provided
// parents.
//
func breedGeneration(parentA *creature.Creature, parentB *creature.Creature, childCount int) (bestChild *creature.Creature, secondBestChild *creature.Creature) {
	//
	// Spin off goroutines for each child that needs to be bred.
	//
	childChannel := make(chan *creature.Creature, childCount)

	for i := 0; i < childCount; i++ {
		go breedChild(childChannel, parentA, parentB)
	}

	//
	// Wait for all of the goroutines to join back and recieve all of the children that they generate.
	//
	children := make([]*creature.Creature, childCount)

	for i := 0; i < childCount; i++ {
		children[i] = <-childChannel
	}

	//
	// Sort the recieved children and return the best ones.
	//
	sort.SliceStable(children, func(i, j int) bool {
		return (children[i].CurScore > children[j].CurScore)
	})

	return children[0], children[1]
}

//
// breedChild is the implementation of a goroutine that generates, mutates, and performs natural
// selection on a new child birthed from the provided parental Creatures.
//
func breedChild(c chan<- *creature.Creature, parentA *creature.Creature, parentB *creature.Creature) {
	//
	// Breed a child from the two parents.
	//
	child := parentA.Breed(parentB)

	//
	// Subject to random chance and probability, mutate a random gene of the child's genotype.
	//
	child.Mutate(mutationProbability)

	//
	// Return the child and close the goroutine.
	//
	c <- child
}

//
// printSeparator simply outputs a line to separate sections of output.
//
func printSeparator() {
	fmt.Printf("--------------------------------------------------------------------------------\n")
}
