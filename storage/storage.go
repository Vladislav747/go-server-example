package storage

import (
	"errors"
)

type Employee struct {
	Id int
	Name string // пол
	Age int
	Salary int //зарплата 
}

type Storage interface {
	Insert(e Employee) error
	Get(id int) (Employee, error)
	Delete(id int) error 
}

//Обрати внимание что ключи в int
type MemoryStorage struct {
	data map[int]Employee
}

//Конструктор структуры MemoryStorage - 
//вернет тебе ссылку на новую структуру с data мапом внутри
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]Employee),
	}
}

//Изменение по ссылке
//Обрати внимание что это объявление методов 
//func (s structName) methodName(arg1 type, arg2 type ...) returnType {}
func (s *MemoryStorage) Insert(e Employee) error {
	s.data[e.Id] = e

	return nil
}

//Обрати внимание что это объявление метода get структуры MemoryStorage
//func (s structName) methodName(arg1 type, arg2 type ...) returnType {}
func (s *MemoryStorage) Get(id int) (Employee, error) {
	e, exists := s.data[id]
	if !exists {
		return Employee{}, errors.New("employee with such id doesn't exist")
	}

	return e, nil
}

//Обрати внимание что это объявление метода get структуры MemoryStorage
//func (s structName) methodName(arg1 type, arg2 type ...) returnType {}
func (s *MemoryStorage) Delete(id int)  error {
	delete(s.data, id)

	return nil
}