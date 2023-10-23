package defcase

import "github.com/go-shafaq/defcase/internal"

type Case = internal.Case

const (
	NoCase               = internal.NoCase
	Snak_case            = internal.Snak_case
	CamelCase            = internal.CamelCase
	Kebab_Case           = internal.Kebab_Case
	PascalCase           = internal.PascalCase
	SCREAMING_SNAKE_CASE = internal.SCREAMING_SNAKE_CASE
)
