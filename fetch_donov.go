package main

// Fetch выводит ответ на запрос по заданному URL
import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	var phrase = "Go Fetch"
	fmt.Println(phrase)

	for _, url := range os.Args[1:] {
		urlFormatted := formatUrl(url)
		resp, err := http.Get(urlFormatted)
		if err != nil {
			fmt.Printf("fetch: %v\n", err)
			os.Exit(1)
		}
		//Тут выделяется много памяти для хранения всего ответа буфера
		//b, err := ioutil.ReadAll(resp.Body)
		//Другой способ 1.8
		b, err := io.Copy(os.Stdout, resp.Body)

		resp.Body.Close()
		if err != nil {
			fmt.Printf("fetch: Чтение%s: %v\n", url, err)
			os.Exit(1)
		}
		//Выводим то что мы прочитали из Response
		fmt.Printf("%s\n", b)
		fmt.Printf("Статус ответа сервера %s", resp.Status)

	}
}

//Добавить строку 'http://' если таковой нет - 1.9
func formatUrl(url string) string {
	if strings.HasPrefix(url, "http://") {
		return url
	} else {
		return "http://" + url
	}
}
