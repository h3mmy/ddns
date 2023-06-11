package services

import (
	"context"
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/h3mmy/ddns/ddns/internal/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type CloudflareService struct {
	logger *zap.Logger
	api    *cloudflare.API
}

func NewCloudflareService(api *cloudflare.API) *CloudflareService {
	lgr := log.NewLoggerForService("cloudflare_service")
	return &CloudflareService{
		logger: lgr,
		api:    api,
	}
}

// Gets a single A record using domain name
func (cfs *CloudflareService) GetDNSARecord(ctx context.Context, domain string) (*cloudflare.DNSRecord, error) {
	logger := cfs.logger.With(
		zapcore.Field{Key: "method", Type: zapcore.StringType, String: "GetARecord"},
		zapcore.Field{Key: "domain", Type: zapcore.StringType, String: domain},
	)

	var rc *cloudflare.ResourceContainer
	zone, err := cfs.api.ZoneIDByName(domain)
	if err != nil {
		logger.Warn("could not get zoneID from zone name. Will try using Account Identifier instead", zap.Error(err))
		rc = cloudflare.AccountIdentifier(cfs.api.AccountID)
	} else {
		rc = cloudflare.ZoneIdentifier(zone)
	}

	params := cloudflare.ListDNSRecordsParams{
		Type: "A",
		Name: domain,
	}

	dnsRecords, resultInfo, err := cfs.api.ListDNSRecords(
		ctx,
		rc,
		params)
	if err != nil {
		logger.Error("error getting DNSRecord", zap.Error(err))
		return nil, err
	}
	if resultInfo.Count > 1 {
		logger.Warn(fmt.Sprintf("multiple records (%d) found for domain? Will use first one", resultInfo.Count))
	}
	if resultInfo.Count == 0 {
		logger.Info(fmt.Sprintf("No results for %s", domain))
		return nil, nil
	}
	return &dnsRecords[0], nil
}

// Gets single AAAA Record using domain name
func (cfs *CloudflareService) GetDNSAAAARecord(ctx context.Context, domain string) (*cloudflare.DNSRecord, error) {
	logger := cfs.logger.With(
		zapcore.Field{Key: "method", Type: zapcore.StringType, String: "GetAAAARecord"},
		zapcore.Field{Key: "domain", Type: zapcore.StringType, String: domain},
	)

	var rc *cloudflare.ResourceContainer
	zone, err := cfs.api.ZoneIDByName(domain)
	if err != nil {
		logger.Warn("could not get zoneID from zone name. Will try using Account Identifier instead", zap.Error(err))
		rc = cloudflare.AccountIdentifier(cfs.api.AccountID)
	} else {
		rc = cloudflare.ZoneIdentifier(zone)
	}

	params := cloudflare.ListDNSRecordsParams{
		Type: "AAAA",
		Name: domain,
	}

	dnsRecords, resultInfo, err := cfs.api.ListDNSRecords(
		ctx,
		rc,
		params)
	if err != nil {
		logger.Error("error getting DNSRecord", zap.Error(err))
		return nil, err
	}
	if resultInfo.Count > 1 {
		logger.Warn(fmt.Sprintf("multiple records (%d) found for domain? Will use first one", resultInfo.Count))
	}
	if resultInfo.Count == 0 {
		logger.Info(fmt.Sprintf("No results for %s", domain))
		return nil, nil
	}
	return &dnsRecords[0], nil
}
