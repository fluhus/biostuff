biostuff
========

[![Go Reference](https://pkg.go.dev/badge/github.com/fluhus/biostuff.svg)](https://pkg.go.dev/github.com/fluhus/biostuff)
[![Go Report Card](https://goreportcard.com/badge/github.com/fluhus/biostuff)](https://goreportcard.com/report/github.com/fluhus/biostuff)

Computational biology packages for Go, with emphasis on minimialism.

```
go get github.com/fluhus/biostuff/...
```

Package Overview
----------------

* Data formats
  * [bed](https://pkg.go.dev/github.com/fluhus/biostuff/formats/bed)
  * [fasta](https://pkg.go.dev/github.com/fluhus/biostuff/formats/fasta)
  * [fastq](https://pkg.go.dev/github.com/fluhus/biostuff/formats/fastq)
  * [newick](https://pkg.go.dev/github.com/fluhus/biostuff/formats/newick)
  * [sam](https://pkg.go.dev/github.com/fluhus/biostuff/formats/sam)
* Nucleotide & amino-acid sequence utilities
  * [sequtil](https://pkg.go.dev/github.com/fluhus/biostuff/sequtil)
* Algorithms & data structures
  * [align](https://pkg.go.dev/github.com/fluhus/biostuff/align)
    sequence alignment logic
  * [regions](https://pkg.go.dev/github.com/fluhus/biostuff/regions)
    an index for interval (genes, etc.) overlap lookup
  * [trie](https://pkg.go.dev/github.com/fluhus/biostuff/trie)
    a prefix tree for sequence lookups

Help or Get Help
----------------

Found a bug? Got feedback? Questions? Feel free to
[open an issue](https://github.com/fluhus/biostuff/issues/new)
and let me know!
