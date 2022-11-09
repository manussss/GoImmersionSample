package database

import (
	"database/sql"
	"gosample/internal/order/domain"
	"testing"

	"github.com/stretchr/testify/suite"
)

type OrderRepositoryTestSuit struct {
	suite.Suite
	Db *sql.DB
}

func (suite *OrderRepositoryTestSuit) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")

	suite.NoError(err)
	db.Exec("CREATE TABLE orders (id varchar(255), price float, tax float, final_price float)")
	suite.Db = db
}

func (suit *OrderRepositoryTestSuit) TearDownTest() {
	suit.Db.Close()
}

func TestSuit(t *testing.T) {
	suite.Run(t, new(OrderRepositoryTestSuit))
}

func (suite *OrderRepositoryTestSuit) GivenOrder_WhenSave_ThenShouldHaveOrder() {
	order, err := domain.NewOrder("123", 10, 2)
	suite.NoError(err)
	suite.NoError(order.CalculateFinalPrice())

	repo := NewOrderRepository(suite.Db)
	err = repo.Save(order)
	suite.NoError(err)

	var orderResult domain.Order
	err = suite.Db.QueryRow("SELECT id, price, tax, final_price FROM orders WHERE id = ?", order.ID).Scan(&orderResult.ID, &orderResult.Price, &orderResult.Tax, &orderResult.FinalPrice)

	suite.NoError(err)
	suite.Equal(order.ID, orderResult.ID)
	suite.Equal(order.Price, orderResult.Price)
	suite.Equal(order.Tax, orderResult.Tax)
	suite.Equal(order.FinalPrice, orderResult.FinalPrice)
}
