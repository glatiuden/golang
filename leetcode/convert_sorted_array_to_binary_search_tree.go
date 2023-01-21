package convert_sorted_array_to_binary_search_tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
    return createBST(nums, 0, len(nums) -1)
}

func createBST(nums []int, l, r int) *TreeNode {
    if l > r {
        return nil
    }

    mid := (l+r)/2
    node := &TreeNode{
        Val: nums[mid],
        Left: createBST(nums, l, mid-1),
        Right: createBST(nums, mid+1, r),
    }
    return node
}