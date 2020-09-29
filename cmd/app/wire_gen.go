// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package app

import (
	"gitlab.silkrode.com.tw/golang/kbc_pkg/database"
	"gitlab.silkrode.com.tw/golang/kbc_pkg/logger"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/config"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/delivery/grpc"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/repository/mysql"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/usecase"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func Initialize(configPath string) (*Application, error) {
	configConfig, err := config.NewConfig(configPath)
	if err != nil {
		return nil, err
	}
	logger, err := NewLogger(configConfig)
	if err != nil {
		return nil, err
	}
	db, err := NewDatabase(configConfig, logger)
	if err != nil {
		return nil, err
	}
	withdrawalRepository := mysql.NewWithdrawalRepository(db, logger)
	withdrawalUsecase := usecase.NewWithdrawalUsecase(withdrawalRepository, logger)
	grpcClient := grpc.NewGRPCClient(configConfig, withdrawalUsecase, logger)
	application := newApplication(configConfig, logger, grpcClient)
	return application, nil
}

// wire.go:

func NewLogger(config2 config.Config) (*logger.Logger, error) {
	return logger.NewLogger(config2.LogLevel, config2.LogFormat)
}

func NewDatabase(config2 config.Config, logger2 *logger.Logger) (*gorm.DB, error) {
	return database.NewDB(config2.
		Database.Driver, config2.
		Database.Host, config2.
		Database.Port, config2.
		Database.DBName, config2.
		Database.InstanceName, config2.
		Database.User, config2.
		Database.Password, true,
		"10s",
		"10s",
		"10s",
	)
}
