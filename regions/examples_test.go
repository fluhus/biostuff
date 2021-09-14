package regions

import "fmt"

func ExampleIndex() {
	gene0start, gene0end := 100, 200
	gene1start, gene1end := 150, 250
	gene2start, gene2end := 130, 300

	starts := []int{gene0start, gene1start, gene2start}
	ends := []int{gene0end, gene1end, gene2end}
	idx := NewIndex(starts, ends)

	fmt.Println(idx.At(140)) // Genes that overlap with position 140
	fmt.Println(idx.At(200)) // Genes that overlap with position 200
	//Output:
	//[0 2]
	//[1 2]
}
