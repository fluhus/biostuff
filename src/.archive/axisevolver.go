/* ***************************************************************************
 * Optimizes the axis partition into groups.
 * ***************************************************************************/

package main

import "math/rand"

/*
 * Number of new axes.
 */
const newAxes = 4

/*
 * Holds a partition array and its evaluation.
 * Able to create mutants.
 */
type AxisEvolver struct {
	nodes []Node
	partition [16]int
	value float64
}

/*
 * Returns the value of this partition.
 */
func (ae *AxisEvolver) Value() float64 {
	return ae.value
}

/*
 * Returns a mutated instance based on this one.
 */
func (ae* AxisEvolver) Mutate(strength float64) Evolver {
	result := &AxisEvolver{}
	*result = *ae

	// Go over each axis
	for i := range result.partition {
		if rand.Float64() < strength {
			result.partition[i] = rand.Intn(newAxes)
		}
	}

	// Set new value
	result.SetValue()

	return result
}

/*
 * Sets the value of this instance according to its partition.
 */
func (ae* AxisEvolver) SetValue() {
	// Create help buffer
	var parts [newAxes][]float64
	for i := range parts {
		parts[i] = make([]float64, len(ae.nodes))
	}

	// For each node
	for n := range ae.nodes {
		// Add each old axis to the new one
		for p := range ae.partition {
			parts[ae.partition[p]][n] += float64(ae.nodes[n][p])
		}
	}

	// Set the new value - sum the variances
	ae.value = 0
	for i := range parts {
		ae.value += Variance(parts[i])
	}

	return
}




