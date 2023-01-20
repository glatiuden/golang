package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
    var normal []int
    var reverse []int

    var traverse func(node *ListNode)
    traverse = func(node *ListNode) {
        if (node == nil) {
            return
        }

        normal = append(normal, node.Val)
        traverse(node.Next)
        reverse = append(reverse, node.Val)
    }

    traverse(head)
    for i, val := range normal {
        if (val != reverse[i]) {
            return false
        }
    }
    return true
}