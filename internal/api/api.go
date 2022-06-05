package api

import "context"

type Scaffold interface {
	Create(ctx context.Context, req *CreateRequest) (resp *CreateResponse, err error)
	Retrieve(ctx context.Context, req *RetrieveRequest) (resp *RetrieveResponse, err error)
	Delete(ctx context.Context, req *DeleteRequest) (resp *DeleteResponse, err error)
}

type ItemRepr struct {
	ID   int
	Text string
}

type CreateRequest struct {
	Text string
}

type CreateResponse struct {
	Item *ItemRepr
}

type RetrieveRequest struct {
	ID int
}

type RetrieveResponse struct {
	Item *ItemRepr
}

type DeleteRequest struct {
	ID int
}

type DeleteResponse struct {
}
