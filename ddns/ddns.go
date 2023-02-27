package ddns

import (
	"context"

	"github.com/h3mmy/ddns/ddns/internal/config"
	"github.com/h3mmy/ddns/ddns/internal/models"
	"github.com/h3mmy/ddns/ddns/internal/providers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
)

type chId string

var (
	ipdiscovery chId = "ipdiscovery"
)

type DDNSWorker struct {
	logger          *zap.Logger
	config          *config.DDNSConfig
	ipDiscoveryTask models.DiscoveryTask
	chMap           map[chId](chan *models.IPSet)
}

func NewDDNSWorker() *DDNSWorker {
	lgr := providers.NewLogger().With(
		zap.Field{Key: "group", Type: zapcore.StringType, String: "ddns_worker"},
	)
	return &DDNSWorker{
		config:          providers.GetConfig(),
		ipDiscoveryTask: providers.IfConfigTask(),
		logger:          lgr,
		chMap:           make(map[chId]chan *models.IPSet),
	}
}

func (dw *DDNSWorker) Start(ctx context.Context) error {
	processId := ctx.Value(models.ProcessId).(string)
	logger := dw.logger.With(
		zapcore.Field{Key: "method", Type: zapcore.StringType, String: "Start"},
		zapcore.Field{Key: string(models.ProcessId), Type: zapcore.StringType, String: processId},
	)

	errGroup, ctx := errgroup.WithContext(ctx)

	logger.Debug("Starting IP Discovery")
	errGroup.Go(func() error {
		return dw.GetCurrentIPs(ctx)
	})

	return errGroup.Wait()
}

func (dw *DDNSWorker) GetCurrentIPs(ctx context.Context) error {
	logger := dw.logger.With(
		zapcore.Field{Key: "method", Type: zapcore.StringType, String: "GetCurrentIPs"},
	)
	errGroup, ctx := errgroup.WithContext(ctx)

	errGroup.Go(func() error {
		return dw.ipDiscoveryTask.DiscoverIPs(ctx)
	})

	err := errGroup.Wait()
	if err != nil {
		logger.Fatal("error while discovering IPs", zap.Error(err))
		return err
	}
	dw.chMap[ipdiscovery] <- dw.ipDiscoveryTask.GetResultSet()
	close(dw.chMap[ipdiscovery])
	return nil
}
