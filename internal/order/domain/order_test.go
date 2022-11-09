package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func GivenAnEmptyID_WhenCreateNewOrder_ThenShouldReturnError(t *testing.T) {
	order := Order{}

	assert.Error(t, order.IsValid(), "invalid id")
}

func GivenAnEmptyPrice_WhenCreateNewOrder_ThenShouldReturnError(t *testing.T) {
	order := Order{ID: "123"}

	assert.Error(t, order.IsValid(), "invalid price")
}

func GivenAnEmptyTax_WhenCreateNewOrder_ThenShouldReturnError(t *testing.T) {
	order := Order{ID: "123", Price: 10}

	assert.Error(t, order.IsValid(), "invalid price")
}

func GivenValidParams_WhenCreateNewOrder_ThenShouldReturnOrder(t *testing.T) {
	order := Order{ID: "123", Price: 10, Tax: 10}

	assert.Equal(t, "123", order.ID)
	assert.Nil(t, order.IsValid())
}

func GivenValidParams_WhenCreateNewOrderFunc_ThenShouldReturnOrder(t *testing.T) {
	order, err := NewOrder("123", 10, 10)

	assert.Nil(t, err)
	assert.Equal(t, "123", order.ID)
}

func GivenPriceAndTax_WhenCalculatePrice_ThenShouldSetFinalPrice(t *testing.T) {
	order, err := NewOrder("123", 10, 10)

	order.CalculateFinalPrice()

	assert.Nil(t, err)
	assert.Equal(t, 20, order.FinalPrice)
}
