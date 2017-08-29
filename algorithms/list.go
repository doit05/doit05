package algorithms

import "fmt"

type Node struct {
	Data int
	Next *Node
}

func (node Node) Node(data int) *Node {
	newnode := Node{}
	newnode.Data = data
	newnode.Next = nil
	return &newnode
}

type List struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (list List) List() (newlist List) {
	newlist.Head = nil
	newlist.Tail = nil
	list.Length = 0
	return newlist
}

func (list *List) ListPut(node *Node) {
	//空链表
	if list.Head == nil {
		list.Head = node
		list.Tail = list.Head

		fmt.Printf("创建第一个元素 %d \n", list.Head.Data)
	} else {
		list.Tail.Next = node
		list.Tail = node
		fmt.Printf("创建第%d个元素 %d \n", list.Length, list.Tail.Data)
	}
	list.Length++
	return
}

//遍历列表
func (list *List) ListTravel() {
	tmp := list.Head
	for ; tmp != nil; tmp = tmp.Next {
		fmt.Println(tmp.Data)
	}
}

//遍历列表
func (list *List) ListHasCicle() (bool, int, *Node, *Node) {
	ret := false
	count := 0
	var start, end *Node
	hashset := make([]*Node, list.Length)
	tmp := list.Head
	for count = 0; tmp != nil; tmp = tmp.Next {
		fmt.Println(tmp.Data)
		if In_slice_p(hashset, tmp) < 0 {
			hashset = append(hashset, tmp)
		} else {
			count++
			if count == 1 {
				start = tmp
				ret = true
			} else {
				if tmp.Next == start {
					end = tmp
					return ret, count, start, end
				}
			}

		}
	}
	return ret, count, start, end
}

func In_slice_p(sl []*Node, p *Node) int {
	if p == nil || len(sl) == 0 {
		return -1
	}

	for i := 0; i < len(sl); i++ {
		if sl[i] == p {
			return i + 1
		}
	}
	return -1
}

//查找前驱节点
func (list *List) FindPre(node *Node) (pre *Node) {
	count := 0
	tmp := list.Head
	for ; tmp.Next != nil && tmp.Next != node; tmp = tmp.Next {
		count++
	}
	if count < list.Length {
		pre = tmp
	} else {
		pre = nil
	}
	return
}

func (list *List) FindOne(data int) (ret int, pre *Node) {
	count := 0
	tmp := list.Head
	for ; tmp.Next != nil && tmp.Next.Data != data; tmp = tmp.Next {
		count++
	}
	if count < list.Length {
		pre = tmp
		ret = count + 2
	} else {
		pre = nil
	}
	return
}

func (list *List) DelFirstOne(data int) (ret int, pre *Node) {
	count := 0
	tmp := list.Head
	for ; tmp.Next != nil && tmp.Next.Data != data; tmp = tmp.Next {
		count++
	}
	if count < list.Length {
		pre = tmp
		list.Listdel(pre.Next)
		ret = count + 1
	} else {
		pre = nil
	}
	return
}

//删除下一个节点
func (list *List) Listdel(node *Node) (ret int) {
	if list.Length == 0 { //空链表
		ret = 1
	} else if list.Length == 1 && list.Head == node {
		list.Head = nil
		list.Tail = nil
		list.Length = 0
		ret = 0
	} else if node == list.Head {
		list.Head = node.Next
		list.Length--
	} else if list.Length == 2 && list.Tail == node {
		list.Tail = list.Head
		list.Length--
	} else {
		pre := list.FindPre(node)
		pre.Next = node.Next
		if node == list.Tail {
			list.Tail = pre
		}
	}
	return
}
