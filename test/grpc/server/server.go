package grpcserver

import (
	"context"
	"fmt"
	"net"
	"test/configs"
	"test/usecases/v1"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer interface {
	Start() error
	Stop()
}

type grpcServer struct {
	Usecase usecases.Usecase
	Logs    zerolog.Logger
	Server  *grpc.Server
}

func InitializeGrpcServer(u usecases.Usecase, l zerolog.Logger) GrpcServer {
	s := grpc.NewServer()
	// proto.RegisterYourServiceServer(s, &yourServer{Usecase: u})
	reflection.Register(s)
	
	return &grpcServer{
		Usecase: u,
		Logs:    l,
		Server:  s,
	}
}

func (s *grpcServer) Start() error {
	port := configs.Cfg.Server.Port // Or separate GRPC port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		s.Logs.Error().Err(err).Msg("Failed to listen for gRPC")
		return err
	}

	s.Logs.Info().Msgf("gRPC server listening on %s", port)
	return s.Server.Serve(lis)
}

func (s *grpcServer) Stop() {
	s.Server.GracefulStop()
}
