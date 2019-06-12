/* ***************************************************************************
 * An evolver learning machine.
 * ***************************************************************************/

package main

import "fmt"
import "runtime"

/*
 * An interface of a mutable instance, that can take part in the evolver
 * machine.
 * Mutate returns the mutated instance. Strength represents the amount of
 * mutation, such that 1 is a completely randomized instance, and 0 is not
 * mutated at all. Only values within [0,1] are allowed.
 * Value returns how good this instance is, so that the best instances can be
 * selected. A higher value means a better (more optimal) instance.
 */
type Evolver interface {
	Mutate(strength float64) Evolver
	Value() float64
}

/*
 * Sorts the evolvers according to their values, in DESCENDING order.
 */
func SortEvolvers(e []Evolver) {
	// From last element to first
	for i := len(e) - 1; i >= 0; i-- {
		// Bubble until i'th element
		for j := 0; j < i; j++ {
			if e[j].Value() < e[j+1].Value() {
				t := e[j]
				e[j] = e[j+1]
				e[j+1] = t
			}
		}
	}

	return
}

/*
 * The actual learning process.
 * Creates 'selected'**2 instances by invoking Mutate(1). Then for the desired
 * number of iterations, mutation coefficient is multiplied by 'mutratio' each
 * iteration and the best 'selected' instances are preserved.
 * Returns the best instance.
 */
func Evolve(source Evolver, mutratio float64, selected int, iterations int, verbose bool) Evolver {
	// Create an array of evolvers
	evolvers := make([]Evolver, selected * selected)

	// Each element is a mutant
	for i := range evolvers {
		evolvers[i] = source.Mutate(1)
	}
	SortEvolvers(evolvers)

	// Mutation strength counter
	mutstrength := mutratio

	// Helper buffer
	ebuffer := make([]Evolver, selected)

	// Perform iterations
	for i := 0; i < iterations; i++ {
		if verbose {
			fmt.Printf("Evolve:\titeration=%d\tbest=%.3f\tmutratio=%.3f\n", i, evolvers[0].Value(), mutstrength)
		}

		// Copy selected into buffer
		copy(ebuffer, evolvers)

		// Create threads
		numProcs := runtime.NumCPU()
		c := make(chan int, numProcs)
		done := make(chan int, numProcs)
		for i := 0; i < numProcs; i++ {
			go func() {
				for e := range c {
					evolvers[e] = ebuffer[e % selected].Mutate(mutstrength)
				}
				done <- 1
			}()
		}

		// Put indexes in the channel
		for i := range evolvers {
			c <- i
		}
		close(c)

		// Wait for threads to terminate
		for i := 0; i < numProcs; i++ {
			<-done
		}

		// Sort
		SortEvolvers(evolvers)

		// Update mutation strengh
		mutstrength *= mutratio
	}

	// Return top evolver
	return evolvers[0]
}






