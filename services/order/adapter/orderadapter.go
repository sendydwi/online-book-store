package adapter

import (
	apiorder "github.com/sendydwi/online-book-store/api/order"
	apiproduct "github.com/sendydwi/online-book-store/api/product"
	"github.com/sendydwi/online-book-store/services/order/entity"
)

func ConvertToProductSnapshot(detail apiproduct.ProductDetail, stock int, price float32) entity.ProductSnapshot {
	return entity.ProductSnapshot{
		ISBN:           detail.ISBN,
		Title:          detail.Title,
		Subtitle:       detail.Subtitle,
		Author:         detail.Author,
		PublishTime:    detail.PublishTime,
		Publisher:      detail.Publisher,
		TotalPage:      detail.TotalPage,
		Description:    detail.Description,
		AvailableStock: stock,
		Price:          price,
	}
}

func ConvertToApiOrder(entity entity.Order) apiorder.Order {
	return apiorder.Order{
		ID:              entity.ID,
		UserId:          entity.UserId,
		Status:          entity.Status,
		TotalPrice:      entity.TotalPrice,
		DeliveryAddress: entity.DeliveryAddress,
	}
}

func ConvertToApiOrderItem(entity entity.OrderItem) apiorder.OrderItem {
	return apiorder.OrderItem{
		ID:              entity.ID,
		OrderId:         entity.OrderId,
		ProductId:       entity.ProductId,
		SubtotalPrice:   entity.SubtotalPrice,
		Quantity:        entity.Quantity,
		ProductSnapshot: ConvertToApiProductSnapshot(entity.ProductSnapshot),
	}
}

// Adapter function to convert OriginalProductSnapshot to SimplifiedProductSnapshot
func ConvertToApiProductSnapshot(entity entity.ProductSnapshot) apiorder.ProductSnapshot {
	return apiorder.ProductSnapshot{
		ISBN:           entity.ISBN,
		Title:          entity.Title,
		Subtitle:       entity.Subtitle,
		Author:         entity.Author,
		PublishTime:    entity.PublishTime,
		Publisher:      entity.Publisher,
		TotalPage:      entity.TotalPage,
		Description:    entity.Description,
		AvailableStock: entity.AvailableStock,
		Price:          entity.Price,
	}
}
