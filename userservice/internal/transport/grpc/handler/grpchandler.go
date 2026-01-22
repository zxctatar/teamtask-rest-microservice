package grpchandler

import (
	"context"
	"log/slog"
	userservicev1 "userservice/proto/userservice"
)

type GRPCHandler struct {
	log *slog.Logger
	userservicev1.UnimplementedUserServiceServer
}

func NewGRPCHandler(log *slog.Logger) *GRPCHandler {
	return &GRPCHandler{
		log: log,
	}
}

func (g *GRPCHandler) GetIdBySession(ctx context.Context, req *userservicev1.GetIdBySessionRequest) (*userservicev1.GetIdBySessionResponse, error) {
	panic("not implemented")
}
