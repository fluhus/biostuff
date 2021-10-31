// Package trie provides a prefix tree implementation.
//
// Add, Has and Delete operations are linear in query length, regardless of the
// size of the trie. All operations are non-recursive except for marshaling.
package trie

import "encoding/json"

// A Trie is a prefix tree that supports lookups.
// A zero value trie is invalid; use New to create a new instance.
type Trie struct {
	m map[byte]*Trie
}

// New returns an empty trie.
func New() *Trie {
	return &Trie{m: map[byte]*Trie{}}
}

// Add inserts b and its prefixes to the trie. If b was already added, the object
// is unchanged.
func (t *Trie) Add(b []byte) {
	cur := t
	for len(b) > 0 {
		next := cur.m[b[0]]
		if next == nil {
			next = New()
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
	M map[byte]*Trie `json:"m"`
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

// ForEach calls f for each final sequence (leaf) in the trie. A final sequence
// is a sequence that is not a prefix of a longer sequence. f should return
// whether or not the iteration should continue. f's input slice may be
// overwritten by subsequent iterations.
func (t *Trie) ForEach(f func([]byte) bool) {
	stack := []*forEachStep{{t, t.keys(), 0}}
	var cur []byte
	for {
		step := stack[len(stack)-1]
		if len(step.t.m) == 0 { // Reached a leaf.
			if len(cur) > 0 && !f(cur) {
				break
			}
		}
		if step.i == len(step.t.m) { // Finished with this branch.
			stack = stack[:len(stack)-1]
			if len(stack) == 0 { // Done.
				break
			}
			cur = cur[:len(cur)-1]
			continue
		}
		// Handle next child.
		key := step.k[step.i]
		child := step.t.m[key]
		stack = append(stack, &forEachStep{child, child.keys(), 0})
		step.i++
		cur = append(cur, key)
	}
}

// Returns the keys of a trie.
func (t *Trie) keys() []byte {
	result := make([]byte, 0, len(t.m))
	for k := range t.m {
		result = append(result, k)
	}
	return result
}

// A step in the for-each stack.
type forEachStep struct {
	t *Trie  // Current trie
	k []byte // Trie's keys
	i int    // Current key
}
