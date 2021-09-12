package main

import ("fmt")

func main(){
	var phrase = "Go Structures"
	fmt.Println(phrase);

	//Описание структуры - но так крайне редко пишут
	employee12 := struct{
		name string
		sex string // пол
		age int
		salary int //зарплата 
	}{
		name: "Вася",
		sex: "M",
		age: 25,
		salary: 1500,
	}


	fmt.Printf("%v\n", employee12)

	//Намного чаще описывают тип так
	type employee struct{
		name string
		sex string // пол
		age int
		salary int //зарплата 
	}

	employee1 := employee{
		name: "Вася",
		sex: "M",
		age: 25,
		salary: 1500,
	}

	employee2 := employee{
		name: "Вася",
		sex: "M",
		age: 28,
		salary: 2000,
	}

	fmt.Printf("%v\n", employee1)
	fmt.Printf("%v\n", employee2)

}

