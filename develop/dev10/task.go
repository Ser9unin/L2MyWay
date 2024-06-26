package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	mytelnet "github.com/Ser9unin/L2MyWay/develop/dev10/telnet"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	log.SetFlags(0)

	timeout := flag.Duration("timeout", 10*time.Second, "server connect timeout")

	flag.Parse()
	if flag.NArg() != 2 {
		log.Fatal("Please define address and port")
	}

	address := flag.Arg(0) + ":" + flag.Arg(1)

	ctx, cancel := context.WithCancel(context.Background())

	go spectateSignals(cancel)

	client := mytelnet.NewTelnetClient(address, *timeout, os.Stdin, os.Stdout)
	err := client.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	go send(client, cancel)
	go receive(client, cancel)

	<-ctx.Done()
}

func spectateSignals(cancel context.CancelFunc) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	<-signals
	cancel()
}

func send(client mytelnet.TelnetClient, cancel context.CancelFunc) {
	err := client.Send()
	if err != nil {
		log.Println(err)
	}
	cancel()
}

func receive(client mytelnet.TelnetClient, cancel context.CancelFunc) {
	err := client.Receive()
	if err != nil {
		log.Println(err)
	}
	cancel()
}
