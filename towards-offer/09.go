package offer

// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail
// 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。
// (若队列中没有元素，deleteHead 操作返回 -1 )

// reference: https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof

type CQueue struct {
	data []int
}

func Constructor09() CQueue {
	return CQueue{
		data: make([]int, 0, 10000),
	}
}

func (c *CQueue) AppendTail(value int) {
	c.data = append(c.data, value)
}

func (c *CQueue) DeleteHead() int {
	if len(c.data) == 0 {
		return -1
	}
	head := c.data[0]
	c.data = c.data[1:]
	return head
}
