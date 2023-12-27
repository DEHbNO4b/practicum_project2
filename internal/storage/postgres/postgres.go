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
func (s *Storage) SaveLogPass(ctx context.Context, lp models.LogPassData) error {

	op := "storage.postgres.SetLogPass"

	local := domainLpToLocal(lp)

	_, err := s.db.ExecContext(ctx,
		`INSERT INTO logpass_data(user_id,login,pass,meta) VALUES($1,$2,$3,$4)`,
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

	rows, err := s.db.QueryContext(ctx, `SELECT login, pass, meta from logpass_data WHERE user_id=$1`, id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return ans, fmt.Errorf("%s %w", op, storage.ErrNoDataFound)
		}

		return ans, fmt.Errorf("%s %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var local LogPassData
		local.UserID = id

		if err := rows.Scan(&local.Login, &local.Pass, &local.Meta); err != nil {
			return ans, fmt.Errorf("%s %w", op, err)
		}

		dlp := lpToDomain(local)

		ans = append(ans, dlp)

	}
	return ans, nil
}

// TextKeeper interface implementation

func (s *Storage) SaveText(ctx context.Context, td models.TextData) error {

	op := "storage.postgres.SaveText"

	local := domainTextToLocal(td)

	_, err := s.db.ExecContext(ctx,
		`INSERT INTO text_data(user_id,text,meta) VALUES($1,$2,$3)`,
		local.UserID, local.Text, local.Meta,
	)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}
func (s *Storage) TextData(ctx context.Context, id int64) ([]models.TextData, error) {

	op := "storage.postgres.TextData"

	ans := make([]models.TextData, 0, 10)

	rows, err := s.db.QueryContext(ctx, `SELECT text, meta from text_data WHERE user_id=$1`, id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return ans, fmt.Errorf("%s %w", op, storage.ErrNoDataFound)
		}

		return ans, fmt.Errorf("%s %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var local TextData
		local.UserID = id

		if err := rows.Scan(&local.Text, &local.Meta); err != nil {
			return ans, fmt.Errorf("%s %w", op, err)
		}

		td := textToDomain(local)

		ans = append(ans, td)

	}
	return ans, nil
}

// BinaryKeeper interface implementation
func (s *Storage) SaveBinary(ctx context.Context, bd models.BinaryData) error {

	op := "storage.postgres.SaveBinary"

	local := domainBinaryToLocal(bd)

	_, err := s.db.ExecContext(ctx,
		`INSERT INTO binary_data(user_id,data,meta) VALUES($1,$2,$3)`,
		local.UserID, local.Data, local.Meta,
	)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}
	return nil
}
func (s *Storage) BinaryData(ctx context.Context, id int64) ([]models.BinaryData, error) {
	op := "storage.postgres.BinaryData"

	ans := make([]models.BinaryData, 0, 10)

	rows, err := s.db.QueryContext(ctx, `SELECT data, meta from binary_data WHERE user_id=$1`, id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return ans, fmt.Errorf("%s %w", op, storage.ErrNoDataFound)
		}

		return ans, fmt.Errorf("%s %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {
		var local BinaryData
		local.UserID = id

		if err := rows.Scan(&local.Data, &local.Meta); err != nil {
			return ans, fmt.Errorf("%s %w", op, err)
		}

		bd := binaryToDomain(local)

		ans = append(ans, bd)

	}
	return ans, nil
}

// CardKeeper interface implementation
func (s *Storage) SaveCard(ctx context.Context, cd models.Card) error {

	op := "storage.postgres.SaveCard"

	local := domainCardToLocal(cd)

	_, err := s.db.ExecContext(ctx,
		`INSERT INTO card_data(user_id,card_id,pass,date,meta) VALUES($1,$2,$3,$4,$5)`,
		local.UserID, local.CardID, local.Pass, local.Date, local.Meta,
	)
	if err != nil {
		return fmt.Errorf("%s %w", op, err)
	}

	return nil
}
func (s *Storage) CardData(ctx context.Context, id int64) ([]models.Card, error) {

	op := "storage.postgres.CardData"

	ans := make([]models.Card, 0, 10)

	rows, err := s.db.QueryContext(ctx, `SELECT card_id, pass,date,meta from card_data WHERE user_id=$1`, id)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return ans, fmt.Errorf("%s %w", op, storage.ErrNoDataFound)
		}

		return ans, fmt.Errorf("%s %w", op, err)
	}

	defer rows.Close()

	for rows.Next() {

		var local Card
		local.UserID = id

		if err := rows.Scan(&local.CardID, &local.Pass, &local.Date, &local.Meta); err != nil {
			return ans, fmt.Errorf("%s %w", op, err)
		}

		cd, err := cardToDomain(local)
		if err != nil {
			return ans, err
		}

		ans = append(ans, cd)

	}
	return ans, nil
}
