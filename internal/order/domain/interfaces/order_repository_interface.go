package interfaces

import "gosample/internal/order/domain"

type OrderRepositoryInterface interface {
	Save(order *domain.Order) error
}
