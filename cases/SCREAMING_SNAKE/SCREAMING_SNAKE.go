package SCREAMING_SNAKE

import "github.com/iancoleman/strcase"

func Func(name string) string {
	return strcase.ToScreamingSnake(name)
}
