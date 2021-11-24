package offer

// 请实现 copyRandomList 函数，复制一个复杂链表。在复杂链表中，
// 每个节点除了有一个 next 指针指向下一个节点，还有一个 random
// 指针指向链表中的任意节点或者 null

// reference: https://leetcode-cn.com/problems/fu-za-lian-biao-de-fu-zhi-lcof/

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var Cache map[*Node]*Node

func CopyNode(node *Node) *Node {
	if node == nil {
		return node
	}
	if copynode, ok := Cache[node]; ok {
		return copynode
	}
	copynode := &Node{Val: node.Val}
	Cache[node] = copynode
	copynode.Next = CopyNode(node.Next)
	copynode.Random = CopyNode(node.Random)
	return copynode
}

func CopyRandomList(head *Node) *Node {
	if Cache == nil {
		Cache = make(map[*Node]*Node)
	}
	return CopyNode(head)
}
