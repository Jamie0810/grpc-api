//+build wireinject

//The build tag makes sure the stub is not built in the final build.
package app

import (
	"github.com/google/wire"
	"gitlab.silkrode.com.tw/golang/kbc_pkg/database"
	log "gitlab.silkrode.com.tw/golang/kbc_pkg/logger"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/config"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/delivery/grpc"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/repository/mysql"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/usecase"
	"gorm.io/gorm"
)

func Initialize(configPath string) (*Application, error) {
	wire.Build(
		newApplication,
		config.NewConfig,
		grpc.NewGRPCClient,
		mysql.NewWithdrawalRepository,
		NewLogger,
		usecase.NewWithdrawalUsecase,
		NewDatabase,
	)
	return &Application{}, nil
}

func NewLogger(config config.Config) (*log.Logger, error) {
	return log.NewLogger(config.LogLevel, config.LogFormat)
}

func NewDatabase(config config.Config, logger *log.Logger) (*gorm.DB, error) {
	return database.NewDB(
		config.Database.Driver,
		config.Database.Host,
		config.Database.Port,
		config.Database.DBName,
		config.Database.InstanceName,
		config.Database.User,
		config.Database.Password,
		true,
		"10s",
		"10s",
		"10s",
	)
}
