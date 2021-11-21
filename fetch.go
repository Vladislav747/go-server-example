package main

// Fetch выводит ответ на запрос по заданному URL
import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	var phrase = "Go Fetch"
	fmt.Println(phrase)

	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch: %v\n", err)
			os.Exit(1)
		}
		//Тут выделяется много памяти для хранения всего ответа буфера
		//b, err := ioutil.ReadAll(resp.Body)
		//Другой способ
		b, err := io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Printf("fetch: Чтение%s: %v\n", url, err)
			os.Exit(1)
		}
		//Выводим то что мы прочитали из Response
		fmt.Printf("%s", b)

	}

}
