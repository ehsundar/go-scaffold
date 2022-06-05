package storage

import (
	"context"
	"database/sql"
)

type postgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage(db *sql.DB) Storage {
	return &postgresStorage{
		db: db,
	}
}

func (s *postgresStorage) Select(ctx context.Context, ID int) (*Item, error) {
	row := s.db.QueryRowContext(ctx, "select * from items where id=?", ID)
	if row.Err() != nil {
		return nil, row.Err()
	}

	item := Item{}
	err := row.Scan(&item.ID, &item.Text)
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *postgresStorage) Insert(ctx context.Context, item *Item) (*Item, error) {
	//TODO implement me
	panic("implement me")
}

func (s *postgresStorage) Update(ctx context.Context, item *Item) error {
	//TODO implement me
	panic("implement me")
}

func (s *postgresStorage) Delete(ctx context.Context, ID int) error {
	//TODO implement me
	panic("implement me")
}
