package handler

import (
	"context"
	"github.com/wuqinqiang/product-srv/proto/product"
	"github.com/wuqinqiang/product-srv/service"
)

var _ product.ProductServer = (*ProductHandler)(nil)

type ProductHandler struct {
	product.UnimplementedProductServer
	server service.ProductServer
}

func NewProductHandler(server service.ProductServer) *ProductHandler {
	return &ProductHandler{server: server}
}

func (p *ProductHandler) GetProductList(ctx context.Context, req *product.GetReq) (*product.GetReply, error) {
	resp := new(product.GetReply)
	list, err := p.server.GetProductList()
	if err != nil {
		return nil, err
	}

	for index := range list {
		productModel := list[index]
		productPb := product.ProductEntity{
			Id:          productModel.Id,
			Title:       productModel.Title,
			Description: productModel.Description,
			Image:       productModel.Image,
			OnSale:      uint32(productModel.OnSale),
			SoldCount:   productModel.SoldCount,
			ReviewCount: productModel.ReviewCount,
			Price:       productModel.Price,
			CreatedAt:   productModel.CreatedAt,
		}

		skus := productModel.Skus
		for j := range skus {
			skuPb := product.ProductSku{
				Id:          skus[j].Id,
				Title:       skus[j].Title,
				Description: skus[j].Description,
				Price:       skus[j].Price,
				Stock:       skus[j].Stock,
				ProductId:   skus[j].ProductId,
				CreatedAt:   skus[j].CreatedAt,
			}
			productPb.Skus = append(productPb.Skus, &skuPb)
		}

		resp.ProductList = append(resp.ProductList, &productPb)
	}
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
