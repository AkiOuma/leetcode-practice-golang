package offer

import (
	"fmt"
	"testing"
)

func TestLevelOrder2(t *testing.T) {
	root := &TreeNode2{
		Val: 3,
	}
	root.Left = &TreeNode2{
		Val: 9,
	}
	root.Right = &TreeNode2{
		Val: 20,
	}
	root.Right.Left = &TreeNode2{
		Val: 15,
	}
	root.Right.Right = &TreeNode2{
		Val: 7,
	}
	res := levelOrder2(root)
	fmt.Printf("%v\n", res)

}
