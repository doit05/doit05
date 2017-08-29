package study


import (
		"fmt"
)

type Node struct {
	name string
	value       int
	child_list []*Node
}

func (root *Node) FindNode(nodea *Node) (ret []*Node, find bool) {
	ret = make([]*Node,1)
	ret = append(ret,root)
	find = false
	if(in_array(nodea, root.child_list)){
		find = true
		return
	}
	if(root.child_list != nil && len(root.child_list) > 0){
		for _, val := range root.child_list {
		ret1, find_tmp := x.FindNode(val, nodea)
		if(find_tmp){
			ret = apped(ret,ret1)
    }
		}
	}
	
	
}

func in_array(val *Node, array []*Node) (exists bool, index int) {
    exists = false
    index = -1;

    for i, v := range array {
        if val.name == v.name {
            index = i
            exists = true
            return
        }   
    }
    return
}




func (tree *Node) Create() (root *Node){
	null_child_list := make([]*Node,0) 
	noder := &Node{
		value : 1,
		name : "r"
		child_list : null_child_list
	}
	nodes := &Node{
		value : 1,
		name : "s"
		child_list : null_child_list
	}
	nodet := &Node{
		value : 1,
		name : "t"
		child_list : null_child_list
	}
	nodek := &Node{
		value : 1,
		name : "k"
		child_list : null_child_list
	}
	nodel := &Node{
		value : 1,
		name : "l"
		child_list : null_child_list
	}
	n_list := make([]*Node, 3)
	n_list = append(n_list,noder,nodes,nodet)
	noden := &Node{
		value : 1,
		name : "n"
		child_list : n_list
	}
	nodeo := &Node{
		value : 1,
		name : "o"
		child_list : null_child_list
	}
	nodep := &Node{
		value : 1,
		name : "p"
		child_list : null_child_list
	}
	nodeq := &Node{
		value : 1,
		name : "q"
		child_list : null_child_list
	}
	nodee := &Node{
		value : 1,
		name : "e"
		child_list : null_child_list
	}
	f_list := make([]*Node, 3)
	f_list = append(n_list,nodek,nodel,noden)
	nodef := &Node{
		value : 1,
		name : "f"
		child_list : f_list
	}
	nodeg := &Node{
		value : 1,
		name : "g"
		child_list : null_child_list
	}
	h_list := make([]*Node, 3)
	h_list = append(n_list,nodeo,nodep,nodeq)
	nodeh := &Node{
		value : 1,
		name : "h"
		child_list : h_list
	}
	nodei := &Node{
		value : 1,
		name : "i"
		child_list : null_child_list
	}
	nodej := &Node{
		value : 1,
		name : "j"
		child_list : null_child_list
	}
	b_list := make([]*Node, 3)
	b_list = append(n_list,nodee,nodef,nodeg)
	nodeb := &Node{
		value : 1,
		name : "b"
		child_list : b_list
	}
	nodec := &Node{
		value : 1,
		name : "c"
		child_list : null_child_list
	}
	d_list := make([]*Node, 3)
	d_list = append(n_list,nodeh,nodei,nodej)
	noded := &Node{
		value : 1,
		name : "d"
		child_list : d_list
	}
	a_list := make([]*Node, 3)
	a_list = append(n_list,nodeb,nodec,noded)
	nodea := &Node{
		value : 1,
		name : "a"
		child_list : a_list
	}

	return nodea
}




func main() {
		root := Node.Create()
		nodec := &Node{
		value : 1,
		name : "c"
		child_list : nil
	}
	ret,find := root.FindNode(nodea)
	fmt.Println(ret)
	fmt.Println(find)
		
}
