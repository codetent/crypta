// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: control/v1/control.proto

package controlv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/codetent/crypta/gen/control/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// ControlServiceName is the fully-qualified name of the ControlService service.
	ControlServiceName = "control.v1.ControlService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ControlServiceShutdownProcedure is the fully-qualified name of the ControlService's Shutdown RPC.
	ControlServiceShutdownProcedure = "/control.v1.ControlService/Shutdown"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	controlServiceServiceDescriptor        = v1.File_control_v1_control_proto.Services().ByName("ControlService")
	controlServiceShutdownMethodDescriptor = controlServiceServiceDescriptor.Methods().ByName("Shutdown")
)

// ControlServiceClient is a client for the control.v1.ControlService service.
type ControlServiceClient interface {
	Shutdown(context.Context, *connect.Request[v1.ShutdownRequest]) (*connect.Response[v1.ShutdownResponse], error)
}

// NewControlServiceClient constructs a client for the control.v1.ControlService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewControlServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) ControlServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &controlServiceClient{
		shutdown: connect.NewClient[v1.ShutdownRequest, v1.ShutdownResponse](
			httpClient,
			baseURL+ControlServiceShutdownProcedure,
			connect.WithSchema(controlServiceShutdownMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// controlServiceClient implements ControlServiceClient.
type controlServiceClient struct {
	shutdown *connect.Client[v1.ShutdownRequest, v1.ShutdownResponse]
}

// Shutdown calls control.v1.ControlService.Shutdown.
func (c *controlServiceClient) Shutdown(ctx context.Context, req *connect.Request[v1.ShutdownRequest]) (*connect.Response[v1.ShutdownResponse], error) {
	return c.shutdown.CallUnary(ctx, req)
}

// ControlServiceHandler is an implementation of the control.v1.ControlService service.
type ControlServiceHandler interface {
	Shutdown(context.Context, *connect.Request[v1.ShutdownRequest]) (*connect.Response[v1.ShutdownResponse], error)
}

// NewControlServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewControlServiceHandler(svc ControlServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	controlServiceShutdownHandler := connect.NewUnaryHandler(
		ControlServiceShutdownProcedure,
		svc.Shutdown,
		connect.WithSchema(controlServiceShutdownMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/control.v1.ControlService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ControlServiceShutdownProcedure:
			controlServiceShutdownHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedControlServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedControlServiceHandler struct{}

func (UnimplementedControlServiceHandler) Shutdown(context.Context, *connect.Request[v1.ShutdownRequest]) (*connect.Response[v1.ShutdownResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("control.v1.ControlService.Shutdown is not implemented"))
}