package main

import ("fmt"; "time")

func main(){
	var phrase = "Goрутины"
	fmt.Println(phrase);


	t:= time.Now() //время старта

	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339))

	//Горутины это по сути потоки в go
	fmt.Println("Сравните время исполнения");

	fmt.Println("Время исполнения без горутинов");
	//Очень сложные вычисления
	calculateSomething(1000);
	fmt.Println("Время исполнения с горутинами");



	
}

func calculateSomething(n int){
	t := time.Now()

	result := 0
	for i:=-; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339))
}

