package easy

import (
	"fmt"
	"testing"
)

//链表的定义
//链表是一种物理存储单元上非连续、非顺序的存储结构，数据元素的逻辑顺序是通过链表中的指针链接次序实现的。
//链表由一系列结点（链表中每一个元素称为结点）组成，结点可以在运行时动态生成。每个结点包括两个部分：一个是存储数据元素的数据域，另一个是存储下一个结点地址的指针域。
//使用链表结构可以避免在使用数组时需要预先知道数据大小的缺点，链表结构可以充分利用计算机内存空间，实现灵活的内存动态管理。但是链表失去了数组随机读取的优点，同时链表由于增加了结点的指针域，空间开销比较大。
//链表允许插入和移除表上任意位置上的结点，但是不允许随机存取。
//链表有三种类型：单向链表、双向链表、循环链表。

// 定义一个单向链表
type Student struct {
	Name string
	next *Student
}

// 为链表赋值
func createList() {
	var head Student
	head.Name = "张三"

	var stu1 Student
	stu1.Name = "李四"

	head.next = &stu1

	var stu2 Student
	stu2.Name = "王五"

	stu1.next = &stu2

	Req(&head)
}

func Req(tmp *Student) { //tmp指针是指向下一个结构体的地址，加*就是下一个结构体
	for tmp != nil { //遍历输出链表中每个结构体，判断是否为空
		fmt.Println(*tmp)
		tmp = tmp.next //tmp变更为下一个结构体地址
	}
}

func TestCreateList(t *testing.T) {
	createList()
	fmt.Println("Done")
}
