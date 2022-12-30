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
	firstName string
	lastName  string
}

func main() {
	var crud = gocrud.NewCrudRepository[Person](gocrud.NewIntGenerator())

	// Save
	id, _ := crud.Save(&Person{
		firstName: "Elon",
		lastName:  "Musk",
	})

	// FindById
	person, _ := crud.FindById(id)

	fmt.Printf("#%d: %s %s", id, person.firstName, person.lastName)
	// Output: #1: Elon Musk
}
```