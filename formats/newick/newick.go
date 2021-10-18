// Package newick handles reading and writing Newick-formatted trees.
//
// This package uses the format described in:
// https://en.wikipedia.org/wiki/Newick_format
package newick

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
)

// A Node is a single node in a tree, along with the subtree that is under it.
type Node struct {
	Name     string  // Node name. Can be empty.
	Distance float64 // Distance from parent. 0 is treated as none.
	Children []*Node // Child nodes. Can be nil.
}

// Newick returns a condensed Newick-format representation of this node.
func (n *Node) Newick() []byte {
	buf := bytes.NewBuffer(nil)
	n.newick(buf)
	buf.WriteByte(';')
	return buf.Bytes()
}

// Writes a single node/subtree to the buffer.
func (n *Node) newick(buf *bytes.Buffer) {
	if len(n.Children) > 0 {
		buf.WriteByte('(')
		for i, c := range n.Children {
			if i > 0 {
				buf.WriteByte(',')
			}
			c.newick(buf)
		}
		buf.WriteByte(')')
	}
	buf.WriteString(n.Name)
	if n.Distance != 0 {
		fmt.Fprint(buf, ":", n.Distance)
	}
}

// A Reader reads Newick-formatted trees.
type Reader struct {
	r *bufio.Reader
	b *bytes.Buffer
}

// NewReader returns a Reader that reads from the given stream.
func NewReader(r io.Reader) *Reader {
	return &Reader{bufio.NewReader(r), &bytes.Buffer{}}
}

// Read reads a single tree from the input. Can be called multiple times to read
// subsequent trees.
func (r *Reader) Read() (*Node, error) {
	// TODO(amit): Break this down to subfunctions.

	// Possible states.
	const (
		beforeNode = iota
		afterName
		afterColon
		afterDist
		afterChildren
	)

	stack := []*Node{{}}
	state := beforeNode
	readAny := false // Accepting EOF only if no tokens are available.

loop:
	for {
		token, err := r.nextToken()
		if err != nil {
			if err == io.EOF && readAny {
				return nil, io.ErrUnexpectedEOF
			}
			return nil, err
		}
		readAny = true
		switch token {
		case "(":
			if state != beforeNode {
				return nil, fmt.Errorf("unexpected '('")
			}
			cur := stack[len(stack)-1]
			node := &Node{}
			cur.Children = append(cur.Children, node)
			stack = append(stack, node)
		case ")":
			if state == afterColon || state == afterChildren {
				return nil, fmt.Errorf("unexpected ')'")
			}
			if len(stack) == 1 {
				return nil, fmt.Errorf("too many ')'")
			}
			stack = stack[:len(stack)-1]
			state = afterChildren
		case ",":
			if state == afterColon {
				return nil, fmt.Errorf("unexpected ',' after ':'")
			}
			if len(stack) == 1 {
				return nil, fmt.Errorf("found ',' after top level node")
			}
			node := &Node{}
			parent := stack[len(stack)-2]
			parent.Children = append(parent.Children, node)
			stack[len(stack)-1] = node
			state = beforeNode
		case ":":
			if state == afterColon || state == afterDist {
				return nil, fmt.Errorf("unexpected ':' after ':'")
			}
			state = afterColon
		case ";":
			if len(stack) != 1 {
				return nil, fmt.Errorf("unexpected ';' at depth %d", len(stack)+1)
			}
			if state == afterColon {
				return nil, fmt.Errorf("unexpected ';' after ':'")
			}
			break loop
		default:
			if state == afterName || state == afterDist {
				return nil, fmt.Errorf("unexpected token: %q", token)
			}
			cur := stack[len(stack)-1]
			if state == beforeNode || state == afterChildren {
				cur.Name = token
				state = afterName
				break
			}
			if state != afterColon {
				panic(fmt.Sprintf("unexpected state: %d", state))
			}
			dist, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return nil, err
			}
			cur.Distance = dist
			state = afterDist
		}
	}
	return stack[0], nil
}

// Reads a single token from the input stream.
func (r *Reader) nextToken() (string, error) {
	r.b.Reset()
loop:
	for {
		b, err := r.r.ReadByte()
		if err != nil {
			if err == io.EOF && r.b.Len() > 0 {
				break loop
			}
			return "", err
		}
		switch b {
		case '(', ')', ',', ':', ';':
			if r.b.Len() > 0 {
				r.r.UnreadByte()
				break loop
			}
			return string([]byte{b}), nil
		case ' ', '\t', '\n', '\r':
			if r.b.Len() > 0 {
				break loop
			}
		default:
			r.b.WriteByte(b)
		}
	}
	return r.b.String(), nil
}