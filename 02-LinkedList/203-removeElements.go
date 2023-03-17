/*
Leetcode:203 移除链表元素
https://leetcode.cn/problems/remove-linked-list-elements/
给你一个链表的头节点 head 和一个整数 val ，
请你删除链表中所有满足 ListNode.Val == val 的节点，并返回新的头节点。
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func createLinkedList() *ListNode {
	// 创建节点
	node1 := NewListNode(1)
	node2 := NewListNode(2)
	node3 := NewListNode(6)
	node4 := NewListNode(3)
	node5 := NewListNode(4)
	node6 := NewListNode(5)
	node7 := NewListNode(6)
	// 构建链表
	node1.Next = node2
	node2.Next = node3
	node3.Next = node4
	node4.Next = node5
	node5.Next = node6
	node6.Next = node7
	//默认尾结点的next就是nil
	return node1
}

func printLinkedList(head *ListNode) {
	//遍历打印链表中所有元素的值
	dummyHead := NewListNode(0)
	dummyHead.Next = head
	fmt.Println("打印链表:")
	tmp := dummyHead
	i := 0
	for tmp.Next != nil {
		fmt.Printf("%v ", tmp.Next.Val)
		i++
		tmp = tmp.Next
	}
	fmt.Printf("\n打印完毕！共有%v个结点。\n", i)
}

func removeElements(head *ListNode, val int) *ListNode {
	//虚拟头结点方法：统一删除逻辑
	dummyHead := NewListNode(0)
	dummyHead.Next = head
	tmp := dummyHead
	for tmp != nil && tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return dummyHead.Next
}

func printLinkedList2(head *ListNode) {
	//遍历打印链表中所有元素的值
	if head == nil {
		return
	}
	tmp := head //用于遍历
	i := 1
	fmt.Println("打印链表:")
	for tmp.Next != nil {
		fmt.Printf("%v ", tmp.Val)
		i++
		tmp = tmp.Next
	}
	//别忘了尾结点（tmp.Next==nil）的操作
	fmt.Printf("%v\n", tmp.Val)
	fmt.Printf("打印完毕！共有%v个结点。\n", i)
	return
}

func removeElements2(head *ListNode, val int) *ListNode {
	//常规方法：先删除头结点，再删除其他节点
	//删除头结点
	for head != nil && head.Val == val {
		//要确保删除后的头结点已经满足条件（不需再删）,因此需要循环
		head = head.Next
	}
	//【注意】：上面删除头结点就有可能已经删光光了，所以下面的tmp有可能是nil
	//nil没有Next，所以要【先判断tmp!=nil】，否则tmp.Next会报错

	tmp := head
	//删除中间结点或尾结点
	for tmp != nil && tmp.Next != nil { //要先判断 tmp != nil，否则会访问空指针
		if tmp.Next.Val == val { //该删
			tmp.Next = tmp.Next.Next
		} else { //不该删，就访问下一结点
			tmp = tmp.Next
		}
	}
	return head
}

/*
func main() {
	head := createLinkedList() //创建好一个单链表[1,2,6,3,4,5,6]
	printLinkedList(head)
	head = removeElements(head, 6) //删除数值为6的结点
	printLinkedList(head)
}
*/
