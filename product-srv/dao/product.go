package dao

import (
	"github.com/wuqinqiang/product-srv/model"
	"gorm.io/gorm"
)

type ProductDao interface {
	GetProductList() (list []*model.Product, err error)
}

var _ ProductDao = (*ProductImpl)(nil)

type ProductImpl struct {
	db *gorm.DB
}

func NewProductImpl(db *gorm.DB) ProductDao {
	return &ProductImpl{
		db: db,
	}
}

func (p *ProductImpl) GetProductList() (list []*model.Product, err error) {
	err = p.db.Model(&model.Product{}).
		Scopes(model.GetWithOnSale(model.ProductInSale)).
		Preload("Skus", func(db *gorm.DB) *gorm.DB {
			return db.Select(
				"id,title,description,price,stock,product_id,created_at")
		}).
		Find(&list).Error
	return
}

func (p *ProductImpl) CreateProduct() {

}
