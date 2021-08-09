package model

type ProductSku struct {
	Id          uint64 `gorm:"column:id" json:"id"`
	Title       string `gorm:"column:title" json:"title"`
	Description string `gorm:"column:description" json:"description"`
	Price       uint32 `gorm:"column:price" json:"price"`
	Stock       uint32 `gorm:"column:stock" json:"stock"`
	ProductId   uint64 `gorm:"column:product_id" json:"productId"`
	CreatedAt   string `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   string `gorm:"column:updated_at" json:"updatedAt"`
}

func (ProductSku) TableName() string {
	return "product_skus"
}
