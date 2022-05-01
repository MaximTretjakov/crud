package main

import (
	"context"
	"fmt"

	"github.com/MaximTretjakov/CRUD/pkg/crud"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func main() {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewBundle().TransportCredentials()),
	}

	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		grpclog.Fatal(err)
	}
	defer conn.Close()

	c := crud.NewCRUDClient(conn)

	request := crud.CreateRequest{
		Title: "Create blog title",
		Text:  "Test create operation",
		Tags:  "CRUD",
	}

	response, err := c.CreateNote(context.Background(), &request)
	if err != nil {
		grpclog.Fatal(err)
	}

	fmt.Println(response.Status)
}
