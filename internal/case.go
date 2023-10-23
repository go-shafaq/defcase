package internal

type Case int

const (
	noDefined Case = iota - 1
	NoCase
	Snak_case
	CamelCase
	Kebab_Case
	PascalCase
	SCREAMING_SNAKE_CASE
)
