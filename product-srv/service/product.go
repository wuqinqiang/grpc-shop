package service

import (
	"github.com/wuqinqiang/product-srv/dao"
	"github.com/wuqinqiang/product-srv/model"
)

type ProductServer interface {
	GetProductList() (list []*model.Product, err error)
}

var _ ProductServer = (*ProductServerImpl)(nil)

type ProductServerImpl struct {
	productDao dao.ProductDao
}

func NewProductServerImpl(dao dao.ProductDao) ProductServer {
	return &ProductServerImpl{productDao: dao}
}

func (s *ProductServerImpl) GetProductList() (list []*model.Product, err error) {
	return s.productDao.GetProductList()
}
