package leetcode

type LruListNode struct {
	Before *LruListNode
	After  *LruListNode
	Key    int
	Val    int
}

type LRUCache struct {
	M    map[int]*LruListNode
	Head *LruListNode
	Tail *LruListNode
	Cap  int
	Len  int
}

//func Constructor(capacity int) LRUCache {
//	head, tail := &LruListNode{}, &LruListNode{}
//	head.After = tail
//	tail.Before = head
//	return LRUCache{
//		M:    make(map[int]*LruListNode, capacity),
//		Head: head,
//		Tail: tail,
//		Cap:  capacity,
//	}
//}

func (c *LRUCache) headInsert(node *LruListNode) {
	node.After = c.Head.After
	node.Before = c.Head
	node.Before.After = node
	node.After.Before = node
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.M[key]; ok {
		before, after := node.Before, node.After
		before.After = after
		after.Before = before
		c.headInsert(node)
		return node.Val
	} else {
		return -1
	}
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.M[key]; ok {
		node.Val = value
		c.Get(key)
	} else {
		node = &LruListNode{Key: key, Val: value}
		c.headInsert(node)
		c.M[key] = node
		if c.Cap == 0 {
			tail := c.Tail.Before
			c.Tail.Before = tail.Before
			tail.Before.After = c.Tail
			delete(c.M, tail.Key)
		} else {
			c.Cap--
		}
	}
}
