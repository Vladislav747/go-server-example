package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	var phrase = "Go Остановка Контекста"
	fmt.Println(phrase)

	ctx, cancel := context.WithCancel(context.Background())
	go handleSignals(cancel)

	if err := startServer(ctx); err != nil {
		log.Fatal(err)
	}
}

func handleSignals(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal)
	//Запись событий сигналов в канал
	signal.Notify(sigCh, os.Interrupt)
	for {
		sig := <-sigCh
		switch sig {
		//Отлавливаем сигнал остановки ctrl+c
		case os.Interrupt:
			cancel()
			return
		}
	}
}

func startServer(ctx context.Context) error {
	laddr, err := net.ResolveTCPAddr("tcp", ":8080")

	if err != nil {
		return err
	}

	l, err := net.ListenTCP("tcp", laddr)

	if err != nil {
		return err
	}

	defer l.Close()

	for {
		select {
		//Чтение с канала
		case <-ctx.Done():
			log.Println("server stopped")
			return nil

		default:
			//Если эта ошибка эта не Timeout
			if err := l.SetDeadline(time.Now().Add(time.Second)); err != nil {
				return err
			}

			_, err := l.Accept()
			//Если это все таки Timeout то продолжаем читать с канала
			if err != nil {
				if os.IsTimeout(err) {
					continue
				}

				return err
			}

			log.Println("new client connected")
		}
	}
}
