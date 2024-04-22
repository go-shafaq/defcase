package defcase

import (
	"github.com/go-shafaq/defcase/internal"
)

type DefCase interface {
	SetCase(tag, pkgPath string, case_ Case)
	GetCase(tag, pkgPath string) Case
	Convert(tag, pkgPath, name string) string
	Converter(tag, pkgPath string) func(name string) string
}

var public DefCase = New()

func Get() DefCase {
	return public
}

func New() DefCase {
	return &defCase{tagNode: internal.NewPkgNode()}
}

type defCase struct {
	tagNode *internal.PkgNode
}

func (d *defCase) SetCase(tag, pkgPath string, case_ Case) {
	d.tagNode.SetCase(tag, pkgPath, case_)
}

func (d *defCase) GetCase(tag, pkgPath string) Case {
	return d.tagNode.GetCase(tag, pkgPath)
}

func (d *defCase) Convert(tag, pkgPath, name string) string {
	converter := d.Converter(tag, pkgPath)
	return converter(name)
}

func (d *defCase) Converter(tag, pkgPath string) func(name string) string {
	return d.tagNode.GetCase(tag, pkgPath).GetFunc()
}
