package rpc

type RPC interface {
	GetUser(ctx context.Context, req client.GetUserRequest) (*client.GetUserResponse, error)
}

type rpc struct {
	app app.App
}

func New(app app.App) RPC {
	return &rpc{
		app
	}
}
