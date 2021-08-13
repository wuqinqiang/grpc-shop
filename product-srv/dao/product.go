package dao

import (
	"github.com/wuqinqiang/product-srv/model"
	"gorm.io/gorm"
)

type ProductDao interface {
	GetProductList() (list []*model.Product, err error)
	CreateProduct(product *model.Product) (id int64, err error)
	UpdateProduct(product *model.Product) error
	FirstProductById(id int64) (*model.Product, error)
	DeleteProductById(id int64) (count int64, err error)
	ListingProductById(id int64) (count int64, err error)
	DeListingProductById(id int64) (count int64, err error)
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

func (p *ProductImpl) CreateProduct(product *model.Product) (id int64, err error) {
	err = p.db.Model(&model.Product{}).Create(product).Error
	id = product.Id
	return
}

func (p *ProductImpl) UpdateProduct(product *model.Product) error {
	return p.db.Model(&model.Product{}).Updates(product).Error
}

func (p *ProductImpl) DeleteProductById(id int64) (count int64, err error) {
	base := p.db.Model(&model.Product{}).
		Select("Skus").
		Delete(&model.Product{}, id)
	count = base.RowsAffected
	err = base.Error
	return
}

func (p *ProductImpl) FirstProductById(id int64) (*model.Product, error) {
	var (
		product model.Product
	)
	err := p.db.Model(model.Product{}).First(&product, id).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &product, nil
}

func (p *ProductImpl) ListingProductById(id int64) (count int64, err error) {
	base := p.db.Model(&model.Product{}).
		Where("id=?", id).
		Update("on_sale", model.ProductInSale)
	count = base.RowsAffected
	err = base.Error
	return
}

func (p *ProductImpl) DeListingProductById(id int64) (count int64, err error) {
	base := p.db.Model(&model.Product{}).
		Where("id=?", id).
		Update("on_sale", model.ProductNotSale)
	count = base.RowsAffected
	err = base.Error
	return
}
