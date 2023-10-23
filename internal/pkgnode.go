package internal

import (
	"github.com/go-shafaq/defcase/internal/smap"
	"strings"
)

func NewPkgNode() *PkgNode {
	n := new(PkgNode)

	n.Case = noDefined
	n.Dirs = smap.NewMap[PkgNode]()

	return n
}

type PkgNode struct {
	Case Case
	Name string
	Dirs smap.Map[PkgNode]
}

func (n *PkgNode) GetCase(tag, pkgPath string) Case {
	if n == nil {
		return NoCase
	}

	names := SplitTagPath(tag, pkgPath)

	return n.getCase(names)
}

func (n *PkgNode) GetCaseByPkg(pkgPath string) Case {
	if n == nil {
		return NoCase
	}

	names := strings.FieldsFunc(pkgPath, func(r rune) bool {
		return r == '/'
	})

	return n.getCase(names)
}

func (n *PkgNode) getCase(names []string) Case {
	if n == nil {
		return NoCase
	}

	currNode := n
	lastCase := currNode.Case
	for _, name := range names {
		currNode, _ = currNode.Dirs.Get(name)
		if currNode == nil {
			break
		}
		if currNode.Case != noDefined {
			lastCase = currNode.Case
		}
	}

	if lastCase == noDefined {
		return NoCase
	}

	return lastCase
}

func (n *PkgNode) SetCase(tag, pkgPath string, c Case) {
	if n == nil {
		panic("PkgNode struct pointer is nil")
	}

	currNode := n.GetOrSetSubNode(tag)

	if pkgPath == "" || pkgPath == "*" {
		currNode.Case = c
		return
	}

	names := SplitTagPath(tag, pkgPath)

	for _, name := range names {
		currNode = currNode.GetOrSetSubNode(name)
	}

	currNode.Case = c
}

func (n *PkgNode) GetOrSetSubNode(name string) *PkgNode {
	if node, ok := n.Dirs.Get(name); ok {
		return node
	}

	subNode := n.Dirs.GetOrSet(name, NewPkgNode())
	subNode.Name = name

	return subNode
}

func SplitTagPath(tag, path string) []string {
	type span struct {
		start uint16
		end   uint16
	}
	spans := make([]span, 0, 8)

	start := -1
	for end, r := range path {
		if r == '/' {
			if start >= 0 {
				spans = append(spans, span{uint16(start), uint16(end)})
				start = ^start
			}
		} else {
			if start < 0 {
				start = end
			}
		}
	}

	if start >= 0 {
		spans = append(spans, span{uint16(start), uint16(len(path))})
	}

	a := make([]string, len(spans)+1)
	a[0] = tag

	for i, s := range spans {
		a[i+1] = path[s.start:s.end]
	}

	return a
}
