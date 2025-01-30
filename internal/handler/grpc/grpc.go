package handler

import (
	"context"

	"grpc-example-with-go/internal/app"
	gen "grpc-example-with-go/internal/handler/grpc/generated"
)

type ProductGrpcHandler struct {
	gen.UnimplementedProductHandlerServer
	svc *app.ProductService
}

func NewProductGrpcHandler(svc *app.ProductService) *ProductGrpcHandler {
	return &ProductGrpcHandler{
		svc: svc,
	}
}

func (h *ProductGrpcHandler) Create(ctx context.Context, in *gen.ProductRequest) (*gen.ProductResponse, error) {
	p := app.NewProduct(in.GetName())
	h.svc.Add(p)

	return &gen.ProductResponse{
		Id:   p.Id,
		Name: p.Name,
	}, nil
}
