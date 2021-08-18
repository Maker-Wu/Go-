package main

import "fmt"

type BinaryTree struct {
	root *HeroNode
	len  int
}

func (b BinaryTree) preOrder() {
	if b.root == nil {
		fmt.Println("二叉树为空，无法遍历")
	}
	b.root.preOrder()
}

func (b BinaryTree) infixOrder() {
	if b.root == nil {
		fmt.Println("二叉树为空，无法遍历")
	}
	b.root.infixOrder()
}

func (b BinaryTree) postOrder() {
	if b.root == nil {
		fmt.Println("二叉树为空，无法遍历")
	}
	b.root.postOrder()
}

func (b BinaryTree) preOrderSearch(no int) {
	if b.root == nil {
		fmt.Println("根节点为空")
		return
	}
	res := b.root.preOrderSearch(no)
	if res != nil {
		fmt.Printf("找到了，%d -> %s\n", res.no, res.name)
	} else {
		fmt.Println("没有找到")
	}
}

func (b BinaryTree) infixOrderSearch(no int) {
	if b.root == nil {
		fmt.Println("根节点为空")
		return
	}
	res := b.root.infixOrderSearch(no)
	if res != nil {
		fmt.Printf("%d -> %s\n", res.no, res.name)
	} else {
		fmt.Println("没有找到")
	}
}

func (b BinaryTree) postOrderSearch(no int) {
	if b.root == nil {
		fmt.Println("根节点为空")
		return
	}
	res := b.root.postOrderSearch(no)
	if res != nil {
		fmt.Printf("%d -> %s\n", res.no, res.name)
	} else {
		fmt.Println("没有找到")
	}
}

// 递归删除节点
/*
如果删除的节点是叶子节点，则删除该节点
如果删除的节点是非叶子节点，则删除该子树
*/
func (b *BinaryTree) delNode(no int) {
	if b.root == nil {
		fmt.Println("空树，无法删除")
		return
	}
	b.root.delNode(no, &b.root)
}

type HeroNode struct {
	no    int
	name  string
	left  *HeroNode
	right *HeroNode
}

// 前序遍历
func (h *HeroNode) preOrder() {
	if h == nil {
		return
	}
	fmt.Printf("HeroNode [no=%d, name=%s]\n", h.no, h.name)
	h.left.preOrder()
	h.right.preOrder()
}

// 中序遍历
func (h *HeroNode) infixOrder() {
	if h == nil {
		return
	}
	h.left.infixOrder()
	fmt.Printf("HeroNode [no=%d, name=%s]\n", h.no, h.name)
	h.right.infixOrder()
}

// 后序遍历
func (h *HeroNode) postOrder() {
	if h == nil {
		return
	}
	h.left.postOrder()
	h.right.postOrder()
	fmt.Printf("HeroNode [no=%d, name=%s]\n", h.no, h.name)

}

// 前序查找
func (h *HeroNode) preOrderSearch(no int) *HeroNode {
	if h == nil {
		return nil
	}
	if h.no == no {
		return h
	}
	res := h.left.preOrderSearch(no)
	if res != nil {
		return res
	}
	res = h.right.preOrderSearch(no)
	if res != nil {
		return res
	}
	return nil
}

// 中序查找
func (h *HeroNode) infixOrderSearch(no int) *HeroNode {
	if h == nil {
		return nil
	}
	res := h.left.infixOrderSearch(no)
	if res != nil {
		return res
	}
	if h.no == no {
		return h
	}
	res = h.right.infixOrderSearch(no)
	if res != nil {
		return res
	}
	return nil
}

// 后序查找
func (h *HeroNode) postOrderSearch(no int) *HeroNode {
	if h == nil {
		return nil
	}
	res := h.left.postOrderSearch(no)
	if res != nil {
		return res
	}
	res = h.right.postOrderSearch(no)
	if res != nil {
		return res
	}
	if h.no == no {
		return h
	}
	return nil
}

// 删除节点
func (h *HeroNode) delNode(no int, parent **HeroNode) {
	if h.no == no {
		*parent = nil
	}
	if h.left != nil {
		h.left.delNode(no, &h.left)
	}
	if h.right != nil {
		h.right.delNode(no, &h.right)
	}
}

func main() {
	binaryTree := BinaryTree{}
	root := HeroNode{no: 1, name: "宋江"}
	hero2 := HeroNode{no: 2, name: "卢俊义"}
	hero3 := HeroNode{no: 3, name: "吴用"}
	hero4 := HeroNode{no: 4, name: "公孙胜"}
	hero5 := HeroNode{no: 5, name: "关胜"}
	root.left = &hero2
	root.right = &hero3
	hero3.left = &hero5
	hero3.right = &hero4
	binaryTree.root = &root
	binaryTree.preOrder()
	binaryTree.infixOrder()
	binaryTree.postOrder()

	binaryTree.postOrderSearch(4)
	binaryTree.delNode(3)
	binaryTree.preOrder()

}
