package ddns

import (
	"context"
	"fmt"
	"net"

	"github.com/h3mmy/ddns/ddns/internal/config"
	"github.com/h3mmy/ddns/ddns/internal/models"
	"github.com/h3mmy/ddns/ddns/internal/providers"
	"github.com/h3mmy/ddns/ddns/pb"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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


	rs, err := dw.ipDiscoveryTask.GetResultSet(ctx)
	if err != nil {
		logger.Fatal("error while discovering IPs", zap.Error(err))
		return err
	}
	dw.chMap[ipdiscovery] <- rs
	close(dw.chMap[ipdiscovery])
	return nil
}

func GetDiscoveryHandler() models.DiscoveryTask {
	return providers.IfConfigTask()
}

func GetServer() *grpc.Server {
	s := grpc.NewServer()
	pb.RegisterDDNSServiceServer(s, &Server{logger: providers.NewLogger().Named("grpc_server")})
	reflection.Register(s)
	return s
}

func StartServer(port int) error {
	logger := providers.NewLogger().With(
		zap.Field{Key: "group", Type: zapcore.StringType, String: "ddns_worker"},
	)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		logger.Fatal("failed to listen", zap.Error(err))
	}
	s := GetServer()
	logger.Info(fmt.Sprintf("Listening on port %d", port))
	return s.Serve(lis)
}

func GetConfig() *config.DDNSConfig {
	return providers.GetConfig()
}
