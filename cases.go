package defcase

import "github.com/go-shafaq/defcase/internal"

type Case = internal.Case

const (
	NoCase = internal.NoCase
	//			ID Name UserId IHopeItWillBeHelpful
	Snak_case = internal.Snak_case
	//			id name user_id i_hope_it_will_be_helpful
	CamelCase = internal.CamelCase
	//			id name userId iHopeItWillBeHelpful
	Kebab_Case = internal.Kebab_Case
	//			id name user-id i-hope-it-will-be-helpful
	PascalCase = internal.PascalCase
	//			Id Name UserId IHopeItWillBeHelpful
	SCREAMING_SNAKE_CASE = internal.SCREAMING_SNAKE_CASE
	//			ID NAME USER_ID I_HOPE_IT_WILL_BE_HELPFUL
)
