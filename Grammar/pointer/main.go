package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func Init(v int, n *ListNode) *ListNode {
	return &ListNode{
		Val:  v,
		Next: n,
	}
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	curr := head.Next
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		if prev == head {
			prev.Next = nil
		}
		prev = curr
		curr = next
	}
	return prev
}

func main() {
	e5 := Init(5, nil)
	e4 := Init(4, e5)
	e3 := Init(3, e4)
	e2 := Init(2, e3)
	e1 := Init(1, e2)

	for cur := e1; cur != nil; cur = cur.Next {
		fmt.Printf("Before:%v\n", cur.Val)
	}

	res := reverseList(e1)
	fmt.Printf("test head:%v\n", res.Val)

	for cur := res; cur != nil; cur = cur.Next {
		fmt.Printf("After:%v\n", cur.Val)
	}

	var test int // 默认为0
	var e6 *ListNode //默认为nil
	fmt.Printf("test:%v\t%v\n", test, e6)
}

//func main()  {
//	//指针的常见使用（两符号）
//	a := 10
//	b := &a
//	fmt.Printf("a:%d ptr:%p\n", a, &a)
//	fmt.Printf("b:%p type:%T\n", b, b)
//	fmt.Println(&b)
//	c := *b // 指针取值（根据指针去内存取值）
//	fmt.Printf("type of c:%T\n", c)
//	fmt.Printf("value of c:%v\n", c)
//	//make与new的使用
//	/*这一段为引发panic，因为没分配空间
//	var a *int
//	*a = 100
//	fmt.Println(*a)
//
//	var b map[string]int
//	b["沙河娜扎"] = 100
//	fmt.Println(b)*/
//	var d *int
//	d = new(int)
//	*d = 10
//	fmt.Println(*d)
//
//	var e map[string]int
//	e = make(map[string]int, 10)
//	e["沙河娜扎"] = 100
//	fmt.Println(e)
//
//	//自定义类型基础使用（大写加注释）
//
//	//类型别名
//
//}
