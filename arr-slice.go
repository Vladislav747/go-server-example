package main

import ("fmt")

const pi = 3.14

func main(){
	var phrase = "Go Массивы"
	fmt.Println(phrase);

	var todoList = [3]string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
	}


	fmt.Printf("Кол-во элементов в списке: %d\n", len(todoList))

	//Нельзя выйти за пределы массива !! ОШИБКА
	//fmt.Prinln(todoList[4])

}

