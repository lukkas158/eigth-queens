package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Individual []int

type Population [][]int

type Fitness func(Individual) int

func randomSelection(population Population, fitnessFN Fitness) Individual {

	scores := make([]int, len(population))
	probabilities := make([]float64, len(population))

	// Calculates the score and sum of scores for each individual.
	sum := 0
	for i := range population {
		scores[i] = fitnessFN(population[i])
		sum += scores[i]
	}

	// Accumulative probabilities of individuals.
	previousAccumulative := 0.0
	for i := range scores {

		probabilities[i] = float64(scores[i]) / float64(sum)
		probabilities[i] += previousAccumulative
		previousAccumulative = probabilities[i]
	}

	// Generates a random number and compare which individuals has won.
	result := population[0]
	number := rand.Float64() // [0,1)
	for i, probability := range probabilities {
		if number < probability {
			result = population[i]
			break
		}
	}

	return result

}

func reproduce(x, y Individual) Individual {
	var result Individual
	c := rand.Intn(len(x))
	result = append(result, x[0:c]...)
	result = append(result, y[c:]...)

	return result
}

func mutate(x Individual) Individual {
	element := make(Individual, len(x))
	copy(element, x)
	element[rand.Intn(len(x))] = rand.Intn(8)
	return element
}

func geneticAlgorithm(population Population, fitnessFN Fitness) Individual {

	// This for is like the "time". Just to prevent "ridges" and "plateux" loops.
	rand.Seed(time.Now().UnixNano())

	for t := 0; t < 1000; t++ {

		fmt.Println("Epoch", t)
		newPopulation := make(Population, len(population))

		for i := range population {
			x := randomSelection(population, fitnessFN)
			y := randomSelection(population, fitnessFN)
			child := reproduce(x, y)

			// fmt.Printf("%v %v(%v) %v(%v)", population, x, fitnessFN(x), y, fitnessFN(y))

			if rand.Float64() < 0.1 {
				child = mutate(child)

				// fmt.Printf("Mutates ")

			}

			newPopulation[i] = child
			// fmt.Println("Created", child)

		}

		population = newPopulation

		for _, individual := range newPopulation {

			// if fitnessFN(individual) >= 50 {
			// 	fmt.Printf("%v %v ", individual, fitnessFN(individual))

			// }
			if fitnessFN(individual) == 56 {
				return individual
			}
		}

	}

	return nil

}

// I'm using a line equation to find whether the queens attacking each other in the diagonal and horizontal
// If they have the same y-intercept then they are attacking each other.
// I'm using two equation for the two diagonals. One with the slop position and another negative.
func getAttackingsFrom(queenPosition int, queens Individual) int {

	result := 0
	//  y = mx + b
	//  Slop = 1
	//  b1 = y - x
	// 	Slop = -1
	//  b2 = y + x
	b1 := queens[queenPosition] - queenPosition
	b2 := queens[queenPosition] + queenPosition

	//  For horizontal with slop 0
	// b = y
	bh := queens[queenPosition]

	for x, y := range queens {

		if y-x == b1 {
			result++
		} else if y+x == b2 {
			result++
		} else if y == bh {
			result++
		}
	}
	return result

}

func FitnessOneIndividual(individual Individual) int {
	result := 0
	for queenPosition := range individual {
		result += getAttackingsFrom(queenPosition, individual)
	}

	// Complement of all queens attacking each other. Remenber the min value of result is 8
	//  because each queen can attack itself.
	//  The max returned is 56.
	return 64 - result
}

func print_solution(solution Individual) {

	fmt.Println(solution)
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {

			if solution[x] == 7-y {
				fmt.Printf(" X ")

			} else {
				fmt.Printf(" . ")

			}
		}
		fmt.Println("")
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	population := make(Population, 100)

	for i := range population {
		individual := make(Individual, 8)
		for j := range individual {
			individual[j] = rand.Intn(8)
		}
		population[i] = individual
	}

	solution := geneticAlgorithm(population, FitnessOneIndividual)

	fmt.Println("")

	print_solution(solution)

}
