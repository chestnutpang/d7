package gee

import (
	"fmt"
	"strings"
)


type node struct {
	pattern  string
	part     string
	children []*node
	isWild   bool
}
