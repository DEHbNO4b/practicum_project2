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

// UserSaver interface implementation

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
	return id, nil
}

func (s *Storage) Close() {
	if s.db != nil {
		s.once.Do(func() { s.db.Close() })
	}
}

// UserProvider interface implementation

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

// LogPassStorage interface implementation
func (s *Storage) SetLogPass(ctx context.Context, lp models.LogPassData) error {
	op := "storage.postgres.SetLogPass"

	local := domainLpToLocal(lp)

	_, err := s.db.ExecContext(ctx,
		`INSERT INTO logpassdata(user_id,login,pass,meta) VALUES($1,$2,$3,$4)`,
		local.UserID, local.Login, local.Pass, local.Meta,
	)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}
	return nil

}

func (s *Storage) LogPass(ctx context.Context, id int64) ([]models.LogPassData, error) {

	op := "storage.postgres.LogPass"

	ans := make([]models.LogPassData, 0, 10)

	rows, err := s.db.QueryContext(ctx, `SELECT login, pass, meta from logpassdata WHERE user_id=$1`, id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return ans, fmt.Errorf("%s %w", op, storage.ErrNoDataFound)
		}

		return ans, fmt.Errorf("%s %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var lp LogPassData

		if err := rows.Scan(&lp.Login, &lp.Pass, &lp.Meta); err != nil {
			return ans, fmt.Errorf("%s %w", op, err)
		}

		dlp := lpToDomain(lp)

		ans = append(ans, dlp)

	}
	return ans, nil
}
