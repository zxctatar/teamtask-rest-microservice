package userserviceclient

import (
	"context"
	"fmt"
	"log/slog"
	userservicev1 "projectservice/proto/userservice"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserServiceClient struct {
	log    *slog.Logger
	conn   *grpc.ClientConn
	client userservicev1.UserServiceClient
}

func NewUserServiceClient(log *slog.Logger, host string, port uint32) *UserServiceClient {
	const op = "userserviceclient.NewUserServiceClient"

	log.Info("create grpc client", slog.String("op", op))
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("cannot create new grpc client: " + err.Error())
	}

	client := userservicev1.NewUserServiceClient(conn)

	return &UserServiceClient{
		log:    log,
		conn:   conn,
		client: client,
	}
}

func (u *UserServiceClient) GetIdBySession(ctx context.Context, sessionId string) (uint32, error) {
	in := &userservicev1.GetIdBySessionRequest{
		SessionId: sessionId,
	}

	res, err := u.client.GetIdBySession(ctx, in)
	if err != nil {
		return 0, err
	}

	return res.UserId, nil
}

func (u *UserServiceClient) Stop() {
	u.conn.Close()
}
