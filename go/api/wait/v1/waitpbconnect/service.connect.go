// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/wait/v1/service.proto

package waitpbconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/minauteur/go_meetup_api/go/api/wait/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// WaitAPIName is the fully-qualified name of the WaitAPI service.
	WaitAPIName = "api.wait.v1.WaitAPI"
)

// WaitAPIClient is a client for the api.wait.v1.WaitAPI service.
type WaitAPIClient interface {
	Wait(context.Context, *connect_go.Request[v1.WaitRequest]) (*connect_go.Response[v1.WaitResponse], error)
}

// NewWaitAPIClient constructs a client for the api.wait.v1.WaitAPI service. By default, it uses the
// Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewWaitAPIClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) WaitAPIClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &waitAPIClient{
		wait: connect_go.NewClient[v1.WaitRequest, v1.WaitResponse](
			httpClient,
			baseURL+"/api.wait.v1.WaitAPI/Wait",
			opts...,
		),
	}
}

// waitAPIClient implements WaitAPIClient.
type waitAPIClient struct {
	wait *connect_go.Client[v1.WaitRequest, v1.WaitResponse]
}

// Wait calls api.wait.v1.WaitAPI.Wait.
func (c *waitAPIClient) Wait(ctx context.Context, req *connect_go.Request[v1.WaitRequest]) (*connect_go.Response[v1.WaitResponse], error) {
	return c.wait.CallUnary(ctx, req)
}

// WaitAPIHandler is an implementation of the api.wait.v1.WaitAPI service.
type WaitAPIHandler interface {
	Wait(context.Context, *connect_go.Request[v1.WaitRequest]) (*connect_go.Response[v1.WaitResponse], error)
}

// NewWaitAPIHandler builds an HTTP handler from the service implementation. It returns the path on
// which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewWaitAPIHandler(svc WaitAPIHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/api.wait.v1.WaitAPI/Wait", connect_go.NewUnaryHandler(
		"/api.wait.v1.WaitAPI/Wait",
		svc.Wait,
		opts...,
	))
	return "/api.wait.v1.WaitAPI/", mux
}

// UnimplementedWaitAPIHandler returns CodeUnimplemented from all methods.
type UnimplementedWaitAPIHandler struct{}

func (UnimplementedWaitAPIHandler) Wait(context.Context, *connect_go.Request[v1.WaitRequest]) (*connect_go.Response[v1.WaitResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.wait.v1.WaitAPI.Wait is not implemented"))
}
