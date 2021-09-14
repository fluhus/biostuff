Golgi
=====

[![Go Reference](https://pkg.go.dev/badge/github.com/fluhus/golgi.svg)](https://pkg.go.dev/github.com/fluhus/golgi)
[![Go Report Card](https://goreportcard.com/badge/github.com/fluhus/golgi)](https://goreportcard.com/report/github.com/fluhus/golgi)

Pure Go libraries for handling biological data. Emphasis on minimalism and
efficiency.

Feedback? Suggestions? Requests? Questions?
[Let me know!](https://github.com/fluhus/golgi/issues/new)

Package Overview
----------------

* Data formats
  * [bed](https://pkg.go.dev/github.com/fluhus/golgi/formats/bed)
  * [fasta](https://pkg.go.dev/github.com/fluhus/golgi/formats/fasta)
  * [fastq](https://pkg.go.dev/github.com/fluhus/golgi/formats/fastq)
  * [sam](https://pkg.go.dev/github.com/fluhus/golgi/formats/sam)
* Nucleotide & amino-acid sequence utilities
  * [sequtil](https://pkg.go.dev/github.com/fluhus/golgi/sequtil)
* Algorithms & data structures
  * [trie](https://pkg.go.dev/github.com/fluhus/golgi/trie)
    a prefix tree for sequence lookups
  * [regions](https://pkg.go.dev/github.com/fluhus/golgi/regions)
    an index for interval (genes, etc.) overlap lookup
