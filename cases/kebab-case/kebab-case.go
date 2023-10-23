package kebab

import "github.com/iancoleman/strcase"

func Func(name string) string {
	return strcase.ToKebab(name)
}
