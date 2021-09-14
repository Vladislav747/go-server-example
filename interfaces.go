package main

import ("fmt"; "go-server-example/storage/storage")

type dumbStorage struct{}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}

func (s *dumbStorage) insert(e storage.Employee) error {
	fmt.Printf("Вставка пользователы с id:%d прошла успешно\n,", e.Id)
	return nil
}

func main() {
	var pharse = "Go Интерфейсы"
	fmt.Println(pharse);


	ms:= storage.NewMemoryStorage()
	ds:= newDumbStorage()

	spawnEmployees(ms)
	fmt.Println(ms.Get(3))

	spawnEmployees(ds)

}

func spawnEmployees(s storage.Storage) {
	for i:=1; i <= 10; i ++ {
		s.Insert(storage.Employee{Id:i})
	}
}

