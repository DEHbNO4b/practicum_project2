package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/DEHbNO4b/practicum_project2/internal/storage"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	db *sql.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sql.Open("pgx", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s %w", op, err)
	}
	return &Storage{db: db}, nil
}

func (s *Storage) SaveUser(ctx context.Context, login string, paskHash []byte) (int64, error) {

	op := "storage.postgres.SaveUser"

	_, err := s.db.Exec("INSERT INTO users(login,pass_hash)VALUES($1,$2)", login, paskHash)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == `23505` {
				return 0, fmt.Errorf("%s %w", op, storage.ErrUserExists)
			}

			return 0, fmt.Errorf("%s %w", op, err)
		}
	}

	// id, err := res.LastInsertId()
	// if err != nil {
	// 	return 0, fmt.Errorf("%s %w", op, err)
	// }

	return 0, nil
}

func (s *Storage) User(ctx context.Context, login string) (models.User, error) {
	op := "storage.postgres.User"

	row := s.db.QueryRowContext(ctx, "SELECT id, login, pass_hash FROM users WHERE login = $1", login)

	var user User
	err := row.Scan(&user.Id, &user.Login, &user.PassHash)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, fmt.Errorf("%s %w", op, storage.ErrUserNotFound)
		}
		return models.User{}, fmt.Errorf("%s %w", op, err)
	}
	if err := row.Err(); err != nil {
		return models.User{}, fmt.Errorf("%s %w", op, err)
	}

	u := userToDomain(user)

	return u, nil
}
