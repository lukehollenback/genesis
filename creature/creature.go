package creature

import (
	"fmt"
	"math/rand"
)

//
// Creature structures represent individuals with specific set of genes (and thus traits).
//
type Creature struct {
	//
	// Genotype is a set of bitwise switches that represent which traites the Creature has and does
	// note have.
	//
	Genotype uint32

	//
	// A reference to the Creature's god (a.k.a. the ideal version of itself).
	//
	God *Creature

	//
	// Score is the last calculated rating of the Creature's traits when comparted to its "God
	// Creature".
	//
	CurScore int
}

//
// CreateCreature creates and scores a new Creature structure with a random genotype.
//
func CreateCreature(god *Creature) *Creature {
	var creature Creature = Creature {
		Genotype: rand.Uint32(),
		God: god,
	}

	creature.Score()

	return &creature
}

//
// CreateCreatureWithGenotype creates and scores a new Creature structure with a specified genotype.
//
func CreateCreatureWithGenotype(genotype uint32, god *Creature) *Creature {
	var creature Creature = Creature {
		Genotype: genotype,
		God: god,
	}

	creature.Score()

	return &creature
}

//
// CreateGodCreature creates and scores a new "God Creature" structure with a random genotype.
//
func CreateGodCreature() *Creature {
	var creature Creature = Creature {
		Genotype: rand.Uint32(),
	}

	creature.Score()

	return &creature
}

func (creature *Creature) String() string {
	return fmt.Sprintf("%032b (score = %d)", creature.Genotype, creature.CurScore)
}

//
// Breed creates and scores a new child by mating the Creature with the provided Partner.
//
func (creature *Creature) Breed(partner *Creature) *Creature {
	// TODO: Check that both god's are the same.

	//
	// Actually combine the two Creatures.
	//
	// NOTE: We apply a random filter to each parent so that only certain genes from each are passed
	//  on to the generated kids.
	//
	var creatureGenotypeMasked uint32 = (creature.Genotype & rand.Uint32())
	var partnerGenotypeMasked uint32 = (partner.Genotype & rand.Uint32())
	var combinedGenotype uint32 = (creatureGenotypeMasked | partnerGenotypeMasked)

	return CreateCreatureWithGenotype(combinedGenotype, creature.God)
}

//
// Score determines how close the Creature's traits to those of its god (a.k.a. to those of the
// ideal version of itself). The caluclated score will be be updated on the actual Creature as well
// as returned.
//
func (creature *Creature) Score() int {
	//
	// "God" Creatures are always "perfect" and do not have a god themselves.
	//
	if creature.God == nil {
		creature.CurScore = 32
		return 32
	}

	//
	// Scroll through the Creature's genotype and compare it to that of its god.
	//
	const maskFinal uint32 = 0x80000000

	var mask uint32 = 0x00000001
	var score int = 0

	for {
		creatureMasked := (creature.Genotype & mask)
		godMasked := (creature.God.Genotype & mask)

		if creatureMasked == godMasked {
			score++
		}

		if mask == maskFinal {
			break
		}

		mask <<= 1
	}

	creature.CurScore = score
	return score
}

//
// Mutate randomly turns on and off one of the Individual's traits if a random number <= to the
// provided probability (out of 100) is chosen. After mutation, the Creature's score is
// recalculated.
//
func (creature *Creature) Mutate(probability int) {
	if rand.Intn(100) > probability {
		return
	}

	var genePos int = rand.Intn(32)
	var geneMask uint32 = 0x00000001 << genePos

	var geneCur = creature.Genotype & geneMask
	var geneNew = ^geneCur

	creature.Genotype &^= geneMask
	creature.Genotype |= (geneNew & geneMask)

	creature.Score()
}