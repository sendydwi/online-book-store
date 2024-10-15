package cart_test

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	testutils "github.com/sendydwi/online-book-store/commons"
	"github.com/sendydwi/online-book-store/services/cart"
	"github.com/sendydwi/online-book-store/services/cart/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_GetCurrentActiveCart_Repository(t *testing.T) {
	db, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}

	repo := &cart.CartRepository{DB: db}

	userId := "user123"
	cart := entity.Cart{UserId: userId, Status: entity.CartStatusActive}

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "carts" WHERE user_id = $1 AND status=$2 ORDER BY "carts"."cart_id" LIMIT $3`)).
			WithArgs(userId, entity.CartStatusActive, 1).
			WillReturnRows(sqlmock.NewRows([]string{"user_id", "status"}).AddRow(cart.UserId, cart.Status))

		result, err := repo.GetCurrentActiveCart(userId)
		assert.NoError(t, err)
		assert.Equal(t, &cart, result)
	})

	t.Run("database_error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "carts" WHERE user_id = $1 AND status=$2 ORDER BY "carts"."cart_id" LIMIT $3`)).
			WithArgs(userId, entity.CartStatusActive).
			WillReturnError(gorm.ErrRecordNotFound)

		result, err := repo.GetCurrentActiveCart(userId)
		assert.Error(t, err)
		assert.Nil(t, result)
	})

}

func Test_CreateActiveCart_Repository(t *testing.T) {
	db, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}

	repo := &cart.CartRepository{DB: db}

	cart := entity.Cart{
		CartId:    "cart-id",
		UserId:    "user123",
		Status:    entity.CartStatusActive,
		CreatedAt: time.Now(),
		CreatedBy: "system",
		UpdatedAt: time.Now(),
		UpdatedBy: "system",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "carts"`).
			WithArgs(
				cart.CartId,
				cart.UserId,
				cart.Status,
				testutils.AnyTime{},
				cart.CreatedBy,
				testutils.AnyTime{},
				cart.UpdatedBy).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = repo.CreateActiveCart(cart)
		assert.NoError(t, err)
	})

	t.Run("database_error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "carts"`).WithArgs(cart.UserId, cart.Status).WillReturnError(errors.New("insert error"))
		mock.ExpectRollback()

		err = repo.CreateActiveCart(cart)
		assert.Error(t, err)
	})

}

func Test_UpdateCartItem_Repository(t *testing.T) {
	db, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}

	repo := &cart.CartRepository{DB: db}

	cartItem := entity.CartItem{
		CartItemId: 1,
		CartId:     "cart123",
		ProductId:  1,
		Quantity:   1,
		CreatedAt:  time.Now(),
		CreatedBy:  "system",
		UpdatedAt:  time.Now(),
		UpdatedBy:  "system",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "cart_items" WHERE cart_id = $1 AND product_id = $2 ORDER BY "cart_items"."cart_item_id" LIMIT $3 FOR UPDATE`)).
			WithArgs(cartItem.CartId, cartItem.ProductId, 1).
			WillReturnRows(sqlmock.NewRows([]string{"cart_item_id", "cart_id", "product_id"}).
				AddRow(cartItem.CartItemId, cartItem.CartId, cartItem.ProductId))

		mock.ExpectExec(`UPDATE "cart_items" SET`).
			WithArgs(
				cartItem.CartId,
				cartItem.ProductId,
				cartItem.Quantity,
				testutils.AnyTime{},
				cartItem.CreatedBy,
				testutils.AnyTime{},
				cartItem.UpdatedBy,
				cartItem.CartItemId,
			).WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = repo.UpdateCartItem(cartItem)
		assert.NoError(t, err)
	})

	t.Run("database_error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "cart_items" WHERE cart_id = $1 AND product_id = $2 ORDER BY "cart_items"."cart_item_id" LIMIT $3 FOR UPDATE`)).
			WithArgs(cartItem.CartId, cartItem.ProductId, 1).
			WillReturnError(errors.New("select error"))
		mock.ExpectRollback()

		err = repo.UpdateCartItem(cartItem)
		assert.Error(t, err)
	})

}

func Test_GetCartItemByCartId_Repository(t *testing.T) {
	db, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}

	repo := &cart.CartRepository{DB: db}

	cartId := "cart123"
	cartItems := []entity.CartItem{
		{CartId: cartId, ProductId: 1},
		{CartId: cartId, ProductId: 2},
	}
	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "cart_items" WHERE cart_id = $1`)).
			WithArgs(cartId).
			WillReturnRows(sqlmock.NewRows([]string{"cart_id", "product_id"}).
				AddRow(cartItems[0].CartId, cartItems[0].ProductId).
				AddRow(cartItems[1].CartId, cartItems[1].ProductId))

		result, err := repo.GetCartItemByCartId(cartId)
		assert.NoError(t, err)
		assert.Equal(t, cartItems, result)
	})

	t.Run("database_error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "cart_items" WHERE cart_id = $1`)).
			WithArgs(cartId).
			WillReturnError(errors.New("query error"))

		result, err := repo.GetCartItemByCartId(cartId)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func Test_UpdateCartStatus_Repository(t *testing.T) {
	db, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}

	repo := &cart.CartRepository{DB: db}

	cart := entity.Cart{
		CartId:    "cart123",
		UserId:    "user-id",
		Status:    entity.CartStatusOrdered,
		CreatedAt: time.Now(),
		CreatedBy: "system",
		UpdatedAt: time.Now(),
		UpdatedBy: "system",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "carts" SET`).
			WithArgs(
				cart.UserId,
				cart.Status,
				testutils.AnyTime{},
				cart.CreatedBy,
				testutils.AnyTime{},
				cart.UpdatedBy,
				cart.CartId,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = repo.UpdateCartStatus(cart)
		assert.NoError(t, err)
	})

	t.Run("database_error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "carts" SET`).
			WithArgs(
				cart.UserId,
				cart.Status,
				testutils.AnyTime{},
				cart.CreatedBy,
				testutils.AnyTime{},
				cart.UpdatedBy,
				cart.CartId,
			).
			WillReturnError(errors.New("update error"))
		mock.ExpectRollback()

		err = repo.UpdateCartStatus(cart)
		assert.Error(t, err)
	})
}
