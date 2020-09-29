package grpc

import (
	"net"
	"time"

	pb "gitlab.silkrode.com.tw/golang/kbc2/proto/order"
	"gitlab.silkrode.com.tw/golang/kbc_pkg/logger"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/config"
	"gitlab.silkrode.com.tw/team_golang/kbc2/captain-marvel/internal/pkg/domain"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type GRPCClient struct {
	cfg               config.Config
	log               *logger.Logger
	withdrawalUsecase domain.WithdrawalUsecase
}

func NewGRPCClient(
	cfg config.Config,
	wu domain.WithdrawalUsecase,
	log *logger.Logger,
) *GRPCClient {
	return &GRPCClient{
		cfg:               cfg,
		withdrawalUsecase: wu,
		log:               log,
	}
}

func (c *GRPCClient) Start() (err error) {
	listener, err := net.Listen("tcp", ":"+c.cfg.Server.GrpcPort)
	if err != nil {
		return err
	}

	s := grpc.NewServer(
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:              (time.Duration(5) * time.Second), // Ping the client if it is idle for 5 seconds to ensure the connection is still active
				Timeout:           (time.Duration(5) * time.Second), // Wait 5 second for the ping ack before assuming the connection is dead
				MaxConnectionIdle: 5 * time.Minute,
			},
		),
		grpc.KeepaliveEnforcementPolicy(
			keepalive.EnforcementPolicy{
				MinTime:             (time.Duration(2) * time.Second), // If a client pings more than once every 2 seconds, terminate the connection
				PermitWithoutStream: true,                             // Allow pings even when there are no active streams
			},
		),
	)

	pb.RegisterOrderServer(s, c.withdrawalUsecase)
	reflection.Register(s)

	c.log.Info().Msgf("start grpc server port:%s", c.cfg.Server.GrpcPort)
	return s.Serve(listener)
}
