package storage

import (
	"errors"
	"fmt"
)

type Employee struct {
	Id   int
	Name string
	Age  int
}

type Storage interface {
	insert(e Employee) error
	get(id int) (Employee, error)
	delete(id int) error
}

type MemoryStore struct {
	data map[int]Employee
}

func NewMemoryStorage() *MemoryStore {
	return &MemoryStore{
		data: make(map[int]Employee),
	}
}

func (s *MemoryStore) insert(e Employee) error {
	s.data[e.Id] = e

	return nil
}

func (s *MemoryStore) get(id int) (Employee, error) {
	e, exist := s.data[id]

	if !exist {
		return Employee{}, errors.New("no employee")
	}

	return e, nil
}

func (s *MemoryStore) delete(id int) error {
	delete(s.data, id)

	return nil
}

func (e Employee) GetInfo() string {
	return fmt.Sprintf("name: %v, age %v", e.Name, e.Age)
}

func (e *Employee) SetName(name string) {
	e.Name = name
}
