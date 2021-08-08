package handler

import (
	"context"
	"github.com/wuqinqiang/product-srv/proto/product"
)

var _ product.ProductServer = (*ProductHandler)(nil)

type ProductHandler struct {
	product.UnimplementedProductServer
	//productServer service.ProductServer
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

func (p *ProductHandler) GetProductList(ctx context.Context, req *product.GetReq) (*product.GetReply, error) {
	resp := new(product.GetReply)
	resp.ProductList = append(resp.ProductList, &product.ProductEntity{
		Id:    100,
		Title: "测试商品",
	})
	return resp, nil
}

func (p *ProductHandler) CreateProduct(ctx context.Context, req *product.CreateReq) (*product.CreateReply, error) {
	panic("implement me")
}

func (p *ProductHandler) UpdateProduct(ctx context.Context, req *product.UpdateReq) (*product.UpdateReply, error) {
	panic("implement me")
}

func (p *ProductHandler) DeleteProduct(ctx context.Context, req *product.DeleteReq) (*product.DeleteReply, error) {
	panic("implement me")
}

func (p *ProductHandler) ListingProduct(ctx context.Context, req *product.ListingReq) (*product.ListingReply, error) {
	panic("implement me")
}

func (p *ProductHandler) DeListingProduct(ctx context.Context, req *product.DeListingReq) (*product.DeleteReply, error) {
	panic("implement me")
}
