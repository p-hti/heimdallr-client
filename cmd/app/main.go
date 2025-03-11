package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/p-hti/heimdallr-client/internal/broker"
	"github.com/p-hti/heimdallr-client/internal/config"
	"github.com/p-hti/heimdallr-client/internal/manage"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)

	machine, err := manage.NewMachine()
	if err != nil {
		panic(err)
	}
	time.Sleep(5 * time.Second)
	br := broker.NewBrokerWriter(cfg.Broker.KafkaAddress, cfg.Broker.KafkaTopic, machine)
	fmt.Println("already conn")
	errChan := make(chan error)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go br.SendResourceUsage(ctx, errChan)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigChan:
		log.Println("signal received: ", sig)
		cancel()
		time.Sleep(2 * time.Second)
	case err := <-errChan:
		log.Println("error received: ", err)
	}
}
