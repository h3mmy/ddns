package ddns

import (
	"context"
	"fmt"

	"github.com/h3mmy/ddns/ddns/internal/models"
	"github.com/h3mmy/ddns/ddns/internal/providers"
	"github.com/h3mmy/ddns/ddns/pb"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Server struct {
	pb.UnimplementedDDNSServiceServer
	logger *zap.Logger
}

// Handler for GetCurrentIP endpoint
func (s *Server) GetCurrentIP(ctx context.Context, req *pb.GetCurrentIPRequest) (*pb.GetCurrentIPResponse, error) {
	logger := s.logger.With(zapcore.Field{Key: "method", Type: zapcore.StringType, String: "GetCurrentIP"})
	defer logger.Sync()
	logger.Debug(fmt.Sprintf("Received GetCurrentIPRequest: %v", req))
	handler := GetDiscoveryHandler()
	logger.Debug("Got Discovery Handler")
	rs, err := handler.GetResultSet(ctx)
	if err != nil {
		logger.Error("Error Discovering IPs", zap.Error(err))
		return nil, err
	}
	return mapToIPResponse(rs), err
}

// Utility method for mapping from discoveryHandler output to protobuf response
func mapToIPResponse(resultSet *models.IPSet) *pb.GetCurrentIPResponse {
	logger := providers.NewLogger().With(zapcore.Field{Key: "method", Type: zapcore.StringType, String: "GetCurrentIP"})
	defer logger.Sync()
	if resultSet != nil {
		logger.Debug(fmt.Sprintf("resultSet not-nil: %v", resultSet))
		return &pb.GetCurrentIPResponse{
			IpV4: resultSet.GlobalIPv4.String(),
			IpV6: resultSet.GlobalIPv6.String(),
		}
	}
	return nil
}
