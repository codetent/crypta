// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: secret/v1/secret.proto

package secretv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/codetent/crypta/gen/secret/v1"
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
	// SecretServiceName is the fully-qualified name of the SecretService service.
	SecretServiceName = "secret.v1.SecretService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SecretServiceSetSecretProcedure is the fully-qualified name of the SecretService's SetSecret RPC.
	SecretServiceSetSecretProcedure = "/secret.v1.SecretService/SetSecret"
	// SecretServiceGetSecretProcedure is the fully-qualified name of the SecretService's GetSecret RPC.
	SecretServiceGetSecretProcedure = "/secret.v1.SecretService/GetSecret"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	secretServiceServiceDescriptor         = v1.File_secret_v1_secret_proto.Services().ByName("SecretService")
	secretServiceSetSecretMethodDescriptor = secretServiceServiceDescriptor.Methods().ByName("SetSecret")
	secretServiceGetSecretMethodDescriptor = secretServiceServiceDescriptor.Methods().ByName("GetSecret")
)

// SecretServiceClient is a client for the secret.v1.SecretService service.
type SecretServiceClient interface {
	SetSecret(context.Context, *connect.Request[v1.SetSecretRequest]) (*connect.Response[v1.SetSecretResponse], error)
	GetSecret(context.Context, *connect.Request[v1.GetSecretRequest]) (*connect.Response[v1.GetSecretResponse], error)
}

// NewSecretServiceClient constructs a client for the secret.v1.SecretService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSecretServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SecretServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &secretServiceClient{
		setSecret: connect.NewClient[v1.SetSecretRequest, v1.SetSecretResponse](
			httpClient,
			baseURL+SecretServiceSetSecretProcedure,
			connect.WithSchema(secretServiceSetSecretMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getSecret: connect.NewClient[v1.GetSecretRequest, v1.GetSecretResponse](
			httpClient,
			baseURL+SecretServiceGetSecretProcedure,
			connect.WithSchema(secretServiceGetSecretMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// secretServiceClient implements SecretServiceClient.
type secretServiceClient struct {
	setSecret *connect.Client[v1.SetSecretRequest, v1.SetSecretResponse]
	getSecret *connect.Client[v1.GetSecretRequest, v1.GetSecretResponse]
}

// SetSecret calls secret.v1.SecretService.SetSecret.
func (c *secretServiceClient) SetSecret(ctx context.Context, req *connect.Request[v1.SetSecretRequest]) (*connect.Response[v1.SetSecretResponse], error) {
	return c.setSecret.CallUnary(ctx, req)
}

// GetSecret calls secret.v1.SecretService.GetSecret.
func (c *secretServiceClient) GetSecret(ctx context.Context, req *connect.Request[v1.GetSecretRequest]) (*connect.Response[v1.GetSecretResponse], error) {
	return c.getSecret.CallUnary(ctx, req)
}

// SecretServiceHandler is an implementation of the secret.v1.SecretService service.
type SecretServiceHandler interface {
	SetSecret(context.Context, *connect.Request[v1.SetSecretRequest]) (*connect.Response[v1.SetSecretResponse], error)
	GetSecret(context.Context, *connect.Request[v1.GetSecretRequest]) (*connect.Response[v1.GetSecretResponse], error)
}

// NewSecretServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSecretServiceHandler(svc SecretServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	secretServiceSetSecretHandler := connect.NewUnaryHandler(
		SecretServiceSetSecretProcedure,
		svc.SetSecret,
		connect.WithSchema(secretServiceSetSecretMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	secretServiceGetSecretHandler := connect.NewUnaryHandler(
		SecretServiceGetSecretProcedure,
		svc.GetSecret,
		connect.WithSchema(secretServiceGetSecretMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/secret.v1.SecretService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SecretServiceSetSecretProcedure:
			secretServiceSetSecretHandler.ServeHTTP(w, r)
		case SecretServiceGetSecretProcedure:
			secretServiceGetSecretHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSecretServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSecretServiceHandler struct{}

func (UnimplementedSecretServiceHandler) SetSecret(context.Context, *connect.Request[v1.SetSecretRequest]) (*connect.Response[v1.SetSecretResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("secret.v1.SecretService.SetSecret is not implemented"))
}

func (UnimplementedSecretServiceHandler) GetSecret(context.Context, *connect.Request[v1.GetSecretRequest]) (*connect.Response[v1.GetSecretResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("secret.v1.SecretService.GetSecret is not implemented"))
}
