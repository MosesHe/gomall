package service

import (
	"context"

	"github.com/MosesHe/gomall/app/order/biz/dal/mysql"
	"github.com/MosesHe/gomall/app/order/biz/model"
	"github.com/MosesHe/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/MosesHe/gomall/rpc_gen/kitex_gen/order"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type ListOrderService struct {
	ctx context.Context
} // NewListOrderService new ListOrderService
func NewListOrderService(ctx context.Context) *ListOrderService {
	return &ListOrderService{ctx: ctx}
}

// Run create note info
func (s *ListOrderService) Run(req *order.ListOrderReq) (resp *order.ListOrderResp, err error) {
	// Finish your business logic.
	list, err := model.ListOrder(s.ctx, mysql.DB, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(500001, err.Error())
	}
	var orders []*order.Order
	for _, v := range list {
		var orderItems []*order.OrderItem
		for _, oi := range v.OrderItems {
			orderItems = append(orderItems, &order.OrderItem{
				Item: &cart.CartItem{
					ProductId: oi.ProductId,
					Quantity:  oi.Quantity,
				},
				Cost: oi.Cost,
			})
		}
		orders = append(orders, &order.Order{
			OrderId:      v.OrderId,
			UserId:       v.UserId,
			UserCurrency: v.UserCurrency,
			Email:        v.Consignee.Email,
			Address: &order.Address{
				StreetAddress: v.Consignee.StreetAddress,
				City:          v.Consignee.City,
				State:         v.Consignee.State,
				ZipCode:       v.Consignee.ZipCode,
				Country:       v.Consignee.Country,
			},
			Items: orderItems,
		})
	}

	resp = &order.ListOrderResp{
		Orders: orders,
	}
	return
}
