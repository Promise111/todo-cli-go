package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

func (storage *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	_, statErr := os.Stat("data")
	if statErr != nil && !os.IsNotExist(statErr) {
		fmt.Println(statErr)
		return statErr
	}

	if os.IsNotExist(statErr) {
		err = os.Mkdir("data", 0755)
		if err != nil {
			return err
		}
	}

	return os.WriteFile("data/"+storage.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile("data/" + s.FileName)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return json.Unmarshal(fileData, data)
}
