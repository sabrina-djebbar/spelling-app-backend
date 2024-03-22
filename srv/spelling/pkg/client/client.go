package client

import "net/http"

const ListPath = "/spelling/list"

type Client struct {
	internal http.Handler
}
