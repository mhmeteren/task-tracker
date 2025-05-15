package repository

import (
	"regexp"
	"task-tracker/internal/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to start sqlmock: %s", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn: db,
	})

	gormDB, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		t.Fatalf("failed to initialize gorm DB: %s", err)
	}

	return gormDB, mock
}

func TestUserRepository_FindByID(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewUserRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "email"}).
		AddRow(1, "Bob", "Bob@example.com")

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users" WHERE "users"."id" = $1 AND "users"."deleted_at" IS NULL ORDER BY "users"."id" LIMIT $2`)).
		WithArgs(1, 1).
		WillReturnRows(rows)

	user, err := repo.FindByID(1)
	if err != nil {
		t.Fatalf("an error occurred: %s", err)
	}

	if user.ID != 1 || user.Name != "Bob" || user.Email != "Bob@example.com" {
		t.Errorf("unexpected user result: %+v", user)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unmet expectations: %s", err)
	}
}

func TestUserRepository_Create(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewUserRepository(db)

	user := &model.User{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "hashed-password",
		RoleID:   0,
	}

	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "users"`)).
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			nil,
			"Alice",
			"alice@example.com",
			"hashed-password",
			nil,
			sqlmock.AnyArg(),
			uint(0),
		).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	err := repo.Create(user)
	if err != nil {
		t.Fatalf("failed to create user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met: %s", err)
	}
}

func TestUserRepository_Update(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewUserRepository(db)

	user := &model.User{
		Name:     "Updated Alice",
		Email:    "alice@example.com",
		Password: "new-password",
		RoleID:   0,
	}
	user.ID = 1

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET`)).
		WithArgs(
			sqlmock.AnyArg(),
			sqlmock.AnyArg(),
			nil,
			"Updated Alice",
			"alice@example.com",
			"new-password",
			nil,
			time.Time{},
			uint(0),
			uint(1),
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectCommit()

	err := repo.Update(user)
	if err != nil {
		t.Fatalf("failed to update user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met: %s", err)
	}
}

func TestUserRepository_Delete(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := NewUserRepository(db)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "users" SET "deleted_at"=$1 WHERE "users"."id" = $2 AND "users"."deleted_at" IS NULL`)).
		WithArgs(sqlmock.AnyArg(), 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	err := repo.Delete(1)
	if err != nil {
		t.Fatalf("failed to delete user: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("expectations were not met: %s", err)
	}
}
