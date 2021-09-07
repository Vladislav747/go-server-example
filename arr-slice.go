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

	fmt.Printf("Длина  массива todoList: %d\n", len(todoList))

	/* Определение длины массива динамически через троеточие*/
	var todoListArr = [...]string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
	}

	fmt.Printf("Длина динамичного массива todoListArr: %d\n", len(todoListArr))

	//Нельзя выйти за пределы массива !! ОШИБКА
	//fmt.Prinln(todoList[4])

	/* Срез*/
	var todoListSlice = []string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
	}

	fmt.Printf("Длина среза todoListSlice: %d\n", len(todoListSlice))
	fmt.Printf("Емкость среза todoListSlice: %d\n", cap(todoListSlice))
	
	// Функция append() не изменяет
	// переданный в аргументах срез, а
	// возвращает новый.
	newTodoListSlice := append(todoListSlice, "постирать носки")

	fmt.Printf("Длина среза newTodoListSlice: %d\n", len(newTodoListSlice))
	fmt.Printf("Емкость среза newTodoListSlice: %d\n", cap(newTodoListSlice))

	// 	Синтаксис [m:n] позволяет делать
	// выборку элементов массива или
	// среза от индекса m до n . 
	tasksSlice := newTodoListSlice[1:3]

	for index:= range tasksSlice {
		fmt.Printf("Усеченный слайс %d. %s\n", index, tasksSlice[index])
	}

	// Пустые скобки [:] означают выборку всех элементов.
	tasksSliceFull := newTodoListSlice[:]

	for index:= range tasksSliceFull {
		fmt.Printf("Полный слайс %d. %s\n", index, tasksSliceFull[index])
	}
}

