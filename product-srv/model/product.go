package model

import "gorm.io/gorm"

type ProductState uint8

const (
	ProductInSale ProductState = iota
	ProductNotSale
)

type Product struct {
	Id          uint64       `gorm:"column:id" json:"id"`
	Title       string       `gorm:"column:title" json:"title"`
	Description string       `gorm:"column:description" json:"description"`
	Image       string       `gorm:"column:image" json:"image"`
	OnSale      uint8        `gorm:"column:on_sale" json:"onSale"`
	SoldCount   uint32       `gorm:"column:sold_count" json:"soldCount"`
	ReviewCount uint32       `gorm:"column:review_count" json:"reviewCount"`
	Price       uint32       `gorm:"column:price" json:"price"`
	CreatedAt   string       `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   string       `gorm:"column:updated_at" json:"updatedAt"`
	Skus        []ProductSku `gorm:"foreignKey:product_id"`
}

func (Product) TableName() string {
	return "products"
}

func GetWithOnSale(state ProductState) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("on_sale=?", state)
	}
}
