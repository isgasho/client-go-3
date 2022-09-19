package client_go

import (
	"context"
	"github.com/pkg/errors"
	"github.com/wuqinqiang/easycar/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

var (
	EmptyGroup = errors.New("The Group cannot be empty")
)

type Client struct {
	easyCarServerUri string
	level            Level
	groups           []Group

	grpcConn   *grpc.ClientConn
	easycarCli proto.EasyCarClient

	options *options
}

func NewClient(uri string, fns ...OptFn) *Client {
	opt := &options{
		connTimeOut: 5 * time.Second}

	for _, fn := range fns {
		fn(opt)
	}
	return &Client{
		easyCarServerUri: uri,
		level:            1,
		options:          opt,
	}
}

func (client *Client) AddGroup(skip bool, groups ...Group) *Client {
	if skip {
		client.level++
	}
	for _, group := range groups {
		group.SetLevel(client.level)
	}
	return client
}

func (client *Client) Begin(ctx context.Context) (gid string, err error) {
	cliConn, err := client.getConn(ctx)
	if err != nil {
		return "", err
	}
	resp, err := cliConn.Begin(ctx, &emptypb.Empty{})
	if err != nil {
		return "", err
	}
	return resp.GetGId(), nil
}

func (client *Client) Register(ctx context.Context) error {
	if len(client.groups) == 0 {
		return EmptyGroup
	}
	cliConn, err := client.getConn(ctx)
	if err != nil {
		return err
	}
	var (
		req proto.RegisterReq
	)
	for _, group := range client.groups {
		for _, branch := range group.branches {
			b := branch.Convert()
			b.TranType = ConvertTranTypeToGrpc(group.GetTranType())
			req.Branches = append(req.Branches, b)
		}
	}
	_, err = cliConn.Register(ctx, &req)
	return err
}

func (client *Client) Start(ctx context.Context, gId string) error {
	cliConn, err := client.getConn(ctx)
	if err != nil {
		return err
	}
	var (
		req proto.StartReq
	)
	req.GId = gId
	_, err = cliConn.Start(ctx, &req)
	return err
}

func (client *Client) Commit(ctx context.Context, gId string) error {
	cliConn, err := client.getConn(ctx)
	if err != nil {
		return err
	}
	var (
		req proto.CommitReq
	)
	req.GId = gId
	_, err = cliConn.Commit(ctx, &req)
	return err
}

func (client *Client) Rollback(ctx context.Context, gId string) error {
	cliConn, err := client.getConn(ctx)
	if err != nil {
		return err
	}
	var (
		req proto.RollBckReq
	)
	req.GId = gId
	_, err = cliConn.Rollback(ctx, &req)
	return err
}

func (client *Client) getConn(ctx context.Context) (cli proto.EasyCarClient, err error) {
	if client.easycarCli != nil {
		return client.easycarCli, nil
	}

	forkCtx, cancel := context.WithTimeout(ctx, client.options.connTimeOut)
	defer cancel()
	conn, err := grpc.DialContext(forkCtx, client.easyCarServerUri,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	cli = proto.NewEasyCarClient(conn)
	client.easycarCli = cli
	return
}

func (client *Client) Close(ctx context.Context) error {
	return client.grpcConn.Close()
}
