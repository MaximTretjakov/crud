package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/MaximTretjakov/CRUD/internal/server"
	"github.com/MaximTretjakov/CRUD/pkg/crud"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	// create database connection
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// create grpc server
	s := grpc.NewServer()
	srv := &server.CRUDServer{DbConn: db}
	crud.RegisterCRUDServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
