package schemas

import (
	"encoding/json"
	"os"
)

type Schemas struct {
	Schemas []Schema `json:"schemas"`
}

type Schema struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func New(filename string) (Schemas, error) {
	scs := Schemas{}
	res := scs.initFrom(filename)
	return scs, res
}

func (s *Schemas) initFrom(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, s)
}

func (s *Schemas) All() []Schema {
	return s.Schemas
}

func (s *Schemas) ByName(name string) (Schema, bool) {
	for _, s := range s.Schemas {
		if s.Name == name {
			return s, true
		}
	}

	return Schema{}, false
}
