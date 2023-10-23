package camel

import "github.com/iancoleman/strcase"

func Func(name string) string {
	return strcase.ToLowerCamel(name)
}
