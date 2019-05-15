# ARTS-Week6

## Algorithm
### problem:Merge Two Sorted Lists
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.
```
Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
```
### solution:
We probably meet lots of situation that there are two or more lists need to be merged.However,this problem is easy to solve.
```golang
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
	    return l2
	}
	if l2 == nil {
	    return l1
	}
	// 定义一个结果节点
	var res *ListNode
	// 当l1节点的值大于l2节点的值，那么res指向l2的节点，从l2开始遍历，反之从l1开始
	if l1.Val >= l2.Val {
		res = l2
		res.Next = mergeTwoLists(l1, l2.Next)
	} else {
		res = l1
		res.Next = mergeTwoLists(l1.Next, l2)
	}
	return res
}
```

## Review

## Tips

## Share
[Git 工作流和分支规范](week-6-share.md)