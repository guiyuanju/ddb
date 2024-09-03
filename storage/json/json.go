package jsonstorage

import (
	"ddb/storage/schemas"
	"encoding/json"
	"os"
)

type Storage struct {
	Inner    []Data `json:"data"`
	filename string
}

type Data struct {
	Uuid   int    `json:"uuid"`
	Schema string `json:"schema"`
	Inner  any    `json:"data"`
}

func New(filename string) (Storage, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return Storage{}, err
	}

	s := Storage{nil, filename}
	err = json.Unmarshal(data, &s)

	return s, err
}

func (s *Storage) Add(scm schemas.Schema, data any) error {
	s.Inner = append(s.Inner, Data{0, scm.Name, data})
	return nil
}

func (s *Storage) Remove(uuid int) error {
	var target int
	for i, d := range s.Inner {
		if d.Uuid == uuid {
			target = i
			break
		}
	}
	s.Inner = append(s.Inner[:target], s.Inner[target+1:]...)
	return nil
}

func (s *Storage) Save() error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}

	return os.WriteFile(s.filename, data, os.ModePerm)
}
