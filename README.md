# GoCRUD

This is a generic CRUD.


## Example

```go 
package main

import (
	"fmt"

	"github.com/calogxro/gocrud"
)

type Person struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	var crud = gocrud.New[Person]()

	// Save

	id, _ := crud.Save(&Person{
		FirstName: "Elon",
		LastName:  "Musk",
	})

	// FindById

	p, _ := crud.FindById(id)

	fmt.Printf("%s %s #%d", p.FirstName, p.LastName, p.ID)
	// Output: Elon Musk #1
}
```


## Notes

1. `Save()` always returns an ID
2. If an `ID` field is defined on the struct then it gets an ID value

The file `crud_repository_test.go` contains examples.


