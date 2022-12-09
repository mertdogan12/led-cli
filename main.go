package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		select {
		case <-signalChan:
			log.Printf("Got Signal to exit")
			os.Exit(1)
		}
	}()

	run()
}

func run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		text, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		args := strings.Split(text, " ")

		out := sendCommand(args...)

		println(out)
	}
}

func sendCommand(args ...string) string {
	return ""
}
