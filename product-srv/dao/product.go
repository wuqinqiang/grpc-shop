package dao

type ProductDao interface {
}

var _ ProductDao = (*ProductDao)(nil)

type ProductImpl struct {
}

func (p *ProductImpl) GetProductList() {

}

func (p *ProductImpl) CreateProduct() {

}
