package algorithms

import "fmt"

type Key int

type BNode struct {
	Leaf     bool
	N        int
	Keys     []Key
	Children []*BNode
}

func (x *BNode) Search(k Key) (n *BNode, idx int) {
	i := 0
	for i < x.N && x.Keys[i] < k {
		i += 1
	}
	if i < x.N && k == x.Keys[i] {
		n, idx = x, i
	} else if x.Leaf == false {
		n, idx = x.Children[i].Search(k)
	}
	return
}

func newBNode(n, branch int, leaf bool) *BNode {
	return &BNode{
		Leaf:     leaf,
		N:        n,
		Keys:     make([]Key, branch*2-1),
		Children: make([]*BNode, branch*2),
	}
}

func (parent *BNode) Split(branch, idx int) { //  idx is Children's index
	full := parent.Children[idx]
	// make a new BNode, copy full's right most to it
	n := newBNode(branch-1, branch, full.Leaf)
	for i := 0; i < branch-1; i++ {
		n.Keys[i] = full.Keys[i+branch]
		n.Children[i] = full.Children[i+branch]
	}
	n.Children[branch-1] = full.Children[2*branch-1] // copy last child
	full.N = branch - 1                              // is half full now, copied to n(new one)
	// shift parent, add new key and children
	for i := parent.N; i > idx; i-- {
		parent.Children[i] = parent.Children[i-1]
		parent.Keys[i+1] = parent.Keys[i]
	}
	parent.Keys[idx] = full.Keys[branch-1]
	parent.Children[idx+1] = n
	parent.N += 1
}

func (tree *Btree) Insert(k Key) {
	root := tree.Root
	if root.N == 2*tree.branch-1 {
		s := newBNode(0, tree.branch, false)
		tree.Root = s
		s.Children[0] = root
		s.Split(tree.branch, 0)
		s.InsertNonFull(tree.branch, k)
	} else {
		root.InsertNonFull(tree.branch, k)
	}
}

func (x *BNode) InsertNonFull(branch int, k Key) {
	i := x.N
	if x.Leaf {
		for i > 0 && k < x.Keys[i-1] {
			x.Keys[i] = x.Keys[i-1]
			i -= 1
		}
		x.Keys[i] = k
		x.N += 1
	} else {
		for i > 0 && k < x.Keys[i-1] {
			i -= 1
		}
		c := x.Children[i]
		if c.N == 2*branch-1 {
			x.Split(branch, i)
			if k > x.Keys[i] {
				i += 1
			}
		}
		x.Children[i].InsertNonFull(branch, k)
	}
}

func space(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += " "
	}
	return s
}

func (x *BNode) String() string {
	return fmt.Sprintf("{n=%d, Leaf=%v, Keys=%v, Children=%v}\n",
		x.N, x.Leaf, x.Keys, x.Children)
}

func (tree *Btree) String() string {
	return tree.Root.String()
}

type Btree struct {
	Root   *BNode
	branch int
}

func New(branch int) *Btree {
	return &Btree{Root: newBNode(0, branch, true), branch: branch}
}
func (tree *Btree) Search(k Key) (n *BNode, idx int) {
	return tree.Root.Search(k)
}
