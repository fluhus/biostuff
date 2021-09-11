// Package trie provides a prefix tree implementation.
//
// Operations on a trie are linear in query length, regardless of the size of the
// trie. All operations are non-recursive except for marshaling. Marshaling
// recursiveness depends on the standard library implementation.
package trie

import "encoding/json"

// A Trie is a prefix tree that supports lookups.
// A zero value trie is invalid; use NewTrie to create a new instance.
type Trie struct {
	m map[byte]*Trie
}

// NewTrie returns an empty trie.
func NewTrie() *Trie {
	return &Trie{m: map[byte]*Trie{}}
}

// Add inserts b and its prefixes to the trie. If b was already added, the object
// is unchanged.
func (t *Trie) Add(b []byte) {
	cur := t
	for len(b) > 0 {
		next := cur.m[b[0]]
		if next == nil {
			next = NewTrie()
			cur.m[b[0]] = next
		}
		cur = next
		b = b[1:]
	}
}

// Has checks whether b was added to the trie. b could have been added by being a
// prefix of a longer input.
func (t *Trie) Has(b []byte) bool {
	cur := t
	for len(b) > 0 {
		next := cur.m[b[0]]
		if next == nil {
			return false
		}
		cur = next
		b = b[1:]
	}
	return true
}

// Delete removes prefix b from the trie. All sequences that have b as their prefix
// are removed. All other sequences are unchanged. Returns the result of calling
// Has(b) before deleting.
func (t *Trie) Delete(b []byte) bool {
	// Delve in and create a stack.
	stack := make([]*Trie, len(b))
	cur := t
	for i := range b {
		stack[i] = cur
		cur = cur.m[b[i]]
		if cur == nil {
			return false
		}
	}

	// Go back and delede nodes.
	for i := len(stack) - 1; i >= 0; i-- {
		delete(stack[i].m, b[i])
		if len(stack[i].m) > 0 {
			// Stop deleting if node has other children.
			break
		}
	}

	return true
}

// A trie with exported fields for marshaling.
type marshalTrie struct {
	M map[byte]*Trie
}

// MarshalJSON implements the json.Marshaler interface.
func (t *Trie) MarshalJSON() ([]byte, error) {
	m := marshalTrie{t.m}
	return json.Marshal(m)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (t *Trie) UnmarshalJSON(data []byte) error {
	m := marshalTrie{t.m}
	err := json.Unmarshal(data, &m)
	t.m = m.M
	return err
}
