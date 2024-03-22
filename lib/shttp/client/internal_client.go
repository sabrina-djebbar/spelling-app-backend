package client

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	InternalPreconditionFailed = "internal_precondition_failed"
	InternalRequestFailedCode  = "internal_request_failed"
)

// InternalClient is a client designed to add authorasation headers
type InternalClient struct {
	host    string
	token   string
	timeout time.Duration
	client  http.Client
	logger  *log.Logger
}

type InternalClientOptions struct {
	Name    string `json:"name"`
	Host    string `json:"host"`
	Token   string `json:"token"`
	Timeout int    `json:"timeout"`
}

type InternalReqMaker interface {
	Do(ctx context.Context, route string, req interface{}, resp interface{}) error
}

type noOpRequestResponseLogger struct{}

func NewInternalClient(opts InternalClientOptions) *InternalClient {
	ic := &InternalClient{
		host:    opts.Host,
		token:   opts.Token,
		timeout: time.Duration(opts.Timeout) * time.Second,
		logger:  log.New(os.Stdout, "spelling-api", log.LstdFlags),
	}

	return ic
}

func (client *InternalClient) getPath(route string) (string, error) {
	u, err := url.Parse(client.host)
	if err != nil {
		return "", err
	}

	u.Path = route

	return u.String(), nil
}
func (client *InternalClient) getRequest(ctx context.Context, route string, data []byte) (*http.Request, error) {
	path, err := client.getPath(route)
	if err != nil {
		return http.Request{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewBuffer(data))
	if err != nil {
		return http.Request{}, err
	}

	req.Header.Set("Content-Type", "application/json")

	// TODO: IMPLEMENT setAuthHeader FUNCTION
	return req, nil
}

func (client *InternalClient) do(ctx context.Context, route string, in interface{}, out interface{}) error {
	data, err := json.Marshal(in)
	req, err := client.getRequest(ctx, route, data)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	/* if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted && resp.StatusNoContent {
		errResponse := http.Error{}
	} */

	if len(body) == 0 {
		return nil
	}else{
		err = json.Unmarshal(body, out)
	}
	if err != nil return {err}
	return nil
}
