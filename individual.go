package main

import (
	"math/rand"
)

//
// Individual represents a set of boolean flags indicating which traits exist and which do not.
//
type Individual uint32

//
// Breed creates a new child by mating the Individual that is being acted on with the provided
// Partner.
//
func (individual *Individual) Breed(partner Individual) (child Individual) {
	var parentAMasked Individual = (*individual)
	var parentBMasked Individual = partner

	return (parentAMasked | parentBMasked)
}

//
// Mutate randomly turns on and off one of the Individual's traits.
//
func (individual *Individual) Mutate() {
	var genePos int = rand.Intn(32)
	var geneMask Individual = 0x00000001 << genePos

	var geneCur = (*individual) & geneMask
	var geneNew = ^geneCur

	(*individual) &^= geneMask
	(*individual) |= (geneNew & geneMask)
}

//
// Score determines how close the Individual's traits are to the provided God.
//
func (individual *Individual) Score(god Individual) int {
	const maskFinal Individual = 0x80000000

	var mask Individual = 0x00000001
	var score int = 0

	for {
		individualMasked := ((*individual) & mask)
		godMasked := (god & mask)

		if individualMasked == godMasked {
			score++
		}

		if mask == maskFinal {
			break
		}

		mask <<= 1
	}

	return score
}
