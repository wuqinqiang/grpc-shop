//+build wireinject

package handler

import (
	"github.com/google/wire"
	"github.com/wuqinqiang/product-srv/dao"
	"github.com/wuqinqiang/product-srv/service"
	"gorm.io/gorm"
)

func InitProductHandler(db *gorm.DB) *ProductHandler {
	panic(wire.Build(
		dao.NewProductImpl,
		service.NewProductServerImpl,
		NewProductHandler,
	))
	return  & ProductHandler{}
}
