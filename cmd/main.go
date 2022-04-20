package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/MaximTretjakov/CRUD/internal/db_conn"
	"github.com/MaximTretjakov/CRUD/internal/server"
	"github.com/MaximTretjakov/CRUD/pkg/crud"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func runRestServer(db *sql.DB) error {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	srv := &server.CRUDServer{DbConn: db}
	// Attach the Greeter service to the server
	crud.RegisterCRUDServer(s, srv)
	// Serve gRPC server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = crud.RegisterCRUDHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())

	return nil
}

func runGRPCServer(db *sql.DB) error {
	// create grpc server
	s := grpc.NewServer()
	srv := &server.CRUDServer{DbConn: db}
	crud.RegisterCRUDServer(s, srv)

	l, err := net.Listen("tcp", ":9000")
	if err != nil {
		return err
	}

	if err := s.Serve(l); err != nil {
		return err
	}
	return nil
}

func main() {
	// get server type
	serverType := os.Args[1]

	// create database connection
	db, err := db_conn.NewDBConnection("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if serverType == "rest" {
		if err := runRestServer(db); err != nil {
			log.Fatal(err)
		}
	} else {
		if err := runGRPCServer(db); err != nil {
			log.Fatal(err)
		}
	}
}
