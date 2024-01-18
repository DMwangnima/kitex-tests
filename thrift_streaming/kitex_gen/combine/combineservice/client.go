// Code generated by Kitex v0.8.0. DO NOT EDIT.

package combineservice

import (
	"context"
	combine "github.com/cloudwego/kitex-tests/thrift_streaming/kitex_gen/combine"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/client/callopt/streamcall"
	"github.com/cloudwego/kitex/client/streamclient"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	transport "github.com/cloudwego/kitex/transport"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Foo(ctx context.Context, req *combine.Req, callOptions ...callopt.Option) (r *combine.Rsp, err error)
}

// StreamClient is designed to provide Interface for Streaming APIs.
type StreamClient interface {
	Bar(ctx context.Context, callOptions ...streamcall.Option) (stream B_BarClient, err error)
}

type B_BarClient interface {
	streaming.Stream
	Send(*combine.Req) error
	Recv() (*combine.Rsp, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kCombineServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCombineServiceClient struct {
	*kClient
}

func (p *kCombineServiceClient) Foo(ctx context.Context, req *combine.Req, callOptions ...callopt.Option) (r *combine.Rsp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Foo(ctx, req)
}

// NewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
func NewStreamClient(destService string, opts ...streamclient.Option) (StreamClient, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithTransportProtocol(transport.GRPC))
	options = append(options, streamclient.GetClientOptions(opts)...)

	kc, err := client.NewClient(serviceInfoForStreamClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kCombineServiceStreamClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
// It panics if any error occurs.
func MustNewStreamClient(destService string, opts ...streamclient.Option) StreamClient {
	kc, err := NewStreamClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kCombineServiceStreamClient struct {
	*kClient
}

func (p *kCombineServiceStreamClient) Bar(ctx context.Context, callOptions ...streamcall.Option) (stream B_BarClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.Bar(ctx)
}
