package model

import "ddb/storage/schemas"

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
