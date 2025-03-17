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
  * [genbank](https://pkg.go.dev/github.com/fluhus/biostuff/formats/genbank)
  * [newick](https://pkg.go.dev/github.com/fluhus/biostuff/formats/newick)
  * [sam](https://pkg.go.dev/github.com/fluhus/biostuff/formats/sam)
* Algorithms & data structures
  * [align](https://pkg.go.dev/github.com/fluhus/biostuff/align)
    sequence alignment logic
  * [mash](https://pkg.go.dev/github.com/fluhus/biostuff/mash/v2)
    implementation of Mash distance
  * [rarefy](https://pkg.go.dev/github.com/fluhus/biostuff/rarefy)
    rarefaction by read count
  * [regions](https://pkg.go.dev/github.com/fluhus/biostuff/regions)
    an index for interval (genes, etc.) overlap lookup
* Nucleotide & amino-acid sequence utilities
  * [sequtil](https://pkg.go.dev/github.com/fluhus/biostuff/sequtil)

Help or Get Help
----------------

Found a bug? Got feedback? Questions? Feel free to
[open an issue](https://github.com/fluhus/biostuff/issues/new)
and let me know!
