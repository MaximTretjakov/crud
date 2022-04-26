package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/MaximTretjakov/CRUD/internal/cache"
	"github.com/MaximTretjakov/CRUD/internal/db_conn"
	"github.com/MaximTretjakov/CRUD/internal/server"
	"github.com/MaximTretjakov/CRUD/pkg/crud"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func serveSwagger(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../swagger/www/swagger.json")
}

func runGWServer(db *sql.DB) error {
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

	// Register Greeter
	gwmux := runtime.NewServeMux()
	err = crud.RegisterCRUDHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	// Swagger-UI
	mux := http.NewServeMux()
	mux.Handle("/", gwmux)
	mux.HandleFunc("/swagger.json", serveSwagger)
	fs := http.FileServer(http.Dir("../swagger/www/swagger-ui"))
	mux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui", fs))

	// GW start
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	err = http.ListenAndServe("localhost:8090", mux)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func main() {
	// create redis cache
	c, err := cache.NewRedisConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// create database connection
	db, err := db_conn.NewDBConnection("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// run gw server
	if err := runGWServer(db); err != nil {
		log.Fatal(err)
	}
}
