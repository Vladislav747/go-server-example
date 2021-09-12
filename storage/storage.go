import (
	"fmt"
	"errors"
)

type Employee struct {
	id int
	name string // пол
	age int
	salary int //зарплата 
}

type Storage interface {
	insert(e Employee) error
	get(id int) (Employee, error)
	delete(id int) error 
}

//Обрати внимание что ключи в int
type MemoryStorage struct {
	data map[int]Employee
}

//Конструктор структуры memoryStorage - 
//вернет тебе ссылку на новую структуру с data мапом внутри
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[int]Employee),
	}
}

//Изменение по ссылке
//Обрати внимание что это объявление методов 
//func (s structName) methodName(arg1 type, arg2 type ...) returnType {}
func (s *memoryStorage) insert(e employee) error {
	s.data[e.id] = e

	return nil
}

//Обрати внимание что это объявление метода get структуры memoryStorage
//func (s structName) methodName(arg1 type, arg2 type ...) returnType {}
func (s *memoryStorage) get(id int) (employee, error) {
	e, exists := s.data[id]
	if !exists {
		return employee{}, errors.New("employee with such id doesn't exist")
	}

	return e, nil
}

//Обрати внимание что это объявление метода get структуры memoryStorage
//func (s structName) methodName(arg1 type, arg2 type ...) returnType {}
func (s *memoryStorage) delete(id int)  error {
	delete(s.data, id)

	return nil
}