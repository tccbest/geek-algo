package _6_linkedlist

import "fmt"

//节点定义
type ListNode struct {
	next  *ListNode
	value interface{}
}

//链表定义
type LinkedList struct {
	head   *ListNode
	length uint
}

func (this *ListNode) GetValue() interface{} {
	return this.value
}

func (this *ListNode) GetNext() *ListNode {
	return this.next
}

func NewListNode(v interface{}) *ListNode {
	return &ListNode{nil, v}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(0), 0}
}

//在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *ListNode, v interface{}) bool {
	if nil == p {
		return false
	}

	newNode := NewListNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++

	return true
}

//在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *ListNode, v interface{}) bool {
	if nil == p || p == this.head {
		return false
	}

	//寻找某个节点
	cur := this.head.next
	pre := this.head

	for nil != cur {
		if cur == p {
			break
		}

		pre = cur
		cur = cur.next
	}

	if nil == cur {
		return false
	}

	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++

	return true
}

//在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

//在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	//寻找尾节点
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}

	return this.InsertAfter(cur, v)
}

//通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *ListNode {
	return nil
}

//删除传入的节点
func (this *LinkedList) DeleteNode(p *ListNode) bool {
	return false
}

//打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}

	fmt.Println(format)
}