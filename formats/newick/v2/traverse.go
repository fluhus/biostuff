// Tree traversal functionality.

package newick

import "iter"

// PreOrder returns an iterator over nodes in the tree.
// Every node is called before its children.
//
// Does not make recursive calls.
func (n *Node) PreOrder() iter.Seq[*Node] {
	return n.traverse(true)
}

// PostOrder returns an iterator over nodes in the tree.
// Every node is called after its children.
//
// Does not make recursive calls.
func (n *Node) PostOrder() iter.Seq[*Node] {
	return n.traverse(false)
}

// Returns an iterator over nodes in the tree.
// pre determines whether this is pre- or post-order.
func (n *Node) traverse(pre bool) iter.Seq[*Node] {
	return func(yield func(*Node) bool) {
		stack := []traversalStep{{n, 0}}
		for len(stack) > 0 {
			stepi := len(stack) - 1
			step := stack[stepi]

			if pre && step.i == 0 {
				if !yield(step.n) {
					return
				}
			}
			if step.i == len(step.n.Children) {
				if !pre {
					if !yield(step.n) {
						return
					}
				}
				stack = stack[:len(stack)-1]
				continue
			}

			stack = append(stack, traversalStep{step.n.Children[step.i], 0})
			stack[stepi].i++
		}
	}
}

// A step in the traversal stack.
type traversalStep struct {
	n *Node
	i int
}
