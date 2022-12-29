package crud

import (
	"testing"
)

type person struct {
	firstName string
	lastName  string
}

func TestCrud(t *testing.T) {
	var g IDGenerator[int] = &IntGenerator{}

	var crud = NewCrudRepository[person](g)
	id, err := crud.Save(&person{
		firstName: "Elon",
		lastName:  "Musk",
	})

	if id != 1 {
		t.Fatal("id should be 1")
	}

	if err != nil {
		t.Fatal("error should be nil")
	}
}
