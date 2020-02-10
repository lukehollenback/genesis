package creature

import (
	"testing"
)

func TestMutate(t *testing.T) {
	god := CreateGodCreature()
	creature := CreateCreatureWithGenotype(god.Genotype, god)

	creature.Mutate(100)

	if (god.CurScore - creature.CurScore) != 1 {
		t.Errorf("Mutation did not modify exactly one gene in the creature's genome. The God " +
			"Creature's score was %d, but the mutated Creature's score was %d.", god.CurScore,
			creature.CurScore)
	}
}