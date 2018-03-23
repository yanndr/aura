package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"

	"github.com/yanndr/aura"
	"github.com/yanndr/aura/pb"
	"github.com/yanndr/aura/transport"
	"github.com/yanndr/temperature"
)

func main() {

	var config struct {
		Port string `default:":8888"`
	}

	if err := envconfig.Process("", &config); err != nil {
		fmt.Fprintf(os.Stderr, "error loading configuration: %v\n", err)
		os.Exit(1)
	}

	srv := grpc.NewServer()
	s := transport.New(aura.New(temperature.NewWithHandler(0, temperature.Celsius, handler)))
	pb.RegisterAuraServer(srv, s)
	l, err := net.Listen("tcp", config.Port)
	if err != nil {
		log.Fatalf("could not listen to %s %v", config.Port, err)
	}

	go func() {
		// graceful shutdown
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		<-interrupt
		log.Print("app is shutting down...")
		srv.Stop()
	}()

	log.Printf("app is ready to listen and serve on port %s", config.Port)

	if err := srv.Serve(l); err != nil {
		log.Fatalf("server failed %v", err)
	}

	fmt.Println("Good bye.")
}

func handler(t temperature.Temperature) {
	fmt.Printf("Temperature updated to %v\n", t)
}
