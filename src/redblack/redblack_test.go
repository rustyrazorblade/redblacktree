package redblack

import (
	"testing"
	. "gopkg.in/check.v1"
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
	rb.Insert(10)
	rb.Insert(5)
	rb.Insert(6)
	rb.Insert(3)
	rb.Insert(1)
	rb.Insert(4)
	rb.Insert(11)
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


	five.RotateRight() // rotates on the 5
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

	five.RotateLeft()
	c.Check(five.parent, Equals, six)
	c.Check(three.parent, Equals, five)
	c.Check(six.parent, Equals, ten)
}

func (s *MySuite) TestRotateLeftNoLeaves(c *C) {
	rb := NewTree()
	PopulateTree(rb)
	six, _ := rb.Get(6)

	six.RotateLeft()
}

func (s *MySuite) TestNot(c *C) {
	c.Check(not(1), Equals, 0)
	c.Check(not(0), Equals, 1)
}



func (s *MySuite) TestIsBalance(c *C) {

}
