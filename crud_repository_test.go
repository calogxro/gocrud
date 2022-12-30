package gocrud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	firstName string
	lastName  string
}

func TestCrud(t *testing.T) {
	var crud = NewCrudRepository[Person](NewIntGenerator())

	person := Person{
		firstName: "Elon",
		lastName:  "Musk",
	}

	// Save

	id, err := crud.Save(&person)

	assert.Equal(t, 1, id)
	assert.Nil(t, err)

	// FindById

	want := person
	found, err := crud.FindById(id)

	assert.Equal(t, want, *found)
	assert.Nil(t, err)
}
