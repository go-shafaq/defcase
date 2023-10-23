package defcase

import (
	Pascal "github.com/go-shafaq/defcase/cases/PascalCase"
	SCREAMING_SNAKE "github.com/go-shafaq/defcase/cases/SCREAMING_SNAKE"
	camel "github.com/go-shafaq/defcase/cases/camelCase"
	kebab "github.com/go-shafaq/defcase/cases/kebab-case"
	"github.com/go-shafaq/defcase/cases/none"
	snake "github.com/go-shafaq/defcase/cases/snake_case"
)

func Convert(name string, c Case) string {
	switch c {
	case NoCase:
		return none.Func(name)
	case Snak_case:
		return snake.Func(name)
	case CamelCase:
		return camel.Func(name)
	case Kebab_Case:
		return kebab.Func(name)
	case PascalCase:
		return Pascal.Func(name)
	case SCREAMING_SNAKE_CASE:
		return SCREAMING_SNAKE.Func(name)
	}

	return name
}
