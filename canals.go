package main

import ("fmt"; "time")

func main(){
	//Каналы это механизм коммуникации между горутинами
	var phrase = "Каналы"
	fmt.Println(phrase);


	//Канал по своей структуре схож со срезом или мапой и тоже
	//является ссылочным типом.

	//ch := make(chan int) // инициализация канала
	//ch <- 1 // запись в канал
	//number := <-ch // чтение из канала
	//
	//fmt.Printf("Данные канала: %d\n", number)

	result1 := make(chan int)
	result2 := make(chan int)


	//fmt.Println("Время исполнения без горутинов");
	//Очень сложные вычисления
	go calculateSomething12(1000, result1)
	//fmt.Println("Время исполнения с горутинами");

	//Использование корутин через слово go
	go calculateSomething12(2000, result2)




	
}

func calculateSomething12(n int, res chan int){
	t := time.Now()

	result := 0
	for i:=0; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Время выполнения расчетов %s\n",  time.Since(t))
	res <- result
}


