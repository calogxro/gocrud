package crud

type CrudRepository[T any, I comparable] struct {
	entities  map[I]*T
	generator IDGenerator[I]
}

func NewCrudRepository[T any, I comparable](generator IDGenerator[I]) CrudRepository[T, I] {
	return CrudRepository[T, I]{
		entities:  make(map[I]*T),
		generator: generator,
	}
}

func (r CrudRepository[T, I]) Save(entity *T) (I, error) {
	ID, err := r.generator.Next()
	r.entities[ID] = entity
	return ID, err
}

func (r CrudRepository[T, I]) FindById(ID I) (*T, error) {
	entity, exists := r.entities[ID]
	if !exists {
		return nil, &KeyNotFound{}
	}
	return entity, nil
}
