package main

import (
	"context"
	"fmt"
	"os"

	"github.com/yanndr/aura/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not connect to backend: %v\n", err)
		os.Exit(1)
	}

	client := pb.NewAuraClient(conn)

	_, err = client.UpdateTemperature(context.Background(), &pb.UpdateTemperatureRequest{10, pb.Unit_CELSIUS})
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not update temperature: %v\n", err)
		os.Exit(1)
	}
}
