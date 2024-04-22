package internal

import (
	"github.com/iancoleman/strcase"
)

type Case int

const (
	noDefined            Case = iota - 1
	NoCase                    //	ID Name UserId IHopeItWillBeHelpful
	Snak_case                 //	id name user_id i_hope_it_will_be_helpful
	CamelCase                 //	id name userId iHopeItWillBeHelpful
	Kebab_Case                //	id name user-id i-hope-it-will-be-helpful
	PascalCase                //	Id Name UserId IHopeItWillBeHelpful
	SCREAMING_SNAKE_CASE      //	ID NAME USER_ID I_HOPE_IT_WILL_BE_HELPFUL
)

func (c Case) GetFunc() func(name string) string {
	switch c {
	case Snak_case:
		return strcase.ToSnake
	case CamelCase:
		return strcase.ToLowerCamel
	case Kebab_Case:
		return strcase.ToKebab
	case PascalCase:
		return strcase.ToCamel
	case SCREAMING_SNAKE_CASE:
		return strcase.ToScreamingSnake
	}

	return func(name string) string { return name }
}
