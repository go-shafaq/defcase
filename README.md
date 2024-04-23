# DefCase

## Problem
Tagging structs takes time and sometimes it could be annoying
(probably because we do such easy thing so many times)

```go
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

type Post struct {
	Title       string                `form:"title"`
	Description string                `form:"description"`
	Photo       *multipart.FileHeader `form:"photo"`
}
```

## Solution
```go
package main

import (
	"github.com/go-shafaq/defcase"
	"github.com/go-shafaq/gin"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
}

func main() {
	// Get Default Case
	dcase := defcase.Get()
	// Set a case for "json" and "form" tags
	dcase.SetCase("json", "*", defcase.Snak_case)
	dcase.SetCase("form", "*", defcase.Snak_case)

	// Set a Default Case to library
	gin.SetDefCase(dcase)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, User{
			ID:        1,
			FirstName: "Mahmud",
			LastName:  "Hasan",
			Email:     "MahmudHasan@email.com",
		})
	})
	r.Run()
}
```

### Result of <a href="http://localhost:8080/" target="_blank">localhost:8080</a> request
Pretty-print

```json
{
  "id": 1,
  "first_name": "Mahmud",
  "last_name": "Hasan",
  "email": "MahmudHasan@email.com"
}
```


## Cases

```json
["ID", "Name", "UserId", "IHopeItWillBeHelpful"]
```

| Function               | Result                                      |
|------------------------|---------------------------------------------|
|                        | `ID Name UserId IHopeItWillBeHelpful`       |
| `snak_case`            | `id name user_id i_hope_it_will_be_helpful` |
| `camelCase`            | `id name userId iHopeItWillBeHelpful`       |
| `kebab-case`           | `id name user-id i-hope-it-will-be-helpful` |
| `PascalCase`           | `Id Name UserId IHopeItWillBeHelpful`       |
| `SCREAMING_SNAKE_CASE` | `ID NAME USER_ID I_HOPE_IT_WILL_BE_HELPFUL` |
 
