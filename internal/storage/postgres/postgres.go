package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
	"github.com/DEHbNO4b/practicum_project2/internal/storage"

	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	db   *sql.DB
	once *sync.Once
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

	row := s.db.QueryRowContext(ctx, "INSERT INTO users(login,pass_hash)VALUES($1,$2) returning id", login, paskHash)

	var id int64

	err := row.Scan(&id)
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

	return id, nil
}

func (s *Storage) Close() {
	if s.db != nil {
		s.once.Do(func() { s.db.Close() })
	}
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
