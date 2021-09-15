package main

import ("fmt"; "time")

func main(){
	var phrase = "Goрутины"
	fmt.Println(phrase);


	t:= time.Now() //время старта

	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339))

	go func(){
		for{
			for _, r:=range `-\|/'`{
				fmt.Printf("\r%c", r)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	//Горутины это по сути потоки в go
	fmt.Println("Сравните время исполнения");

	//fmt.Println("Время исполнения без горутинов");
	//Очень сложные вычисления
	go calculateSomething(1000)
	//fmt.Println("Время исполнения с горутинами");

	//Использование корутин через слово go
	go calculateSomething(2000)

	//блокирует выполнение главной горутины,
	//планировщик передает контекст на другие незаблокированные горутины
	time.Sleep(8 * time.Second)
	fmt.Printf("Время выполнения программы: %s\n", time.Since(t))




	
}

func calculateSomething(n int){
	t := time.Now()

	result := 0
	for i:=0; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Результат: %d; Прошло времени %s\n", result, time.Since(t))
}

func calculateSomething1(n int){
	t := time.Now()

	result := 0
	for i:=0; i <= n; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}

	fmt.Printf("Результат2: %d; Прошло времени %s\n", result, time.Since(t))
}


