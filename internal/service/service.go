package service

import (
	"context"

	"github.com/MaximTretjakov/CRUD/pkg/crud"
)

type CRUDServer struct{}

func (c *CRUDServer) CreateNote(ctx context.Context, r *crud.Request) (*crud.Response, error) {
	return &crud.Response{}, nil
}

func (c *CRUDServer) ReadNote(ctx context.Context, r *crud.Request) (*crud.Response, error) {
	return &crud.Response{}, nil
}

func (c *CRUDServer) UpdateNote(ctx context.Context, r *crud.Request) (*crud.Response, error) {
	return &crud.Response{}, nil
}

func (c *CRUDServer) DeleteNote(ctx context.Context, r *crud.Request) (*crud.Response, error) {
	return &crud.Response{}, nil
}
