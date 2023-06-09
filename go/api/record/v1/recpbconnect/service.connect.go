// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: api/record/v1/service.proto

package recpbconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/minauteur/go_meetup_api/go/api/record/v1"
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
	// RecordAPIName is the fully-qualified name of the RecordAPI service.
	RecordAPIName = "api.record.v1.RecordAPI"
)

// RecordAPIClient is a client for the api.record.v1.RecordAPI service.
type RecordAPIClient interface {
	Get(context.Context, *connect_go.Request[v1.GetRecordRequest]) (*connect_go.Response[v1.GetRecordResponse], error)
}

// NewRecordAPIClient constructs a client for the api.record.v1.RecordAPI service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewRecordAPIClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) RecordAPIClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &recordAPIClient{
		get: connect_go.NewClient[v1.GetRecordRequest, v1.GetRecordResponse](
			httpClient,
			baseURL+"/api.record.v1.RecordAPI/Get",
			opts...,
		),
	}
}

// recordAPIClient implements RecordAPIClient.
type recordAPIClient struct {
	get *connect_go.Client[v1.GetRecordRequest, v1.GetRecordResponse]
}

// Get calls api.record.v1.RecordAPI.Get.
func (c *recordAPIClient) Get(ctx context.Context, req *connect_go.Request[v1.GetRecordRequest]) (*connect_go.Response[v1.GetRecordResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// RecordAPIHandler is an implementation of the api.record.v1.RecordAPI service.
type RecordAPIHandler interface {
	Get(context.Context, *connect_go.Request[v1.GetRecordRequest]) (*connect_go.Response[v1.GetRecordResponse], error)
}

// NewRecordAPIHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewRecordAPIHandler(svc RecordAPIHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/api.record.v1.RecordAPI/Get", connect_go.NewUnaryHandler(
		"/api.record.v1.RecordAPI/Get",
		svc.Get,
		opts...,
	))
	return "/api.record.v1.RecordAPI/", mux
}

// UnimplementedRecordAPIHandler returns CodeUnimplemented from all methods.
type UnimplementedRecordAPIHandler struct{}

func (UnimplementedRecordAPIHandler) Get(context.Context, *connect_go.Request[v1.GetRecordRequest]) (*connect_go.Response[v1.GetRecordResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("api.record.v1.RecordAPI.Get is not implemented"))
}
