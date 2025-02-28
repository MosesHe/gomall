package service

import (
	"context"

	"github.com/MosesHe/gomall/app/product/biz/dal/mysql"
	"github.com/MosesHe/gomall/app/product/biz/model"
	product "github.com/MosesHe/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	// Finish your business logic.
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	c, err := categoryQuery.GetProductsByCategoryName(req.CategoryName)
	resp = &product.ListProductsResp{}
	for _, v := range c {
		for _, p := range v.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(p.ID),
				Picture:     p.Picture,
				Price:       p.Price,
				Description: p.Description,
				Name:        p.Name,
			})
		}
	}

	return resp, nil
}
