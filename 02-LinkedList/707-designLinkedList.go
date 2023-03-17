/*
Leetcode:707 设计链表
使用单链表或者双链表，设计并实现自己的链表。
包含以下功能：
	(1)链表结构体定义
	(2)Constructor:链表初始化
	(3)Get:按位查值
	(4)AddAtHead:头插法
	(5)AddAtTail:尾插法
	(6)addAtIndex:按位插值（index从0起）
	(7)deleteAtIndex:按位删除
*/

package main

import (
	"fmt"
)

type MyLinkedList struct {
	cnt       int       //链表长度
	dummyHead *ListNode //虚拟头结点，dummyHead.Next指向真正链表的头结点
}

func Constructor() MyLinkedList {
	return MyLinkedList{0, &ListNode{}}
}

func (this *MyLinkedList) Get(index int) int {
	//获取链表中下标为 index 的节点的值。如果下标无效，则返回 -1 。
	//index从0开始
	if index > this.cnt-1 || index < 0 {
		return -1
	}
	tmp := this.dummyHead
	for ; index >= 0; index-- { //若index=0，则刚好执行一次
		tmp = tmp.Next
	}
	return tmp.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	// 将一个值为 val 的节点插入到链表中第一个元素之前。
	//在插入完成后，新节点会成为链表的第一个节点。
	newHead := NewListNode(val)
	newHead.Next = this.dummyHead.Next
	this.dummyHead.Next = newHead
	this.cnt++
}

func (this *MyLinkedList) AddAtTail(val int) {
	//将一个值为 val 的节点追加到链表中作为链表的最后一个元素。
	newTail := NewListNode(val)
	tmp := this.dummyHead
	for tmp.Next != nil {
		tmp = tmp.Next
	} //此时tmp为旧尾结点
	tmp.Next = newTail //加入新尾结点
	this.cnt++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	//将一个值为 val 的节点插入到链表中下标为 index 的节点之前。（index应该属于[0,cnt]）
	//如果 index 等于链表的长度，那么该节点会被追加到链表的末尾。
	//如果 index 比长度更大，该节点将 不会插入 到链表中。
	if index < 0 || index > this.cnt {
		return
	}
	newNode := NewListNode(val)
	tmp := this.dummyHead
	for ; index > 0; index-- { //若index=1，则移动1次
		tmp = tmp.Next
	} //（index=1）tmp为头结点
	//添加结点
	newNode.Next = tmp.Next
	tmp.Next = newNode
	this.cnt++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	//如果下标有效(index属于[0,cnt-1])，则删除链表中下标为 index 的节点。
	if index < 0 || index > this.cnt-1 {
		return
	}
	tmp := this.dummyHead
	for ; index > 0; index-- { //若index=1，则移动1次
		tmp = tmp.Next
	} //（index=1）tmp移到头结点
	//删除操作
	tmp.Next = tmp.Next.Next
	this.cnt--
}

func (this *MyLinkedList) printMyLinkedList() {
	fmt.Println("\n打印链表：")
	tmp := this.dummyHead
	i := 0
	for tmp.Next != nil {
		fmt.Printf("%v ", tmp.Next.Val)
		tmp = tmp.Next
		i++
	}
	fmt.Printf("\n打印完毕！共有%v个结点。\n\n", i)
}

func main() {
	obj := Constructor()

	obj.AddAtHead(1)
	obj.printMyLinkedList()

	obj.AddAtTail(3)
	obj.printMyLinkedList()

	obj.AddAtIndex(1, 2)
	obj.printMyLinkedList()

	fmt.Println("list[1]:", obj.Get(1))

	obj.DeleteAtIndex(1)
	obj.printMyLinkedList()

	fmt.Println("list[1]:", obj.Get(1))

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
