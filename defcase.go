package defcase

import (
	"github.com/go-shafaq/defcase/internal"
)

var tagNode *internal.PkgNode = internal.NewPkgNode()

func DefCase(tag, pkgPath string, c Case) {
	tagNode.SetCase(tag, pkgPath, c)
}

func ToDefCase(tag, pkgPath, name string) string {
	c := tagNode.GetCase(tag, pkgPath)

	return Convert(name, c)
}

func DefCaseFn(tag string) func(pkgPath, name string) string {
	pkgNode := tagNode.GetOrSetSubNode(tag)

	return func(pkgPath, name string) string {
		c := pkgNode.GetCaseByPkg(pkgPath)

		return Convert(name, c)
	}
}
