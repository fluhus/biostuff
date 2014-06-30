/* ***************************************************************************
 * A node that converts DNA sequences to metric coordinates.
 * ***************************************************************************/

package main

// *** NODE *******************************************************************

/*
 * Node type - 1 coordinate for each pair of nucleotides.
 */
type Node [16]uint8

/*
 * Locates this node according to the given sequence.
 */
func (n *Node) Locate(sequence []byte) {
	// Reset to zeros
	for i := range n {
		n[i] = 0
	}

	// Locate
	for i := range sequence {
		// Ignore first char (no pair)
		if i == 0 {continue}

		// Find coordinate
		a := Ntoi(sequence[i])
		b := Ntoi(sequence[i-1])
		//c := Ntoi(sequence[i-2])

		// Check for illegal characters
		if a == -1 || b == -1 {continue}

		coord := a + b*4 //+ c*16
		n[coord]++
	}
}

/*
 * Returns L1 distance between 2 nodes.
 */
func (n *Node) Distance(other *Node) float64 {
	d := float64(0)

	for i := range n {
		if n[i] > other[i] {
			d += float64(n[i] - other[i])
		} else {
			d += float64(other[i] - n[i])
		}
	}

	return d
}


// *** MINI-NODE **************************************************************

/*
 * The number of coordinates in a mini-node.
 */
const miniNodeSize = 4

/*
 * A mini-node - holds a compact optimized representation of a node.
 */
type MiniNode [miniNodeSize]int

/*
 * Locates this node according to the grouping constant.
 */
func (n *MiniNode) Locate(sequence []byte) {
	// Reset to zeros
	for i := range n {
		n[i] = 0
	}

	// Create a node
	var node Node
	node.Locate(sequence)

	// Grouping of the regular's node in the mini-node.
	// [could not declare a constant array]
	miniNodeGrouping := []int{3, 2, 3, 1, 2, 2, 2, 2, 3, 2, 0, 0, 1, 2, 0, 1}

	// Group coordinates
	for i := range node {
		n[miniNodeGrouping[i]] += int(node[i])
	}
}





