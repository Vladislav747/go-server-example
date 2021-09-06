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

	/* Определение длины массива динамически через троеточие*/
	var todoListArr = [...]string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
	}

	/* Срез*/
	var todoListSlice = []string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
	}

	fmt.Printf("Длина среза todoListSlice: %d\n", len(todoListSlice))
	fmt.Printf("Емкость среза todoListSlice: %d\n", cap(todoListSlice))
	




	fmt.Printf("Кол-во элементов в списке: %d\n", len(todoList))
	fmt.Printf("Кол-во элементов в todoListArr: %d\n", len(todoListArr))
	//Нельзя выйти за пределы массива !! ОШИБКА
	//fmt.Prinln(todoList[4])

}

