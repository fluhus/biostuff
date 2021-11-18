// Tree traversal functionality.

package newick

// PreOrder calls forEach on every node in the tree, non-recursively. Every node is
// called before its children. The iteration stops when forEach returns false.
func (n *Node) PreOrder(forEach func(*Node) bool) {
	n.traverse(true, forEach)
}

// PostOrder calls forEach on every node in the tree, non-recursively. Every node is
// called after its children. The iteration stops when forEach returns false.
func (n *Node) PostOrder(forEach func(*Node) bool) {
	n.traverse(false, forEach)
}

// Calls forEach on every node in the tree. pre determines whether this is pre- or
// post-order.
func (n *Node) traverse(pre bool, forEach func(*Node) bool) {
	stack := []traversalStep{{n, 0}}
	for len(stack) > 0 {
		stepi := len(stack) - 1
		step := stack[stepi]

		if pre && step.i == 0 {
			if !forEach(step.n) {
				return
			}
		}
		if step.i == len(step.n.Children) {
			if !pre {
				if !forEach(step.n) {
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

// A step in the traversal stack.
type traversalStep struct {
	n *Node
	i int
}
