package gocrud

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Person struct {
	ID        int
	firstName string
	lastName  string
}

type StructWithoutID struct {
	field string
}

func TestCrud(t *testing.T) {
	var crud = New[Person]()

	person := Person{
		firstName: "Elon",
		lastName:  "Musk",
	}

	// Save

	id, err := crud.Save(&person)

	assert.Equal(t, id, person.ID)
	assert.Equal(t, 1, person.ID)
	assert.Nil(t, err)

	// FindById

	expected := person
	found, err := crud.FindById(person.ID)

	assert.Equal(t, 1, expected.ID)
	assert.Equal(t, expected, *found)
	assert.Nil(t, err)
}

func TestCrudWithoutID(t *testing.T) {
	var crud = New[StructWithoutID]()

	object := StructWithoutID{
		field: "something random",
	}

	// Save

	id, err := crud.Save(&object)

	assert.Equal(t, 1, id)
	assert.Nil(t, err)

	// FindById

	expected := object
	found, err := crud.FindById(id)

	assert.Equal(t, expected, *found)
	assert.Nil(t, err)
}
