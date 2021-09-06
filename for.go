package main

import ("fmt")

const pi = 3.14

func main(){
	var phrase = "Go For"
	fmt.Println(phrase);

	var todoList = [3]string{
		"купить хлеб",
		"купить молоко",
		"купить пиво",
	}


	fmt.Printf("Кол-во элементов в списке: %d\n", len(todoList))
	

	for index:= range todoList {
		fmt.Printf("%d. %s\n", index, todoList[index])
	}

	for _, item:= range todoList {
		fmt.Printf("%s\n", item)
	}

}

