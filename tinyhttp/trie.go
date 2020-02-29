package tinyhttp

import "strings"

//参数匹配:。例如 /p/:lang/doc，可以匹配 /p/c/doc 和 /p/go/doc
//通配*。例如 /static/*filepath，可以匹配/static/fav.ico，也可以匹配/static/js/jQuery.js，这种模式常用于静态服务器，能够递归地匹配子路径
type node struct {
	pattern string //待匹配路由
	part string //路由中的一部分
	children []*node //子节点
	isWild bool //是否精确匹配
}

func (n *node)matchChild(part string) *node{
	for _,child :=range n.children {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

func (n *node)findMatchNode(part string) []*node {
	nodes := make([]*node, 0)
	for _,child := range n.children{
		if child.part == part || child.isWild{
			nodes = append(nodes, child)
		}
	}
	return nodes
}

func (n *node)insert(pattern string, parts []string, height int)  {
	if len(parts) == height {
		n.pattern = pattern
		return
	}

	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(pattern, parts, height+1)//递归
}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.findMatchNode(part)

	for _, child := range children {
		result := child.search(parts, height+1)//递归查询
		if result != nil {
			return result
		}
	}

	return nil
}