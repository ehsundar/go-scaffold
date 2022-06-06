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

func (svc *scaffoldImpl) Create(ctx context.Context, req *CreateRequest) (resp *CreateResponse, err error) {
	item, err := svc.s.Insert(ctx, &storage.Item{
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

func (svc *scaffoldImpl) Retrieve(ctx context.Context, req *RetrieveRequest) (resp *RetrieveResponse, err error) {
	item, err := svc.s.Select(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &RetrieveResponse{Item: &ItemRepr{
		ID:   item.ID,
		Text: item.Text,
	}}, nil
}

func (svc *scaffoldImpl) Delete(ctx context.Context, req *DeleteRequest) (resp *DeleteResponse, err error) {
	err = svc.s.Delete(ctx, req.ID)
	return &DeleteResponse{}, err
}
