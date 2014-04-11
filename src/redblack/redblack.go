package redblack
import (
	"fmt"
)

/*
1. A node is either red or black.
2. The root is black. (This rule is sometimes omitted. Since the root can always be changed from red to black, but not necessarily vice-versa, this rule has little effect on analysis.)
3. All leaves (NIL) are black. (All leaves are same color as the root.)
4. Every red node must have two black child nodes.
5. Every path from a given node to any of its descendant leaves contains the same number of black nodes.
 */

type Node struct {
	links [2]*Node // left = 0
	parent *Node
	value int
	red int8
}

func NewNode(value int) *Node {
	tmp := Node{value:value}
	tmp.red = 1
	return &tmp
}

type RedBlackTree struct {
	root *Node
	balance bool
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{balance:true}

}

func NewTree() *RedBlackTree {
	// returns a new tree, but it won't auto balance.  this will make testing certain
	// cases easier
	return &RedBlackTree{balance:false}
}

func (rb *RedBlackTree) RotateLeft(n *Node) {
	n.Rotate(0)
	if n == rb.root {
		// no longer the root, stupid, the parent is
		rb.root = n.parent
	}
}
func (rb *RedBlackTree) RotateRight(n *Node) {
	n.Rotate(1)
	if n == rb.root {
		// no longer the root, stupid, the parent is
		rb.root = n.parent
	}
}
func (t *RedBlackTree) Rotate(n *Node, dir int8) {
	if dir == 0 {
		t.RotateLeft(n)
		return
	}
	t.RotateRight(n)
}
func (t *RedBlackTree) fixUp(n *Node) {
	if n.parent == nil {
		// root case
		return
	}
	if !is_red(n.parent) {
		// black parent, we're cool
		return
	}
	// get grand parent

	parent := n.parent
	gp := parent.parent

	uncle_op_side := btoi(gp.links[0] == parent)

	uncle := gp.links[uncle_op_side]

	if is_red(uncle) { // we know the parent is red already
		// case 3 recolor
		// repaint parent & uncle black, gp red
		// then fixUp(gp)
		recolor(parent,  0)
		recolor(uncle, 0)
		recolor(parent, 1)

		t.fixUp(gp)
		return
	}
	// we already know the parent is red, the uncle black
	// case 4
	// detect a zig zag
	this_node_side := btoi(parent.links[1] == n)
	parent_node_side := not(uncle_op_side)
	if this_node_side != parent_node_side {
		// we need to know the direction the parent & GP are swinging
		// if we hit, we rotate in the direction of the gp->p
		// left rotation is 0
		fmt.Println("zig zag detected", parent_node_side, this_node_side)
		t.Rotate(n.parent, not(parent_node_side))
		return
	}

	// case 5
	// detect a LEFT LEFT or RIGHT RIGHT

}

func recolor(n *Node, red int8) {
	if n != nil {
		n.red = red
	}
}
func (rb *RedBlackTree) Insert(value int) *Node {
	if rb.root == nil {
		rb.root = NewNode(value)
		rb.root.red = 0
		return rb.root
	}
	inserted := rb.root.Insert(value)
	if rb.balance {
		rb.fixUp(inserted)
	}
	return inserted
}

func btoi(b bool) int8 {
	if b {
		return 1
	}
	return 0
}


func (r *RedBlackTree) Get(value int) (*Node, bool) {
	// returns the node matching the value,true or nil, false
	tmp := r.root.Get(value)
	if tmp != nil {
		return tmp, true
	}
	return nil, false

}

func (n *Node) Get(value int) *Node {
	if n.value == value {
		return n
	} else {
		next := n.links[btoi(value > n.value)]
		if next == nil {
			return nil
		} else {
			return next.Get(value)
		}
	}
}

func (n *Node) Insert(value int) *Node {
	i := value > n.value
	next := n.links[btoi(i)]
	if next != nil {
		return next.Insert(value)
	} else {
		tmp := NewNode(value)
		tmp.parent = n
		n.links[btoi(i)] = tmp
		return tmp

	}

}

func (n *Node) Rotate(dir int8) {
	// dir is direction, 0 left, 1 right
	opposite_child := n.links[not(dir)]
	affected_grandchild := opposite_child.links[dir]
	n.links[not(dir)] = affected_grandchild
	opposite_child.parent = n.parent
	n.parent = opposite_child
}

func not(i int8) int8 {
	// stupid utility to flip 1/0 since we can't do !int checks
	// i usually hate these 1 liners but it makes the array index more readable
	return i ^ 1
}

func (t *RedBlackTree) IsBalanced() bool {
	if is_red(t.root) {
		return false
	}
	return t.root.IsBalanced()

}

func (n *Node) IsBalanced() bool {
	if n.red == 0 {
		// make sure the children are both black
		if is_red(n.links[0]) || is_red(n.links[1]) {
			return false
		}
	}
	return true
}

func is_red(n *Node) bool {
	return n != nil && n.red == 1
}
