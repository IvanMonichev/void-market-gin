package mapper

import (
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/model"
	"github.com/IvanMonichev/void-market-gin/order-svc/internal/transport"
)

func ToOrderRDO(o model.Order) transport.OrderRDO {
	items := make([]transport.OrderItemRDO, len(o.Items))
	for i, item := range o.Items {
		items[i] = transport.OrderItemRDO{
			ID:        item.ID,
			Name:      item.Name,
			Quantity:  item.Quantity,
			UnitPrice: item.UnitPrice,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
	}

	return transport.OrderRDO{
		ID:        o.ID,
		UserID:    o.UserID,
		Status:    string(o.Status),
		Total:     o.Total,
		Items:     items,
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}
