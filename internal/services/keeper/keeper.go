package keeper

import (
	"context"
	"log/slog"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
)

// type LogPassStorage interface {
// 	LogPass(ctx context.Context, id int64) ([]models.LogPassData, error)
// 	SetLogPass(ctx context.Context, lp models.LogPassData) error
// 	Close()
// }

type LogPassStorage interface {
	SaveLogPass(ctx context.Context, lp models.LogPassData) error
	LogPass(ctx context.Context, id int64) ([]models.LogPassData, error)
}

type TextStorage interface {
	SaveText(ctx context.Context, lp models.TextData) error
	TextData(ctx context.Context, id int64) ([]models.TextData, error)
}

type BinaryStorage interface {
	SaveBinary(ctx context.Context, lp models.BinaryData) error
	BinaryData(ctx context.Context, id int64) ([]models.BinaryData, error)
}

type CardStorage interface {
	SaveCard(ctx context.Context, lp models.Card) error
	CardData(ctx context.Context, id int64) ([]models.Card, error)
}

type Keeper struct {
	log     *slog.Logger
	logPass LogPassStorage
	text    TextStorage
	binary  BinaryStorage
	card    CardStorage
}

// New returns a new intance of Keeper
func New(
	log *slog.Logger,
	logPass LogPassStorage,
	text TextStorage,
	binary BinaryStorage,
	card CardStorage,

) *Keeper {
	return &Keeper{
		log:     log,
		logPass: logPass,
		text:    text,
		binary:  binary,
		card:    card,
	}
}
func (k *Keeper) SaveLogPass(ctx context.Context, lp models.LogPassData) error {

	err := k.logPass.SaveLogPass(ctx, lp)
	if err != nil {
		return err
	}

	return nil
}
func (k *Keeper) LogPass(ctx context.Context, id int64) ([]models.LogPassData, error) {

	lpd, err := k.logPass.LogPass(ctx, id)
	if err != nil {
		return nil, err
	}

	return lpd, nil
}

func (k *Keeper) SaveText(ctx context.Context, td models.TextData) error {

	err := k.text.SaveText(ctx, td)
	if err != nil {
		return err
	}

	return nil
}
func (k *Keeper) TextData(ctx context.Context, id int64) ([]models.TextData, error) {

	td, err := k.text.TextData(ctx, id)
	if err != nil {
		return nil, err
	}

	return td, err
}

func (k *Keeper) SaveBinary(ctx context.Context, bd models.BinaryData) error {

	err := k.binary.SaveBinary(ctx, bd)
	if err != nil {
		return err
	}

	return nil
}
func (k *Keeper) BinaryData(ctx context.Context, id int64) ([]models.BinaryData, error) {

	bd, err := k.binary.BinaryData(ctx, id)
	if err != nil {
		return nil, err
	}

	return bd, err
}

func (k *Keeper) SaveCard(ctx context.Context, c models.Card) error {

	err := k.card.SaveCard(ctx, c)
	if err != nil {
		return err
	}

	return nil
}
func (k *Keeper) CardData(ctx context.Context, id int64) ([]models.Card, error) {

	cd, err := k.card.CardData(ctx, id)
	if err != nil {
		return nil, err
	}

	return cd, nil
}
