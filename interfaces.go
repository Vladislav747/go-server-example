package main

import ("fmt"; "errors")

type employee struct {
	id int
	name string // пол
	age int
	salary int //зарплата 
}

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error 
}

//Обрати внимание что ключи в int
type memoryStorage struct {
	data map[int]employee
}

//Конструктор структуры memoryStorage - 
//вернет тебе ссылку на новую структуру с data мапом внутри
func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
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

type dumbStorage struct{}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (s *dumbStorage) insert(e employee) error {
	fmt.Printf("Вставка пользователы с id:%d прошла успешно\n,", e.id)
	return nil
}

func main() {
	var pharse = "Go Интерфейсы"
	fmt.Println(pharse);


	


	


}

