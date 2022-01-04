package gee

import (
	"fmt"
	"strings"
)


// 树节点结构
type node struct {
	pattern  string  // 完整路由？
	part     string  // 当前路由部分
	children []*node  // 子节点
	isWild   bool  // 是否模糊匹配 part 含有 :/* 时为true
}


// 打印节点数据
func (n *node) String() string {
	return fmt.Sprintf("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}


// 路由注册
func (n *node) insert(pattern string, parts []string, height int) {
	// 到达完整路由段
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]

	// 查询匹配的子节点
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}

	// 子节点继续匹配
	child.insert(pattern, parts, height + 1)
}


// 路由检索
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)
	for _, child := range children {
		result := child.search(parts, height + 1)
		if result != nil {
			return result
		}
	}

	return nil
}




func (n *node) travel(list *([]*node)) {
	if n.pattern != "" {
		*list = append(*list, n)
	}
	for _, child := range n.children {
		child.travel(list)
	}
}


// 第一个匹配成功的节点，用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}



// 所有匹配成功的节点，用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
