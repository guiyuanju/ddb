package model

import (
	"ddb/storage/schemas"

	"golang.org/x/mod/sumdb/storage"
)

type Adder interface {
	Add(sc schemas.Schema, data any) error
}

type Remover interface {
	Remove(uuid int) error
}

type Saver interface {
	Save() error
}

type Aller interface {
	All() ([]any, error)
}

type Filterer interface {
	Filter(func(any) any) []any
}

type Model struct {
	storage any
}

func (m Model) Add(sc schemas.Schema, data any) error {
	return storage.Add()
}
