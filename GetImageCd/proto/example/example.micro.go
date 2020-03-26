// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: example.proto

package go_micro_srv_GetImageCd

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

// Client API for Example service

type ExampleService interface {
	GetImageCd(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type exampleService struct {
	c    client.Client
	name string
}

func NewExampleService(name string, c client.Client) ExampleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.srv.GetImageCd"
	}
	return &exampleService{
		c:    c,
		name: name,
	}
}

func (c *exampleService) GetImageCd(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Example.GetImageCd", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Example service

type ExampleHandler interface {
	GetImageCd(context.Context, *Request, *Response) error
}

func RegisterExampleHandler(s server.Server, hdlr ExampleHandler, opts ...server.HandlerOption) error {
	type example interface {
		GetImageCd(ctx context.Context, in *Request, out *Response) error
	}
	type Example struct {
		example
	}
	h := &exampleHandler{hdlr}
	return s.Handle(s.NewHandler(&Example{h}, opts...))
}

type exampleHandler struct {
	ExampleHandler
}

func (h *exampleHandler) GetImageCd(ctx context.Context, in *Request, out *Response) error {
	return h.ExampleHandler.GetImageCd(ctx, in, out)
}
