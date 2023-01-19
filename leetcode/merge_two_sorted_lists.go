package merge_two_sorted_lists

type ListNode struct {
   	Val int
	Next *ListNode
}

// Iterative
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    merged := new(ListNode)
    curr := merged
    
    for list1 != nil || list2 != nil {
        if list2 == nil || (list1 != nil && list1.Val < list2.Val) {
            curr.Next = list1
            list1 = list1.Next
        } else {
            curr.Next = list2
            list2 = list2.Next
        }
        curr = curr.Next
    }
    return merged.Next
}

// Recursive
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
    merged := new (ListNode)

    if (list1 == nil) {
        return list2
    }

    if (list2 == nil) {
        return list1
    }

    if list1.Val < list2.Val {
        merged.Val = list1.Val
        merged.Next = mergeTwoLists(list1.Next, list2)
    } else {
        merged.Val = list2.Val
        merged.Next = mergeTwoLists(list1, list2.Next)
    }

    return merged
}