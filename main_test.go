package main

import (
	"math/rand"
	"testing"
)

func TestAttackThreeElementsSameRow(t *testing.T) {
	value := getAttackingsFrom(2, []int{2, 2, 2})

	if value != 3 {
		t.Fatalf("Must return 3 points, returned %v", value)
	}
}

func TestAttackFourElementsTwoInDiagonal(t *testing.T) {
	value := getAttackingsFrom(0, []int{0, 1, 2, 1})

	if value != 3 {
		t.Fatalf("Must return 3 points, returned %v", value)
	}
}

func TestFitnessOneIndividualAllSameRow(t *testing.T) {
	value := FitnessOneIndividual(Individual{1, 1, 1, 1, 1, 1, 1, 1})

	if value != 0 {
		t.Fatalf("Must return 0 points, returned %v", value)
	}
}

func TestFitnessOneIndividualRowsDifferents(t *testing.T) {
	value := FitnessOneIndividual(Individual{1, 1, 1, 1, 1, 1, 1, 0})

	if value != 12 {
		t.Fatalf("Must return 12 points, returned %v", value)
	}
}

func TestFitnessSolution(t *testing.T) {

	value := FitnessOneIndividual(Individual{1, 3, 5, 7, 2, 0, 6, 4})

	if value != 56 {
		t.Fatalf("Must return 56 points, returned %v", value)
	}
}

func TestMutate(t *testing.T) {

	rand.Seed(1)
	i := Individual{1, 1, 1, 1, 1, 1}
	result := mutate(i)

	if result[5] != 7 {
		t.Fatalf("Last element must be 7. Got %v", result)
	}

}

func TestReproducing(t *testing.T) {

	result := reproduce(Individual{1, 1, 2, 2}, Individual{2, 2, 3, 3})

	if (result[2] != 2) || result[3] != 3 {
		t.Fatalf("Expected [1 1 2 3] but got [%v]", result)
	}

}

func TestRandomSelection(t *testing.T) {

	rand.Seed(10)
	result := randomSelection(
		Population{
			Individual{1, 2, 2, 3, 4, 5, 7},
			Individual{7, 3, 1, 1, 1, 1, 6},
			Individual{1, 1, 3, 1, 1, 1, 2}}, FitnessOneIndividual)

	if (result[5] != 1) || result[6] != 6 {
		t.Fatalf("Expected [7 3 1 1 1 1 6] but got [%v]", result)
	}

}
