// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/ops/ops.proto

package go_micro_api_ops

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

// Interface API for Ops service

type OpsService interface {
	ListRepos(ctx context.Context, in *RepoRequest, opts ...client.CallOption) (*RepoResponse, error)
	ListBranchs(ctx context.Context, in *BranchRequest, opts ...client.CallOption) (*BranchResponse, error)
	BuidProject(ctx context.Context, in *BuildRequest, opts ...client.CallOption) (*BuildResponse, error)
}

type opsService struct {
	c    client.Client
	name string
}

func NewOpsService(name string, c client.Client) OpsService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "go.micro.api.ops"
	}
	return &opsService{
		c:    c,
		name: name,
	}
}

func (c *opsService) ListRepos(ctx context.Context, in *RepoRequest, opts ...client.CallOption) (*RepoResponse, error) {
	req := c.c.NewRequest(c.name, "Ops.ListRepos", in)
	out := new(RepoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsService) ListBranchs(ctx context.Context, in *BranchRequest, opts ...client.CallOption) (*BranchResponse, error) {
	req := c.c.NewRequest(c.name, "Ops.ListBranchs", in)
	out := new(BranchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsService) BuidProject(ctx context.Context, in *BuildRequest, opts ...client.CallOption) (*BuildResponse, error) {
	req := c.c.NewRequest(c.name, "Ops.BuidProject", in)
	out := new(BuildResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Ops service

type OpsHandler interface {
	ListRepos(context.Context, *RepoRequest, *RepoResponse) error
	ListBranchs(context.Context, *BranchRequest, *BranchResponse) error
	BuidProject(context.Context, *BuildRequest, *BuildResponse) error
}

func RegisterOpsHandler(s server.Server, hdlr OpsHandler, opts ...server.HandlerOption) error {
	type ops interface {
		ListRepos(ctx context.Context, in *RepoRequest, out *RepoResponse) error
		ListBranchs(ctx context.Context, in *BranchRequest, out *BranchResponse) error
		BuidProject(ctx context.Context, in *BuildRequest, out *BuildResponse) error
	}
	type Ops struct {
		ops
	}
	h := &opsHandler{hdlr}
	return s.Handle(s.NewHandler(&Ops{h}, opts...))
}

type opsHandler struct {
	OpsHandler
}

func (h *opsHandler) ListRepos(ctx context.Context, in *RepoRequest, out *RepoResponse) error {
	return h.OpsHandler.ListRepos(ctx, in, out)
}

func (h *opsHandler) ListBranchs(ctx context.Context, in *BranchRequest, out *BranchResponse) error {
	return h.OpsHandler.ListBranchs(ctx, in, out)
}

func (h *opsHandler) BuidProject(ctx context.Context, in *BuildRequest, out *BuildResponse) error {
	return h.OpsHandler.BuidProject(ctx, in, out)
}
