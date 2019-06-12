package plutus

import gonanoid "github.com/matoous/go-nanoid"

type idGenerator struct {
	alphabet string
}

func (gen *idGenerator) New() string {
	id, _ := gonanoid.Generate(gen.alphabet, 20)
	return id
}

var ids = &idGenerator{
	alphabet: "abcdefghijklmnopqrstuvwxyz0123456789",
}
