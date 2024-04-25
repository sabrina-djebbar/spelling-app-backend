package client

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sabrina-djebbar/spelling-app-backend/lib/shttp/middleware"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var (
	internalRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "internal_request_totals",
			Help: "how many internal requests were sent",
		},
		[]string{"code", "endpoint", "host"},
	)
	internalDurations = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "internal_latency_totals",
			Help: "HTTP internal request latency duration",
		},
		[]string{"code", "endpoint", "host"},
	)
)

func init() {
	prometheus.MustRegister(internalRequests)
	prometheus.MustRegister(internalDurations)
}

type InternalClient struct {
	host    string
	token   string
	timeout int
	http    *http.Client
}

type InternalClientOptions struct {
	Host    string `json:"host"`
	Token   string `json:"token"`
	Timeout int    `json:"timeout"`
}

type InternalReqMaker interface {
	Do(ctx context.Context, route string, in interface{}, out interface{}) error
}

func NewInternalClient(options InternalClientOptions) *InternalClient {
	http := &http.Client{
		Timeout: time.Duration(options.Timeout) * time.Second,
	}
	return &InternalClient{
		host:    options.Host,
		token:   options.Token,
		timeout: options.Timeout,
		http:    http,
	}
}

func (client *InternalClient) setReferer(_ context.Context, req *http.Request) {
	if middleware.ServedBy != "" {
		req.Header.Set("Referer", middleware.ServedBy)
	}
}

func (client *InternalClient) setRequestID(ctx context.Context, req *http.Request) {
	value := ctx.Value(middleware.HTTPRequestCtxIdentifier)
	if requestID, ok := value.(string); ok {
		req.Header.Set(middleware.HTTPRequestHeaderIdentifier, requestID)
	}
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
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client.setRequestID(ctx, req)
	client.setReferer(ctx, req)

	return req, nil
}

func (client *InternalClient) Do(ctx context.Context, route string, in interface{}, out interface{}) error {
	return client.do(ctx, route, in, out)
}

func (client *InternalClient) do(ctx context.Context, route string, in interface{}, out interface{}) error {
	start := time.Now()

	data, err := json.Marshal(in)
	if err != nil {
		return err
	}

	req, err := client.getRequest(ctx, route, data)
	if err != nil {
		return err
	}

	resp, err := client.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	status := strconv.Itoa(resp.StatusCode)
	elapsed := float64(time.Since(start)) / float64(time.Second)

	internalRequests.WithLabelValues(status, req.URL.Path, req.Host).Inc()
	internalDurations.WithLabelValues(status, req.URL.Path, req.Host).Observe(elapsed)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// If the request is failed then we should try to parse our error response back
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusNoContent {

		err := json.Unmarshal(body, out)
		if err != nil {
			return err
		}

	}

	if len(body) == 0 {
		return nil
	} else {
		// Marshal our request back out
		err = json.Unmarshal(body, out)
	}

	if err != nil {
		return err
	}

	return nil
}
