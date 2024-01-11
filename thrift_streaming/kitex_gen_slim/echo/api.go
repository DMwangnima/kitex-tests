// Code generated by thriftgo (0.3.5). DO NOT EDIT.

package echo

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/streaming"
)

type EchoRequest struct {
	Message string `thrift:"message,1,required" frugal:"1,required,string" json:"message"`
}

func NewEchoRequest() *EchoRequest {
	return &EchoRequest{}
}

func (p *EchoRequest) InitDefault() {
	*p = EchoRequest{}
}

func (p *EchoRequest) GetMessage() (v string) {
	return p.Message
}
func (p *EchoRequest) SetMessage(val string) {
	p.Message = val
}

func (p *EchoRequest) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoRequest(%+v)", *p)
}

type EchoResponse struct {
	Message string `thrift:"message,1,required" frugal:"1,required,string" json:"message"`
}

func NewEchoResponse() *EchoResponse {
	return &EchoResponse{}
}

func (p *EchoResponse) InitDefault() {
	*p = EchoResponse{}
}

func (p *EchoResponse) GetMessage() (v string) {
	return p.Message
}
func (p *EchoResponse) SetMessage(val string) {
	p.Message = val
}

func (p *EchoResponse) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoResponse(%+v)", *p)
}

type EchoException struct {
	Message string `thrift:"message,1" frugal:"1,default,string" json:"message"`
}

func NewEchoException() *EchoException {
	return &EchoException{}
}

func (p *EchoException) InitDefault() {
	*p = EchoException{}
}

func (p *EchoException) GetMessage() (v string) {
	return p.Message
}
func (p *EchoException) SetMessage(val string) {
	p.Message = val
}

func (p *EchoException) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoException(%+v)", *p)
}
func (p *EchoException) Error() string {
	return p.String()
}

type EchoService interface {
	EchoBidirectional(stream EchoService_EchoBidirectionalServer) (err error)

	EchoClient(stream EchoService_EchoClientServer) (err error)

	EchoServer(req *EchoRequest, stream EchoService_EchoServerServer) (err error)

	EchoUnary(ctx context.Context, req1 *EchoRequest) (r *EchoResponse, err error)

	EchoPingPong(ctx context.Context, req1 *EchoRequest, req2 *EchoRequest) (r *EchoResponse, err error)

	EchoOneway(ctx context.Context, req1 *EchoRequest) (err error)

	Ping(ctx context.Context) (err error)
}

type EchoService_EchoBidirectionalServer interface {
	streaming.Stream

	Recv() (*EchoRequest, error)

	Send(*EchoResponse) error
}
type EchoService_EchoClientServer interface {
	streaming.Stream

	Recv() (*EchoRequest, error)

	SendAndClose(*EchoResponse) error
}
type EchoService_EchoServerServer interface {
	streaming.Stream

	Send(*EchoResponse) error
}

// exceptions of methods in EchoService.
var (
	_ error = (*EchoException)(nil)
)
