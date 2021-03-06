// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: pkg/proto/v0/accounts.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Client API for SettingsService service

type SettingsService interface {
	Set(ctx context.Context, in *Record, opts ...client.CallOption) (*Record, error)
	Get(ctx context.Context, in *Query, opts ...client.CallOption) (*Record, error)
	List(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*Records, error)
}

type settingsService struct {
	c    client.Client
	name string
}

func NewSettingsService(name string, c client.Client) SettingsService {
	return &settingsService{
		c:    c,
		name: name,
	}
}

func (c *settingsService) Set(ctx context.Context, in *Record, opts ...client.CallOption) (*Record, error) {
	req := c.c.NewRequest(c.name, "SettingsService.Set", in)
	out := new(Record)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsService) Get(ctx context.Context, in *Query, opts ...client.CallOption) (*Record, error) {
	req := c.c.NewRequest(c.name, "SettingsService.Get", in)
	out := new(Record)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *settingsService) List(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*Records, error) {
	req := c.c.NewRequest(c.name, "SettingsService.List", in)
	out := new(Records)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SettingsService service

type SettingsServiceHandler interface {
	Set(context.Context, *Record, *Record) error
	Get(context.Context, *Query, *Record) error
	List(context.Context, *empty.Empty, *Records) error
}

func RegisterSettingsServiceHandler(s server.Server, hdlr SettingsServiceHandler, opts ...server.HandlerOption) error {
	type settingsService interface {
		Set(ctx context.Context, in *Record, out *Record) error
		Get(ctx context.Context, in *Query, out *Record) error
		List(ctx context.Context, in *empty.Empty, out *Records) error
	}
	type SettingsService struct {
		settingsService
	}
	h := &settingsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SettingsService{h}, opts...))
}

type settingsServiceHandler struct {
	SettingsServiceHandler
}

func (h *settingsServiceHandler) Set(ctx context.Context, in *Record, out *Record) error {
	return h.SettingsServiceHandler.Set(ctx, in, out)
}

func (h *settingsServiceHandler) Get(ctx context.Context, in *Query, out *Record) error {
	return h.SettingsServiceHandler.Get(ctx, in, out)
}

func (h *settingsServiceHandler) List(ctx context.Context, in *empty.Empty, out *Records) error {
	return h.SettingsServiceHandler.List(ctx, in, out)
}
