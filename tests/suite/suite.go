package suite

import (
	"context"
	"strconv"
	"testing"

	"github.com/DEHbNO4b/practicum_project2/internal/config"
	pb "github.com/DEHbNO4b/practicum_project2/proto/gen/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T
	Cfg    config.ClientConfig
	Client pb.KeeperClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	// t.Parallel()

	cfg := config.MustLoadClientCfg()

	ctx, cancel := context.WithTimeout(context.Background(), cfg.FileCfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancel()
	})

	conn, err := grpc.DialContext(
		ctx,
		cfg.FileCfg.GRPC.Host+":"+strconv.Itoa(cfg.FileCfg.GRPC.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)

	}

	return ctx, &Suite{
		T:      t,
		Cfg:    cfg,
		Client: pb.NewKeeperClient(conn),
	}

}
