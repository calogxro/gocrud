package gocrud

import (
	"reflect"
)

type CrudRepository[T any] struct {
	entities  map[int]*T
	generator *IDGenerator
}

func New[T any]() CrudRepository[T] {
	return CrudRepository[T]{
		entities:  make(map[int]*T),
		generator: &IDGenerator{},
	}
}

func (r CrudRepository[T]) Save(entity *T) (int, error) {
	id, err := r.generator.Next()
	r.entities[id] = entity

	// TODO
	// if a field of `T` is annotated with `gocrud:"ID"`
	// && the type of the field is `int`
	// then that field is used as ID and has `id` assigned
	// else `id` is only returned

	// if `ID` type is `int` then it gets `id` value
	value := reflect.ValueOf(entity).Elem().FieldByName("ID")
	if value.IsValid() && value.CanInt() {
		value.SetInt(int64(id))
	}

	return id, err
}

func (r CrudRepository[T]) FindById(id int) (*T, error) {
	entity, exists := r.entities[id]
	if !exists {
		return nil, &KeyNotFound{}
	}
	return entity, nil
}
