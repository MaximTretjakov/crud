package server

import (
	"context"
	"database/sql"

	"github.com/MaximTretjakov/CRUD/internal/kafka"
	"github.com/MaximTretjakov/CRUD/pkg/crud"
)

type CRUDServer struct {
	crud.UnimplementedCRUDServer
	DbConn *sql.DB
}

func (c *CRUDServer) CreateNote(ctx context.Context, r *crud.CreateRequest) (*crud.Status, error) {
	_, err := c.DbConn.Exec("insert into blog (title, text, tags) values ($1, $2, $3)", r.Title, r.Text, r.Tags)
	if err != nil {
		return &crud.Status{Status: "500 Internal error"}, err
	}
	return &crud.Status{Status: "200 Ok"}, nil
}

func (c *CRUDServer) ReadNote(ctx context.Context, r *crud.ReadRequest) (*crud.Response, error) {
	resp := crud.Response{}
	rows := c.DbConn.QueryRow("select id, title, text, tags from blog where id = $1", r.Id)
	rows.Scan(&resp.Id, &resp.Title, &resp.Text, &resp.Tags)
	resp.Status = "200 Ok"
	return &resp, nil
}

func (c *CRUDServer) UpdateNote(ctx context.Context, r *crud.UpdateRequest) (*crud.Status, error) {
	_, err := c.DbConn.Exec("update blog set title = $1, text = $2, tags = $3 where id = $4", r.Title, r.Text, r.Tags, r.Id)
	if err != nil {
		return &crud.Status{Status: "500 Internal error"}, err
	}
	return &crud.Status{Status: "200 Ok"}, nil
}

func (c *CRUDServer) DeleteNote(ctx context.Context, r *crud.DeleteRequest) (*crud.Status, error) {
	_, err := c.DbConn.Exec("delete from blog where id = $1", r.Id)
	if err != nil {
		return &crud.Status{Status: "500 Internal error"}, err
	}
	return &crud.Status{Status: "200 Ok"}, nil
}

func (c *CRUDServer) KafkaSendMessage(ctx context.Context, r *crud.KafkaRequest) (*crud.KafkaResponse, error) {
	if err := kafka.SendMessageToQueue("comments", []byte(r.Request)); err != nil {
		return nil, err
	}
	return &crud.KafkaResponse{Response: "200 Ok"}, nil
}

func (c *CRUDServer) KafkaReadMessage(ctx context.Context, r *crud.KafkaRequest) (*crud.KafkaResponse, error) {
	message, err := kafka.ReadMessageFromQueue("comments")
	if err != nil {
		return nil, err
	}
	return &crud.KafkaResponse{Response: message + "200 Ok"}, nil
}
