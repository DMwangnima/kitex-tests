// Code generated by Kitex v0.8.0. DO NOT EDIT.

package combineservice

import (
	"context"
	"errors"
	"fmt"
	combine "github.com/cloudwego/kitex-tests/thrift_streaming/kitex_gen_slim/combine"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

type CombineService interface {
	combine.A
	combine.B
}

var serviceMethods = map[string]kitex.MethodInfo{
	"Foo": kitex.NewMethodInfo(
		fooHandler,
		newAFooArgs,
		newAFooResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Bar": kitex.NewMethodInfo(
		barHandler,
		newBBarArgs,
		newBBarResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingBidirectional),
	),
}

var (
	combineServiceServiceInfo                = NewServiceInfo()
	combineServiceServiceInfoForClient       = NewServiceInfoForClient()
	combineServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return combineServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return combineServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return combineServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(true, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "CombineService"
	handlerType := (*CombineService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "combine",
	}
	extra["combine_service"] = true
	extra["combined_service_list"] = []string{"A", "B"}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func fooHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*combine.AFooArgs)
	realResult := result.(*combine.AFooResult)
	success, err := handler.(combine.A).Foo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newAFooArgs() interface{} {
	return combine.NewAFooArgs()
}

func newAFooResult() interface{} {
	return combine.NewAFooResult()
}

func barHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st, ok := arg.(*streaming.Args)
	if !ok {
		return errors.New("B.Bar is a thrift streaming method, please call with Kitex StreamClient")
	}
	stream := &bBarServer{st.Stream}
	return handler.(combine.B).Bar(stream)
}

type bBarClient struct {
	streaming.Stream
}

func (x *bBarClient) Send(m *combine.Req) error {
	return x.Stream.SendMsg(m)
}
func (x *bBarClient) Recv() (*combine.Rsp, error) {
	m := new(combine.Rsp)
	return m, x.Stream.RecvMsg(m)
}

type bBarServer struct {
	streaming.Stream
}

func (x *bBarServer) Send(m *combine.Rsp) error {
	return x.Stream.SendMsg(m)
}

func (x *bBarServer) Recv() (*combine.Req, error) {
	m := new(combine.Req)
	return m, x.Stream.RecvMsg(m)
}

func newBBarArgs() interface{} {
	return combine.NewBBarArgs()
}

func newBBarResult() interface{} {
	return combine.NewBBarResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Foo(ctx context.Context, req *combine.Req) (r *combine.Rsp, err error) {
	var _args combine.AFooArgs
	_args.Req = req
	var _result combine.AFooResult
	if err = p.c.Call(ctx, "Foo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Bar(ctx context.Context) (B_BarClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "Bar", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &bBarClient{res.Stream}
	return stream, nil
}
