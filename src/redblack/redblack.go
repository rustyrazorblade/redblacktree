package redblack

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

func (rb *RedBlackTree) Insert(value int) *Node {
	if rb.root == nil {
		rb.root = NewNode(value)
		rb.root.red = 1
		return rb.root
	}
	inserted := rb.root.Insert(value)
	return inserted
}

func btoi(b bool) int {
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

func (n *Node) Rotate(dir int) {
	// dir is direction, 0 left, 1 right
	opposite_child := n.links[not(dir)]
	affected_grandchild := opposite_child.links[dir]
	n.links[not(dir)] = affected_grandchild
	opposite_child.parent = n.parent
	n.parent = opposite_child
}

func not(i int) int {
	// stupid utility to flip 1/0 since we can't do !int checks
	// i usually hate these 1 liners but it makes the array index more readable
	return i ^ 1
}

func (t *RedBlackTree) IsBalanced() bool {
	return false
}

func is_red(n *Node) bool {
	return n != nil && n.red == 1
}
