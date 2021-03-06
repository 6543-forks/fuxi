// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: srv/giter/proto/giter/giter.proto

package go_micro_srv_giter

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Giter service

type GiterService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Giter_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (Giter_PingPongService, error)
}

type giterService struct {
	c    client.Client
	name string
}

func NewGiterService(name string, c client.Client) GiterService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.giter"
	}
	return &giterService{
		c:    c,
		name: name,
	}
}

func (c *giterService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Giter.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giterService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (Giter_StreamService, error) {
	req := c.c.NewRequest(c.name, "Giter.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &giterServiceStream{stream}, nil
}

type Giter_StreamService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type giterServiceStream struct {
	stream client.Stream
}

func (x *giterServiceStream) Close() error {
	return x.stream.Close()
}

func (x *giterServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *giterServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *giterServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *giterService) PingPong(ctx context.Context, opts ...client.CallOption) (Giter_PingPongService, error) {
	req := c.c.NewRequest(c.name, "Giter.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &giterServicePingPong{stream}, nil
}

type Giter_PingPongService interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type giterServicePingPong struct {
	stream client.Stream
}

func (x *giterServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *giterServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *giterServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *giterServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *giterServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Giter service

type GiterHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, Giter_StreamStream) error
	PingPong(context.Context, Giter_PingPongStream) error
}

func RegisterGiterHandler(s server.Server, hdlr GiterHandler, opts ...server.HandlerOption) error {
	type giter interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type Giter struct {
		giter
	}
	h := &giterHandler{hdlr}
	return s.Handle(s.NewHandler(&Giter{h}, opts...))
}

type giterHandler struct {
	GiterHandler
}

func (h *giterHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.GiterHandler.Call(ctx, in, out)
}

func (h *giterHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.GiterHandler.Stream(ctx, m, &giterStreamStream{stream})
}

type Giter_StreamStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type giterStreamStream struct {
	stream server.Stream
}

func (x *giterStreamStream) Close() error {
	return x.stream.Close()
}

func (x *giterStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *giterStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *giterStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *giterHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.GiterHandler.PingPong(ctx, &giterPingPongStream{stream})
}

type Giter_PingPongStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type giterPingPongStream struct {
	stream server.Stream
}

func (x *giterPingPongStream) Close() error {
	return x.stream.Close()
}

func (x *giterPingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *giterPingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *giterPingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *giterPingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
