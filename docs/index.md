# BioStuff

Computational biology packages for Go, with emphasis on minimialism.

```
go get github.com/fluhus/biostuff/...
```

### *Another* computational biology library?

Yes.

This one is optimized for API simplicity. It helps those who need quick,
straightforward solutions where they are not required to learn new things.
It is also optimized for performance, as long as the optimization does not
complicate the API.

### Package overview

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
