package gocrud

type IDGenerator struct {
	nextId int
}

func (g *IDGenerator) Next() (int, error) {
	g.nextId++
	return g.nextId, nil
}
