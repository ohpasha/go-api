package main

import (
	"sync"
)

type Employee struct {
	ID     int    `json:id`
	Name   string `json:name`
	Sex    string `json:sex`
	Age    int    `json:age`
	Salary int    `json:age`
}

type Storage interface {
	Insert(e *Employee)
	Get(id int) (Employee, error)
	Update(id int, e *Employee)
	Delete(id int)
}

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Employee),
		counter: 1,
	}
}

func (s *MemoryStorage) Insert(e *Employee) {
	s.Lock()
	e.ID = s.counter
	s.data[s.counter] = *e

	s.Unlock()
}

func (s *MemoryStorage) Update(id int, e *Employee) {
	s.Lock()
	e.ID = s.counter
	s.data[s.counter] = *e

	s.Unlock()
}

func (s *MemoryStorage) Delete(id int) {
	return
}

func (s *MemoryStorage) Get(id int) (Employee, error) {
	return Employee{
		ID:     1,
		Name:   "s12",
		Sex:    "1212",
		Age:    12,
		Salary: 12,
	}, nil
}
