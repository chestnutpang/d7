package gee

import (
	"fmt"
	"strings"
)


type node struct {
	pattern  string
	part     string
	chidlern []*node
	isWild   bool
}


func (n *node) String() string {
	return fmt.Sprint("node{pattern=%s, part=%s, isWild=%t}", n.pattern, n.part, n.isWild)
}

func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.chidlern = append(n.chidlern, child)
	}
	child.insert(pattern, parts, height + 1)
}


func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchiChildren(part)

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
	for _, child := range n.chidlern {
		child.travel(list)
	}
}


func (n *node) matchChild(part string) *node {
	for _, child := range n.chidlern {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}


func (n *node) matchiChildren(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.chidlern {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}
