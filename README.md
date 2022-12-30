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
    var g = gocrud.NewIntGenerator()
    var crud = gocrud.New[Person](g)

	// Save
	id, _ := crud.Save(&Person{
		firstName: "Elon",
		lastName:  "Musk",
	})

	// FindById
	p, _ := crud.FindById(id)

	fmt.Printf("#%d: %s %s", id, p.firstName, p.lastName)
	// Output: #1: Elon Musk
}
```