package controller

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/shafizod/rabbitmqgo/internal/rabbitmq"
)

func Start(rabbit rabbitmq.Service) {
	msgs, err := rabbit.Subscribe("golang-queue", "merchants")
	if err != nil {
		panic("Failed to register a consumer: " + err.Error())
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	<-sigint
}
