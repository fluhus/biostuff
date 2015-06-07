package main

import (
	"net"
	"encoding/json"
	"fmt"
	"bioformats/fasta"
)

// Handles a single request from a fasta client.
func handleConnection(c net.Conn) {
	report("Got connection.")
	defer c.Close()
	
	// Read incoming request.
	var req request
	err := json.NewDecoder(c).Decode(&req)
	if err != nil {
		reportf("Error parsing request: %v\n", err)
		writeJsonError(c, "Error parsing request: %v", err)
		return
	}

	report("Request type:", req.Type)
	
	switch req.Type {
	case "sequence":
		chr := req.Sequence.Chr
		start := req.Sequence.Start
		length := req.Sequence.Length
	
		// Find fasta entry.
		var entry *fasta.Entry
		for _, e := range fa {
			if e.Name() == chr {
				entry = e
			}
		}
		
		if entry == nil {
			writeJsonError(c, "No such chromosome: '%s'.", chr)
			return
		}
		
		// Check positions.
		if length < 1 {
			writeJsonError(c, "Invalid length: %d.", length)
			return
		}
		
		if start < 0 {
			writeJsonError(c, "Invalid start position: %d.", start)
			return
		}
		
		if start + length > entry.Length() {
			writeJsonError(c, "Position exceeds chromosome length (max %d).",
					entry.Length())
			return
		}
		
		// Everything is ok!
		reportf("chr=%s start=%d len=%d\n", chr, start, length)
		seq := entry.Subsequence(start, length)
		c.Write([]byte("{\"sequence\":\""))
		c.Write(seq)
		c.Write([]byte("\"}"))
	case "meta":
		m := map[string]int{}
		for i := range fa {
			m[fa[i].Name()] = fa[i].Length()
		}
		json.NewEncoder(c).Encode(m)
	default:
		writeJsonError(c, "Unknown request type: '%s'.", req.Type)
	}
}

// A request from a client.
type request struct {
	Type string  // What is the server asking for?
	             // Type should match struct name.
	Sequence struct {
		Chr string
		Start int
		Length int
	}
}

// Writes a json object with an 'error' field.
func writeJsonError(c net.Conn, s string, a ...interface{}) {
	m := map[string]string{"error" : fmt.Sprintf(s, a...)}
	json.NewEncoder(c).Encode(m)
}
