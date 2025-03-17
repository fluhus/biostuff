[![Go Reference](https://pkg.go.dev/badge/github.com/fluhus/biostuff.svg)](https://pkg.go.dev/github.com/fluhus/biostuff)
[![Go Report Card](https://goreportcard.com/badge/github.com/fluhus/biostuff)](https://goreportcard.com/report/github.com/fluhus/biostuff)

Computational biology packages for Go, with emphasis on minimialism.

```
go get github.com/fluhus/biostuff/...
```

## *Another* computational biology library?

Well... Yes.
Each library puts its emphasis on a certain audience and certain use cases.

This one is optimized for API simplicity. It helps those who need quick,
straightforward solutions where they are not required to learn new concepts.
It is also optimized for performance, as long as the optimization does not
complicate the API.

## Package overview

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

## Help or get help

Found a bug? Got feedback? Questions? Feel free to
[open an issue](https://github.com/fluhus/biostuff/issues/new)
and let me know!
