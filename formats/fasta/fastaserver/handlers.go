package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fluhus/biostuff/formats/fasta"
)

// TODO(amit): Put sequences in a map.

// Handles sequence requests.
func sequenceHandler(w http.ResponseWriter, req *http.Request) {
	chr := req.FormValue("chr")
	startS := req.FormValue("start")
	lengthS := req.FormValue("length")

	report("Got sequence request", req.Form)

	if chr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: empty chromosome name")
		return
	}

	// Find fasta.
	var entry *fasta.Fasta
	for _, e := range fa {
		if string(e.Name) == chr {
			entry = e
			break
		}
	}

	if entry == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Error: no such chromosome: %q", chr)
		return
	}

	start := 0
	if startS != "" {
		var err error
		start, err = strconv.Atoi(startS)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: bad start position: %q", startS)
			return
		}
	}

	if start >= len(entry.Sequence) || start < 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: start %d exceeds chromosome bounds (max %d)",
			start, len(entry.Sequence))
		return
	}

	length := len(entry.Sequence) - start
	if lengthS != "" {
		var err error
		length, err = strconv.Atoi(lengthS)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error: bad length: %q", lengthS)
			return
		}
	}

	if length < 1 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: invalid length: %d", length)
		return
	}

	if start+length > len(entry.Sequence) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: start %d + length %d exceed chromosome length "+
			"(max %d)", start, length, len(entry.Sequence))
		return
	}

	seq := entry.Sequence[start : start+length]
	w.Write(seq)
}

// Handles metadata requests.
func metaHandler(w http.ResponseWriter, req *http.Request) {
	reportf("Got meta request")
	for _, entry := range fa {
		fmt.Fprintf(w, "%s: %d\n", entry.Name, len(entry.Sequence))
	}
}
