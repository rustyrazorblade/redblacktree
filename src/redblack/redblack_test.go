package redblack

import (
	"testing"
	. "gopkg.in/check.v1"
	"fmt"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestRootInsert(c *C) {
	rb := NewTree()
	rb.Insert(1)
	c.Check(rb.root.value, Equals, 1)
}

func (s *MySuite) TestChildInsert(c *C) {
	rb := NewTree()
	rb.Insert(2)
	c.Check(rb.root.value, Equals, 2)

	rb.Insert(1)
	c.Check(rb.root.links[0].value, Equals, 1)

	rb.Insert(3)
	c.Check(rb.root.links[1].value, Equals, 3)

	rb.Insert(5)
	rb.Insert(6)
}

func PopulateTree(rb *RedBlackTree) *RedBlackTree {
	/*
			       10
				/	   \
			   5	    11
		   /      \
		  3		   6
		 / \
		1  4

	 */
	keys := []int{10, 5, 6, 3, 1, 4, 11}
	for _, value := range keys {
		fmt.Println(value)
		rb.Insert(value)
	}
//	rb.Insert(10)
//	rb.Insert(5)
//	rb.Insert(6)
//	rb.Insert(3)
//	rb.Insert(1)
//	rb.Insert(4)
//	rb.Insert(11)
	return rb
}


func (s *MySuite) TestGet(c *C) {
	rb := NewTree()
	ten  := rb.Insert(10)
	ten_2, ok := rb.Get(10)
	c.Check(ok, Equals, true)
	c.Check(ten, Equals, ten_2)
}

func (s *MySuite) TestRotateRight(c *C) {
	// we're expecting the 5 to rotate.
	// 5 becomes a child of 3
	// 3 of 10
	// 1 stays to left of 3
	// 4 is left child of 5
	// 6 right child of 5

	rb := NewTree()
	PopulateTree(rb)
	five, _ := rb.Get(5)
	three, _ := rb.Get(3)
	six, _ := rb.Get(6)
	ten, _ := rb.Get(10)


	rb.RotateRight(five)
	c.Check(five.links[0].value, Equals, 4)
	c.Check(five.links[1].value, Equals, 6)
	c.Check(three.parent, Equals, ten)

	// make sure 5 has the new parent, it's old left
	c.Check(five.parent.value, Equals, 3)
	c.Check(six.parent, Equals, five)

}

func (s *MySuite) TestRotateLeft(c *C) {
	rb := NewTree()
	PopulateTree(rb)
	five, _ := rb.Get(5)
	three, _ := rb.Get(3)
	six, _ := rb.Get(6)
	ten, _ := rb.Get(10)

	rb.RotateLeft(five)
	c.Check(five.parent, Equals, six)
	c.Check(three.parent, Equals, five)
	c.Check(six.parent, Equals, ten)
}

func (s *MySuite) TestRotateLeftAtRoot(c *C) {
	rb := NewTree()
	PopulateTree(rb)

	eleven, _ := rb.Get(11)
	ten, _ := rb.Get(10)

	rb.RotateLeft(ten)
	c.Check(eleven, Equals, rb.root)

}
func (s *MySuite) TestRotateRightAtRoot(c *C) {
	rb := NewTree()
	PopulateTree(rb)

	five, _ := rb.Get(5)
	ten, _ := rb.Get(10)

	rb.RotateRight(ten)
	c.Check(five, Equals, rb.root)

}

func (s *MySuite) TestNot(c *C) {
	c.Check(not(1), Equals, int8(0))
	c.Check(not(0), Equals, int8(1))
}



func (s *MySuite) TestIsBalance(c *C) {
	rb := NewRedBlackTree()
	fmt.Println("Setting up IsBalance")
	PopulateTree(rb)

	c.Check(rb.IsBalanced(), Equals, true)
}


func (s *MySuite) TestSideSideRotation(c *C) {
	rb := NewRedBlackTree()
	rb.Insert(10)
	new_root := rb.Insert(9)
	rb.Insert(8)
	// state of the tree
	// 9 should be the new root
	c.Check(rb.root, Equals, new_root)

}

func (s *MySuite) TestZigZagRotation(c *C) {
	fmt.Println("---- zig zag recolor and rotate test ---- ")
	rb := NewRedBlackTree()
	rb.Insert(10)
	rb.Insert(8)
	new_root := rb.Insert(9)
	// state of the tree
	// 8 should be the new root
	c.Check(rb.root, Equals, new_root)

}

func (s *MySuite) TestLeftRotateInZigZag(c *C) {
	fmt.Println("---- check left rotation ---- ")
	rb := NewTree()
	ten := rb.Insert(10)
	eight := rb.Insert(8)
	nine := rb.Insert(9)
	// state of the tree
	// 8 should be the new root
	c.Check(rb.root, Equals, ten)

	rb.RotateLeft(eight)
	c.Check(eight, Equals, nine.links[0])
}


func (s *MySuite) TestLeftLeftRotateRight(c *C) {
	fmt.Println("---- check right rotation ---- ")
	rb := NewTree()
	ten := rb.Insert(10)
	nine := rb.Insert(9)
	eight := rb.Insert(8)

	rb.Print()
	rb.RotateRight(ten)

	c.Check(rb.root, Equals, nine)
	c.Check(rb.root.links[0], Equals, eight)
	c.Check(rb.root.links[1], Equals, ten)
}

func (s *MySuite) TestBlackCount(c *C) {
	t := NewTree()

	n1 := t.Insert(10)
	n2 := t.Insert(11)

	n1.red = 0
	n2.red = 1

	_, ok := n1.CountBlack(0)
	c.Check(ok, Equals, true)

}
