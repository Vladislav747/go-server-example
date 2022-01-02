//Server2 - минимальный "echo"-сервер со счетчиком запросов,
package main

//У сервера имеется два обработчика, и запрашиваемый URL определяет, какой из
//них будет вызван: запрос /c o u n t вызывает c o u n te r, а все прочие — h a n d le r. Сервер
//запускает обработчик для каждого входящего запроса в отдельной go-подпрограмме,
//так что несколько запросов могут обрабатываться одновременно. Однако если два параллельных запроса попытаются обновить счетчик c o u n t в один и тот же момент
//времени, он может быть увеличен не согласованно; в такой программе может возникнуть серьезная ошибка под названием состояние гонки (race condition; см. раздел 9.1).

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Обработчик, возвращающий компонент пути запрашиваемого URL.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	mu.Lock()
	count++
	mu.Unlock()
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	//Для считывания get параметров
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

// Счетчик, возвращающий количество сделанных запросов,
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}
