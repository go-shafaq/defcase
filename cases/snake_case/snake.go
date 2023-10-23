package snake

import "github.com/iancoleman/strcase"

func Func(name string) string {
	return strcase.ToSnake(name)
}
