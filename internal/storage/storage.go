package storage

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("item not found")
)

type Item struct {
	ID   int
	Text string
}

type Storage interface {
	Select(ctx context.Context, ID int) (*Item, error)
	Insert(ctx context.Context, item *Item) (*Item, error)
	Update(ctx context.Context, item *Item) error
	Delete(ctx context.Context, ID int) error
}
