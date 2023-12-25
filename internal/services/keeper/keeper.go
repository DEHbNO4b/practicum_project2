package keeper

import (
	"context"
	"log/slog"

	"github.com/DEHbNO4b/practicum_project2/internal/domain/models"
)

type LogPassStorage interface {
	LogPass(ctx context.Context, id int64) ([]models.LogPassData, error)
	SetLogPass(ctx context.Context, lp models.LogPassData) error
	Close()
}

type Keeper struct {
	log       *slog.Logger
	lpStorage LogPassStorage
}

// New returns a new intance of Keeper
func New(
	log *slog.Logger,
	lp LogPassStorage,

) *Keeper {
	return &Keeper{
		log:       log,
		lpStorage: lp,
	}
}
