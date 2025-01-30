package handler

import (
	"context"

	"grpc-example-with-go/internal/app"
	gen "grpc-example-with-go/internal/handler/grpc/generated"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (h *ProductGrpcHandler) Create(ctx context.Context, in *gen.CreateProductRequest) (*gen.CreateProductResponse, error) {
	p := app.NewProduct(in.GetName())
	h.svc.Add(p)

	return &gen.CreateProductResponse{
		Id:   p.Id,
		Name: p.Name,
	}, nil
}

func (h *ProductGrpcHandler) Delete(ctx context.Context, in *gen.DeleteProductRequest) (*gen.DeleteProductResponse, error) {
	h.svc.Delete(in.Id)

	return &gen.DeleteProductResponse{
		Success: true,
	}, nil
}

func (h *ProductGrpcHandler) Update(ctx context.Context, in *gen.UpdateProductRequest) (*gen.UpdateProductResponse, error) {
	p := &app.Product{
		Id:   in.GetId(),
		Name: in.GetName(),
	}
	h.svc.Update(p)

	return &gen.UpdateProductResponse{
		Success: true,
	}, nil
}

func (h *ProductGrpcHandler) Get(ctx context.Context, in *gen.GetProductRequest) (*gen.GetProductResponse, error) {
	p, err := h.svc.Get(in.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "product not found")
	}

	return &gen.GetProductResponse{
		Success: true,
		Id:      p.Id,
		Name:    p.Name,
	}, nil
}
