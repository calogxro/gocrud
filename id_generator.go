package crud

import "github.com/google/uuid"

type IDGenerator[I comparable] interface {
	Next() (I, error)
}

type StringGenerator struct {
	ids map[string]string
}

func (g StringGenerator) Next() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	_, exists := g.ids[id.String()]
	if exists {
		return "", &KeyExists{}
	}
	return id.String(), nil
}

type IntGenerator struct {
	counter int
}

func (g *IntGenerator) Next() (int, error) {
	g.counter++
	return g.counter, nil
}
