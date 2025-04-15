package account

import (
	"github.com/loctodale/go_api_hubs_microservice/account/global"
	"github.com/loctodale/go_api_hubs_microservice/account/pb"
	"google.golang.org/grpc"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.AccountServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewAccountServiceClient(conn)
	return &Client{conn, c}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) PostAccount(userAccount string, userPassword string) error {
	_, err := c.service.PostAccount(
		global.Ctx,
		&pb.PostAccountRequest{UserAccount: userAccount, UserPassword: userPassword},
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetAccounts() (pb.GetAccountsResponse, error) {
	result, err := c.service.GetAccounts(global.Ctx, &pb.Empty{})
	if err != nil {
		return pb.GetAccountsResponse{}, err
	}

	return pb.GetAccountsResponse{
		Account: result.Account,
	}, nil
}
