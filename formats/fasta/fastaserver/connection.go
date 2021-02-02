package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fluhus/golgi/formats/fasta"
)

// Handles sequence requests.
func sequenceHandler(w http.ResponseWriter, req *http.Request) {
	report("Got sequence request.")

	chr := req.FormValue("chr")
	startS := req.FormValue("start")
	lengthS := req.FormValue("length")

	if chr == "" {
		fmt.Fprintf(w, "Error: Empty chromosome name.")
		return
	}

	start, err := strconv.Atoi(startS)
	if err != nil {
		fmt.Fprintf(w, "Error: Bad start position: '%s'", startS)
		return
	}

	length, err := strconv.Atoi(lengthS)
	if err != nil {
		fmt.Fprintf(w, "Error: Bad length: '%s'", lengthS)
		return
	}

	// Find fasta.
	var entry *fasta.Fasta
	for _, e := range fa {
		if string(e.Name) == chr {
			entry = e
		}
	}

	if entry == nil {
		fmt.Fprintf(w, "Error: No such chromosome: '%s'.", chr)
		return
	}

	// Check positions.
	if length < 1 {
		fmt.Fprintf(w, "Error: Invalid length: %d.", length)
		return
	}

	if start < 0 {
		fmt.Fprintf(w, "Error: Invalid start position: %d.", start)
		return
	}

	if start+length > len(entry.Sequence) {
		fmt.Fprintf(w, "Error: Position exceeds chromosome length (max %d).",
			len(entry.Sequence))
		return
	}

	// Everything is ok!
	reportf("chr=%s start=%d len=%d\n", chr, start, length)
	seq := entry.Sequence[start : start+length]
	w.Write(seq)
}

// Handles metadata requests.
func metaHandler(w http.ResponseWriter, req *http.Request) {
	reportf("Got meta request.")
	for _, entry := range fa {
		fmt.Fprintf(w, "%s: %d\n", entry.Name, len(entry.Sequence))
	}
}
