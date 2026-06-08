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

func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	// check that data dir exists
	_, dirErr := os.Stat("data")
	if dirErr != nil && !os.IsNotExist(dirErr) {
		return dirErr
	}

	// permissions 0755 (Read/Write/Execute for Owner; Read/Execute for others)
	if os.IsNotExist(dirErr) {
		fmt.Println(dirErr)
		err = os.Mkdir("data", 0755)
		if err != nil {
			return err
		}
	}

	return os.WriteFile("data/"+s.FileName, fileData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	fileData, err := os.ReadFile("data/" + s.FileName)
	if err != nil {
		return err
	}
	return json.Unmarshal(fileData, data)
}