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
最近阿里云推了篇文章，一看确实不错。

[架构整洁之道](https://yq.aliyun.com/articles/702344?tk=ZUyGvJlxXjLRCYX4mXHISRgZjacO3YwM%2F5DA2qmvyj4%3D)

文章首先大概介绍了软件系统的价值，阐明软件系统的目标。基于此，提出三种经典的编程范式，以及在各种层面（架构设计，类和代码，组件，处理组件依赖）上的设计原则。最后确定了下架构工作的基本方针。值得细读。

## Tips
Git操作分支命令（持续更新）

删除分支  `git branch -d <branchname>`

查看远程分支  `git branch -r`

删除远程分支：
```git
git branch -r -d origin/branch-name
git push origin :branch-name
```
## Share
[Git 工作流和分支规范](week-6-share.md)