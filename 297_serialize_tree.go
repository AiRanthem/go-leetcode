package leetcode

import (
	"strconv"
	"strings"
)

type Codec struct {
	N int
	D []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	c.D = []string{}
	c.serializeActor(root)
	return strings.Join(c.D, ",")
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	c.N = 0
	c.D = strings.Split(data, ",")
	return c.deserailizeActor()
}

func (c *Codec) deserailizeActor() *TreeNode {
	d := c.D[c.N]
	if d == "" {
		return nil
	}
	val, _ := strconv.Atoi(d)
	node := &TreeNode{Val: val}
	c.N++
	node.Left = c.deserailizeActor()
	c.N++
	node.Right = c.deserailizeActor()
	return node
}

func (c *Codec) serializeActor(node *TreeNode) {
	if node == nil {
		c.D = append(c.D, "")
	} else {
		c.D = append(c.D, strconv.Itoa(node.Val))
		c.serializeActor(node.Left)
		c.serializeActor(node.Right)
	}
}
