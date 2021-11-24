package offer

// 定义栈的数据结构，请在该类型中实现一个能够得到栈的最小元素的
//  min 函数在该栈中，调用 min、push 及 pop 的时间复杂度都是 O(1)。

// reference: https://leetcode-cn.com/problems/bao-han-minhan-shu-de-zhan-lcof/

type MinStack struct {
	data [][2]int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: make([][2]int, 0, 1000),
	}
}

func (s *MinStack) Push(x int) {
	if len(s.data) == 0 {
		s.data = append(s.data, [2]int{x, x})
		return
	}
	top := s.data[len(s.data)-1]
	if x < top[1] {
		s.data = append(s.data, [2]int{x, x})
	} else {
		s.data = append(s.data, [2]int{x, top[1]})
	}
}

func (s *MinStack) Pop() {
	size := len(s.data)
	if size > 0 {
		s.data = s.data[0 : size-1]
	}
}

func (s *MinStack) Top() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1][0]
}

func (s *MinStack) Min() int {
	if len(s.data) == 0 {
		return -1
	}
	return s.data[len(s.data)-1][1]
}
