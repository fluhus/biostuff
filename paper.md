---
title: 'biostuff: a computational biology package suite for Go'
tags:
  - go
  - golang
  - bioinformatics
  - computational biology
authors:
  - name: Amit Lavon
    orcid: 0000-0003-3928-5907
    affiliation: 1
affiliations:
 - name: Weizmann Institute of Science, Israel
   index: 1
date: 24 July 2020
bibliography: paper.bib
---

# Summary

*biostuff* is a computational biology package suite for Go
(a compiled, statically-typed programming language released in 2012).
Go's performance is comparable to C's, which makes it suitable for intensive
data analysis tasks.
*biostuff* is a collection of Go packages which includes parsers, algorithms, and
biological sequence manipulation utilities.
The aim of *biostuff* is to empower researchers using Go who need quick,
straightforward solutions for their computational work,
by providing minimal APIs that require near-zero learning.

# Statement of need

*biostuff* was designed to enable fluent biological research using the
Go programming language.
Resource-intensive software in computational biology is currently written
typically in C, C++ and Java (table 1).
Go is comparable to these languages in terms of performance and
type safety,
but is also garbage-collected, simple and has a rich standard library,
making it an attractive option for
programmers coming from scripting languages such as Python, R and Ruby.

Given the diverse backgrounds of researchers in computational biology,
an accessible implementation of common algorithms and data parsing
routines in Go has the potential to
boost computational biology research.
The current most popular implementation is
*biogo* [@Kortschak:2015].
The aim of *biostuff* is to provide a similar functionality while emphasizing
API minimalism in terms of size and complexity. This can appeal to additional
audiences, specifically researchers who need simplified and straightforward
solutions, and those
without an extensive background in programming or in Go.
It has been shown that error-proneness increases with API complexity
[@Cataldo:2014], so minimizing the API can also contribute to reliability.
Simplifying *biostuff*'s API is achieved by:
1\) preferring use of builtin types as much as possible over introduction of
new types,
2\) allowing configuration of functionality only when absolutely needed, making
the default behavior good for general use, and
3\) exposing few functions, focusing only on what an expert implementer can do
better than a non-expert.

This project strives to make the Go language more accessible to
computational biology researchers and enable them to write fast and scalable
software with ease.

---

**Table 1: Repository counts of top compiled languages on GitHub in topic
"bioinformatics"**

Retrieved from: https://github.com/topics/bioinformatics (accessed 2022-07-24)

|Language|Repository count|
|-|-|
|C++|350|
|Java|242|
|C|203|
|Rust|108|
|Go|87|
|C#|33|
