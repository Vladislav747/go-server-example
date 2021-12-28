package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	var phrase = "Go Lesson 4 Web VK"
	fmt.Println(phrase)

	// Bind на порт ОС
	listener, _ := net.Listen("tcp", ":5000")

	for {
		//ждем пока не придёт клиент
		conn, err := listener.Accept()

		//Неуспешное соединение
		if err != nil {
			fmt.Println("Can not connect!!")
			conn.Close()
			continue
		}

		//Все хорошо
		fmt.Println("Can not connect!!")

		//создаем Reader для чтения инф-ции из сокета
		bufReader := bufio.NewReader(conn)
		fmt.Println("Start reading")

		//defer conn.Close()

		for {
			//побайтово читаем
			rbyte, err := bufReader.ReadByte()

			if err != nil {
				fmt.Println("Can not read!", err)
				break
			}

			fmt.Print(string(rbyte))
		}
	}
}
