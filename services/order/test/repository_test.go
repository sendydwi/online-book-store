package order_test

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	testutils "github.com/sendydwi/online-book-store/commons"
	"github.com/sendydwi/online-book-store/services/order"
	"github.com/sendydwi/online-book-store/services/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrder(t *testing.T) {
	db, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}

	repo := &order.OrderRepository{DB: db}

	order := entity.Order{
		OrderId:         "order123",
		UserId:          "user123",
		Status:          entity.OrderStatusWaitingForPayment,
		TotalPrice:      10,
		DeliveryAddress: "address",
		CreatedAt:       time.Now(),
		CreatedBy:       "system",
		UpdatedAt:       time.Now(),
		UpdatedBy:       "system",
	}
	orderItems := []*entity.OrderItem{
		{
			OrderItemId:     1,
			OrderId:         "order123",
			ProductId:       1,
			SubtotalPrice:   10,
			Quantity:        1,
			ProductSnapshot: entity.ProductSnapshot{},
			CreatedAt:       time.Now(),
			CreatedBy:       "system",
			UpdatedAt:       time.Now(),
			UpdatedBy:       "system",
		},
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "orders"`).
			WithArgs(
				order.OrderId,
				order.UserId,
				order.Status,
				order.TotalPrice,
				order.DeliveryAddress,
				testutils.AnyTime{},
				order.CreatedBy,
				testutils.AnyTime{},
				order.UpdatedBy).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectQuery(`INSERT INTO "order_items" (.+) RETURNING`).WithArgs(
			orderItems[0].OrderId,
			orderItems[0].ProductId,
			orderItems[0].SubtotalPrice,
			orderItems[0].Quantity,
			orderItems[0].ProductSnapshot,
			testutils.AnyTime{},
			orderItems[0].CreatedBy,
			testutils.AnyTime{},
			orderItems[0].UpdatedBy,
			orderItems[0].OrderItemId,
		).WillReturnRows(sqlmock.NewRows([]string{"order_item_id"}).AddRow(1))
		mock.ExpectCommit()

		err = repo.CreateOrder(order, orderItems)
		assert.NoError(t, err)
	})

	t.Run("database_error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "orders"`).WillReturnError(errors.New(`insert order error`))
		mock.ExpectRollback()

		err = repo.CreateOrder(order, orderItems)
		assert.Error(t, err)
	})
}

func TestDeleteOrder(t *testing.T) {
	db, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}

	repo := &order.OrderRepository{DB: db}

	order := entity.Order{OrderId: "order123"}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "order_items"`)).
			WithArgs(order.OrderId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "orders"`)).
			WithArgs(order.OrderId).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = repo.DeleteOrder(order)
		assert.NoError(t, err)
	})

	t.Run("database_error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`DELETE FROM "order_items" WHERE order_id = ?`).WithArgs(order.OrderId).WillReturnError(errors.New(`delete error`))
		mock.ExpectRollback()

		err := repo.DeleteOrder(order)
		assert.Error(t, err)
	})
}

func TestGetOrderById(t *testing.T) {
	db, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}

	repo := &order.OrderRepository{DB: db}

	orderId := "order123"
	order := entity.Order{OrderId: orderId}

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orders" WHERE order_id = $1 ORDER BY "orders"."order_id" LIMIT $2`)).
			WithArgs(orderId, 1).
			WillReturnRows(sqlmock.NewRows([]string{"order_id"}).AddRow(order.OrderId))

		result, err := repo.GetOrderById(orderId)
		assert.NoError(t, err)
		assert.Equal(t, &order, result)
	})

	t.Run("database_error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orders" WHERE order_id = $1 ORDER BY "orders"."order_id" LIMIT $2`)).
			WithArgs(orderId).
			WillReturnError(errors.New(`query error`))

		result, err := repo.GetOrderById(orderId)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func TestGetOrderItemByOrderId(t *testing.T) {
	db, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}

	repo := &order.OrderRepository{DB: db}

	orderId := "order123"
	orderItems := []entity.OrderItem{
		{OrderId: orderId, ProductId: 1},
		{OrderId: orderId, ProductId: 2},
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "order_items" WHERE order_id = $1`)).
			WithArgs(orderId).
			WillReturnRows(sqlmock.NewRows([]string{"order_id", "product_id"}).
				AddRow(orderItems[0].OrderId, orderItems[0].ProductId).
				AddRow(orderItems[1].OrderId, orderItems[1].ProductId))

		result, err := repo.GetOrderItemByOrderId(orderId)
		assert.NoError(t, err)
		assert.Equal(t, orderItems, result)
	})

	t.Run("database_error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "order_items" WHERE order_id = $1`)).
			WithArgs(orderId).
			WillReturnError(errors.New(`query error`))

		result, err := repo.GetOrderItemByOrderId(orderId)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
	// Case 2: Failed to retrieve order items by order ID

}

func TestGetOrderByUserId(t *testing.T) {
	db, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}

	repo := &order.OrderRepository{DB: db}

	userId := "user123"
	orders := []entity.Order{
		{OrderId: "order123", UserId: userId},
		{OrderId: "order124", UserId: userId},
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orders" WHERE user_id = $1`)).
			WithArgs(userId).
			WillReturnRows(sqlmock.NewRows([]string{"order_id", "user_id"}).
				AddRow(orders[0].OrderId, orders[0].UserId).
				AddRow(orders[1].OrderId, orders[1].UserId))

		result, err := repo.GetOrderByUserId(userId)
		assert.NoError(t, err)
		assert.Equal(t, orders, result)
	})

	t.Run("database_error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "orders" WHERE user_id = $1`)).
			WithArgs(userId).
			WillReturnError(errors.New(`query error`))

		result, err := repo.GetOrderByUserId(userId)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
