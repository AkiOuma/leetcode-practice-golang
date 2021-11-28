package offer

// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

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
//   [9,20],
//   [15,7]
// ]
//

// 提示：

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode2 struct {
	Val   int
	Left  *TreeNode2
	Right *TreeNode2
}

func levelOrder2(root *TreeNode2) [][]int {
	result := make([][]int, 0, 20)
	if root == nil {
		return result
	}
	q := &Queue2{
		value: make([]*TreeNode2, 0, 20),
		size:  0,
	}
	result = append(result, []int{root.Val})
	q.Enqueue(root)
	for !q.Empty() {
		values := make([]int, 0, 20)
		size := q.Size()
		for i := 0; i < size; i++ {
			node := q.Dequeue()
			if node.Left != nil {
				values = append(values, node.Left.Val)
				q.Enqueue(node.Left)
			}
			if node.Right != nil {
				values = append(values, node.Right.Val)
				q.Enqueue(node.Right)
			}
		}
		if len(values) <= 0 {
			continue
		}
		result = append(result, values)
	}
	return result
}

type Queue2 struct {
	size  int
	value []*TreeNode2
}

func (q *Queue2) Dequeue() *TreeNode2 {
	if len(q.value) == 0 {
		return nil
	}
	head := q.value[0]
	q.value = q.value[1:]
	q.size--
	return head
}

func (q *Queue2) Enqueue(node *TreeNode2) {
	q.value = append(q.value, node)
	q.size++
}

func (q *Queue2) Empty() bool {
	return len(q.value) == 0
}

func (q *Queue2) Size() int {
	return q.size
}
