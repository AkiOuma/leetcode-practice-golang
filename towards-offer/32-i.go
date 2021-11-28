package offer

// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。

//

// 例如:
// 给定二叉树: [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// 返回：

// [3,9,20,15,7]
//

// 来源：力扣（LeetCode）
// 链接：https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof
// 著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func levelOrder1(root *TreeNode1) []int {
	q := &Queue1{
		value: make([]*TreeNode1, 0, 10),
	}
	result := make([]int, 0, 20)
	if root == nil {
		return result
	}
	q.Enqueue(root)
	for !q.Empty() {
		node := q.Dequeue()
		result = append(result, node.Val)
		if node.Left != nil {
			q.Enqueue(node.Left)
		}
		if node.Right != nil {
			q.Enqueue(node.Right)
		}
	}
	return result
}

type Queue1 struct {
	value []*TreeNode1
}

func (q *Queue1) Dequeue() *TreeNode1 {
	if len(q.value) == 0 {
		return nil
	}
	head := q.value[0]
	q.value = q.value[1:]
	return head
}

func (q *Queue1) Enqueue(node *TreeNode1) {
	q.value = append(q.value, node)
}

func (q *Queue1) Empty() bool {
	return len(q.value) == 0
}
