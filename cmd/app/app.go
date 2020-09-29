package app

import (
	log "gitlab.silkrode.com.tw/golang/kbc_pkg/logger"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/config"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/delivery/grpc"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/utils/crypto"
)

type Application struct {
	cfg        config.Config
	logger     *log.Logger
	grpcClient *grpc.GRPCClient
}

func newApplication(
	cfg config.Config,
	logger *log.Logger,
	grpcClient *grpc.GRPCClient,
) *Application {
	crypto.EncryptoKey = []byte(cfg.EncryptoKey)
	return &Application{
		cfg:        cfg,
		logger:     logger,
		grpcClient: grpcClient,
	}
}

func (app *Application) Launch(version string) (err error) {
	app.logger.Info().Str("version", version).Msg("grpc server")
	return app.grpcClient.Start()
}
