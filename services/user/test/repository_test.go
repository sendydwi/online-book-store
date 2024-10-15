package user_test

import (
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	testutils "github.com/sendydwi/online-book-store/commons"
	"github.com/sendydwi/online-book-store/services/user"
	"github.com/sendydwi/online-book-store/services/user/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_RegisterUser_Repository(t *testing.T) {
	db, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}

	repo := &user.UserRepository{DB: db}

	user := entity.User{
		UserId:    "user-id",
		Email:     "email@example.com",
		Password:  "this_password",
		CreatedAt: time.Now(),
		CreatedBy: "system",
		UpdatedAt: time.Now(),
		UpdatedBy: "system",
	}

	t.Run("success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO \"users\"").
			WithArgs(
				user.UserId,
				user.Email,
				user.Password,
				testutils.AnyTime{},
				user.CreatedBy,
				testutils.AnyTime{},
				user.UpdatedBy).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = repo.RegisterUser(user)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO \"users\"").
			WithArgs(
				user.UserId,
				user.Email,
				user.Password,
				testutils.AnyTime{},
				user.CreatedBy,
				testutils.AnyTime{},
				user.UpdatedBy).
			WillReturnError(errors.New("failed to insert user"))
		mock.ExpectRollback()

		err = repo.RegisterUser(user)
		assert.Error(t, err)
		assert.Equal(t, "failed to insert user", err.Error())
	})
}

func Test_GetUserByEmai_Repository(t *testing.T) {
	db, mock, err := testutils.MockDB()
	if err != nil {
		t.Fatalf("failed to set up mock db: %v", err)
	}

	repo := &user.UserRepository{DB: db}

	rows := sqlmock.NewRows([]string{"user_id", "email", "password"}).
		AddRow("user-id", "email@example.com", "this_password")

	t.Run("success", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."user_id" LIMIT $2`)).
			WithArgs("email@example.com", 1).
			WillReturnRows(rows)

		user, err := repo.GetUserByEmail("email@example.com")
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, "user-id", user.UserId)
	})

	t.Run("error", func(t *testing.T) {
		mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE email = $1 ORDER BY "users"."user_id" LIMIT $2`)).
			WithArgs("email@example.com", 1).
			WillReturnError(gorm.ErrRecordNotFound)

		user, err := repo.GetUserByEmail("email@example.com")
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})

}
