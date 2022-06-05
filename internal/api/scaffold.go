package api

import (
	"context"
	"github.com/ehsundar/scaffold/internal/storage"
)

type scaffoldImpl struct {
	s storage.Storage
}

func NewScaffold(s storage.Storage) Scaffold {
	scaffold := &scaffoldImpl{
		s: s,
	}

	return scaffold
}

func (s *scaffoldImpl) Create(ctx context.Context, req *CreateRequest) (resp *CreateResponse, err error) {
	item, err := s.s.Insert(ctx, &storage.Item{
		Text: req.Text,
	})
	if err != nil {
		return nil, err
	}
	return &CreateResponse{Item: &ItemRepr{
		ID:   item.ID,
		Text: item.Text,
	}}, nil
}

func (s *scaffoldImpl) Retrieve(ctx context.Context, req *RetrieveRequest) (resp *RetrieveResponse, err error) {
	//TODO implement me
	panic("implement me")
}

func (s *scaffoldImpl) Delete(ctx context.Context, req *DeleteRequest) (resp *DeleteResponse, err error) {
	//TODO implement me
	panic("implement me")
}
