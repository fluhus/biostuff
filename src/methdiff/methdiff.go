// Annotates differentially methylated regions.
package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func main() {
	if len(os.Args) <= 2 {
		os.Exit(1)
	}
	
	fmt.Printf("Tiling '%s'...\n", os.Args[1])
	t1 := make(map[string]tiles)
	err := tileFile(os.Args[1], t1)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}
	
	fmt.Printf("Tiling '%s'...\n", os.Args[2])
	t2 := make(map[string]tiles)
	err = tileFile(os.Args[2], t2)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(2)
	}
	
	// Compare methylation rates.
	fmt.Println("Diffing...")
	for chr := range t1 {
		for pos := range t1[chr] {
			tile1 := t1[chr][pos]
			tile2 := t2[chr][pos]
			
			if tile2 == nil { continue }
			
			r1 := float64(tile1.methd) / float64(tile1.total)
			r2 := float64(tile2.methd) / float64(tile2.total)
			
			fmt.Printf("%s\t%d\t%f\t%f\t%f\n", chr, pos, r1, r2,
					tilediff(tile1, tile2))
		}
	}
}

type tile struct {
	total int
	methd int
}

// Maps from position to tile.
type tiles map[int]*tile

func tileFile(file string, out map[string]tiles) error {
	f, err := os.Open(file)
	if err != nil { return err }
	defer f.Close()
	
	scanner := bufio.NewScanner(f)
	scanner.Scan() // Skip header line.
	
	for scanner.Scan() {
		// Split to fields.
		fields := strings.Split(scanner.Text(), "\t")
		if len(fields) != 12 {
			return fmt.Errorf("Bad number of fields: %d", len(fields))
		}
		
		// Extract numbers from line.
		chr := fields[0]
		pos, err := strconv.Atoi(fields[1])
		if err != nil { return err }
		total, err := strconv.ParseFloat(fields[5], 64)
		if err != nil { return err }
		if total == 0 { continue }  // Avoid parsing 'NA'.
		ratio, err := strconv.ParseFloat(fields[4], 64)
		if err != nil { return err }
		
		methd := int( total * ratio )
		
		// Create chromosome.
		if out[chr] == nil {
			out[chr] = make(map[int]*tile)
		}
		
		// Round position to tile.
		pos = pos / 100 * 100
		
		// Create tile.
		if out[chr][pos] == nil {
			out[chr][pos] = &tile{}
		}
		
		out[chr][pos].total += int(total)
		out[chr][pos].methd += methd
	}
	
	if scanner.Err() != nil {
		return scanner.Err()
	}
	
	return nil
}

func tilediff(tile1, tile2 *tile) float64 {
	return bindiff(tile1.total, tile1.methd, tile2.total, tile2.methd)
}

