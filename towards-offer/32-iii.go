package offer

// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。

//

// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回其层次遍历结果：

// [
//   [3],
//   [20,9],
//   [15,7]
// ]
//

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	level := 0
	result := make([][]int, 0, 20)
	if root == nil {
		return result
	}
	q := &Queue{
		value: make([]*TreeNode, 0, 20),
		size:  0,
	}
	result = append(result, []int{root.Val})
	q.Enqueue(root)
	level++
	for !q.Empty() {
		values := make([]int, 0, 20)
		size := q.Size()
		for i := 0; i < size; i++ {
			node := q.Dequeue()
			if node.Left != nil {
				if level%2 != 0 {
					values = append([]int{node.Left.Val}, values...)
				} else {
					values = append(values, node.Left.Val)
				}
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				if level%2 != 0 {
					values = append([]int{node.Right.Val}, values...)
				} else {
					values = append(values, node.Right.Val)
				}
				q.Enqueue(node.Right)
			}
		}
		if len(values) <= 0 {
			continue
		}
		result = append(result, values)
		level++
	}
	return result
}

type Queue struct {
	size  int
	value []*TreeNode
}

func (q *Queue) Dequeue() *TreeNode {
	if len(q.value) == 0 {
		return nil
	}
	head := q.value[0]
	q.value = q.value[1:]
	q.size--
	return head
}

func (q *Queue) Enqueue(node *TreeNode) {
	q.value = append(q.value, node)
	q.size++
}

func (q *Queue) Empty() bool {
	return len(q.value) == 0
}

func (q *Queue) Size() int {
	return q.size
}
