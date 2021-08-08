package service

import "github.com/wuqinqiang/product-srv/dao"

type ProductServer interface {
}

var _ ProductServer = (*ProductServerImpl)(nil)

type ProductServerImpl struct {
	productDao dao.ProductDao
}

func NewProductServerImpl(dao dao.ProductDao) ProductServer {
	return &ProductServerImpl{productDao: dao}
}

func (s *ProductServerImpl) GetProductList() {

}
