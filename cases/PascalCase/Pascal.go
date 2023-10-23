package Pascal

import "github.com/iancoleman/strcase"

func Func(name string) string {
	return strcase.ToCamel(name)
}
