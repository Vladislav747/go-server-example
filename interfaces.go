package main

import ("fmt"; "errors")

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

