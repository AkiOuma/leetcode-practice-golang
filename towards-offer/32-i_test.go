package offer

import (
	"fmt"
	"testing"
)

func TestLevelOrder1(t *testing.T) {
	root := &TreeNode1{
		Val: 3,
	}
	root.Left = &TreeNode1{
		Val: 9,
	}
	root.Right = &TreeNode1{
		Val: 20,
	}
	root.Right.Left = &TreeNode1{
		Val: 15,
	}
	root.Right.Right = &TreeNode1{
		Val: 7,
	}
	res := levelOrder1(root)
	fmt.Printf("%#v\n", res)
	if res[4] != 7 || res[2] != 20 {
		t.Error("unexpected result")
	}
}
