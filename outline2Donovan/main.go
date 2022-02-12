package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	//Map с записями кол-ва элементов
	counts := make(map[string]int)
	result := outline2(nil, doc, counts)
	//
	fmt.Println("map:", result)

}

//Выводит тип html элемента
func outline2(stack []string, n *html.Node, counts map[string]int) map[string]int {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // Внесение дескриптора в стек
		//Показывать текущий стэк
		// fmt.Println(stack)
		//Проверяем что ключ с названием элемента уже есть
		i, ok := counts[n.Data]
		//Если ключа нет то создаем со значением 1 иначе добавляем единицу к имеющимися кол-ву
		if !ok {
			counts[n.Data] = 1
		} else {
			counts[n.Data] = i + 1
		}

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//Рекурсия
		outline2(stack, c, counts)
	}
	return counts
}
