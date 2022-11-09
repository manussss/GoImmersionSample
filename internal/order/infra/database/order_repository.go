package database

import (
	"database/sql"
	"gosample/internal/order/domain"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *domain.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price VALUES (?, ?, ?, ?")

	if err != nil {
		return err
	}

	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)

	if err != nil {
		return err
	}

	return nil
}
