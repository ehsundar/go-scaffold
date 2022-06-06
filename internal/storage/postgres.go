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
	row := s.db.QueryRowContext(ctx, "select id, text from items where id=$1", ID)
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
	row := s.db.QueryRowContext(ctx, "insert into items(text) values ($1) returning id, text", item.Text)
	if row.Err() != nil {
		return nil, row.Err()
	}

	i := Item{}
	err := row.Scan(&i.ID, &i.Text)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (s *postgresStorage) Update(ctx context.Context, item *Item) (*Item, error) {
	row := s.db.QueryRowContext(ctx, "update items set text=$2 where id=$1 returning id, text",
		item.ID, item.Text)
	if row.Err() != nil {
		return nil, row.Err()
	}

	i := Item{}
	err := row.Scan(&i.ID, &i.Text)
	if err != nil {
		return nil, err
	}
	return &i, nil
}

func (s *postgresStorage) Delete(ctx context.Context, ID int) error {
	_, err := s.db.ExecContext(ctx, "delete from items where id=$1", ID)
	return err
}
