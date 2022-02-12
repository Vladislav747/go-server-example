package main

import (
	"fmt"
	"os"
	"reflect"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinksl: %v\n", err)
		os.Exit(1)
	}
	//fmt.Printf("%#v", doc)
	xType := fmt.Sprintf("%T", doc)
	fmt.Println(xType) // "[]int"

	//Определение типа
	rt := reflect.TypeOf(doc)

	switch rt.Kind() {
	case reflect.Slice:
		fmt.Println(doc, "is a slice with element type", rt.Elem())
	case reflect.Array:
		fmt.Println(doc, "is an array with element type", rt.Elem())
	case reflect.Map:
		fmt.Println(doc, "is an MAP", rt.Elem())
	default:
		fmt.Println(doc, "is something else entirely")
	}

	//The range keyword is mainly used in for loops in order to iterate over all the elements of a map, slice, channel, or an array.
	//Здесь range потому что visit возвращает массив
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

/*Для обхода HTML узлов
Функция visit обходит дерево узлов HTML, извлекает ссылки из атрибута href
каждого элемента <а href=' добавляет ссылки в срез строк и возвращает
результирующий срез:
*/
func visit(links []string, n *html.Node) []string {
	//Проходимся по всему html дереву
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	//Для вложенных тоже
	//Рекурсия
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links

}
