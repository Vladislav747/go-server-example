package main

import ("fmt")

func main(){
	var phrase = "Go Мапы"
	fmt.Println(phrase);

	//ключ string
	//значение int
	//подобны объектам

	//var ages map[string]int
	ages := make(map[string]int)
	ages["a"] = 20
	ages["b"] = 25
	ages["c"] = 30

	fmt.Println(ages)

	//Синтаксис объявления мапы следующий:
	// var mapName map[keyType]valueType

	newages := map[string]int{
		"a": 20,
		"b": 25,
		"c":30,
	}

	fmt.Println(newages)

	//exists проверка 
	age, exists := newages["a"]
	if !exists {
		fmt.Println("А нет в списке");
	}else{
		fmt.Println("А есть %d в списке", age);
	}

	//Нельзя дублировать ключ
	// newages := map[string]int{
	// 	"a": 20,
	// 	"a": 20,
	// 	"b": 25,
	// 	"c":30,
	// }

	//Но можно свободно менять значение по ключу
	newages["a"] = 123;

	//Через ф-цию delete можно удалять данные из мапы по ключу
	delete(newages, "a");

	fmt.Println(newages, "После удаления ключа a")

	
}

