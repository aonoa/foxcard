// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v3.19.1
// source: api/card/v1/card.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCardCardHeartbeat = "/api.card.v1.Card/CardHeartbeat"
const OperationCardCardLogin = "/api.card.v1.Card/CardLogin"
const OperationCardCardLogout = "/api.card.v1.Card/CardLogout"
const OperationCardCloneCard = "/api.card.v1.Card/CloneCard"
const OperationCardCreateCard = "/api.card.v1.Card/CreateCard"
const OperationCardDelCard = "/api.card.v1.Card/DelCard"
const OperationCardDurationCard = "/api.card.v1.Card/DurationCard"
const OperationCardFrozenCard = "/api.card.v1.Card/FrozenCard"
const OperationCardThawCard = "/api.card.v1.Card/ThawCard"

type CardHTTPServer interface {
	// CardHeartbeat 卡密心跳
	CardHeartbeat(context.Context, *HeartbeatRequest) (*HeartbeatReply, error)
	// CardLogin 卡密登陆
	CardLogin(context.Context, *LoginRequest) (*LoginReply, error)
	// CardLogout 卡密退出登陆
	CardLogout(context.Context, *LogoutRequest) (*LogoutReply, error)
	// CloneCard 卡密克隆
	CloneCard(context.Context, *CloneCardRequest) (*CardReply, error)
	// CreateCard 生成卡密
	CreateCard(context.Context, *CreateCardRequest) (*CardReply, error)
	// DelCard 销毁卡
	DelCard(context.Context, *DelCardRequest) (*DelCardReply, error)
	// DurationCard 时长控制
	DurationCard(context.Context, *DurationCardRequest) (*CardReply, error)
	// FrozenCard 冻结卡
	FrozenCard(context.Context, *FrozenCardRequest) (*FrozenCardReply, error)
	// ThawCard 解冻卡
	ThawCard(context.Context, *ThawCardRequest) (*ThawCardReply, error)
}

func RegisterCardHTTPServer(s *http.Server, srv CardHTTPServer) {
	r := s.Route("/")
	r.POST("/card/login", _Card_CardLogin0_HTTP_Handler(srv))
	r.POST("/card/heartbeat", _Card_CardHeartbeat0_HTTP_Handler(srv))
	r.POST("/card/logout", _Card_CardLogout0_HTTP_Handler(srv))
	r.POST("/card/create", _Card_CreateCard0_HTTP_Handler(srv))
	r.POST("/card/frozen", _Card_FrozenCard0_HTTP_Handler(srv))
	r.POST("/card/thaw", _Card_ThawCard0_HTTP_Handler(srv))
	r.POST("/card/del", _Card_DelCard0_HTTP_Handler(srv))
	r.POST("/card/clone", _Card_CloneCard0_HTTP_Handler(srv))
	r.POST("/card/duration", _Card_DurationCard0_HTTP_Handler(srv))
}

func _Card_CardLogin0_HTTP_Handler(srv CardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCardCardLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CardLogin(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginReply)
		return ctx.Result(200, reply)
	}
}

func _Card_CardHeartbeat0_HTTP_Handler(srv CardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HeartbeatRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCardCardHeartbeat)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CardHeartbeat(ctx, req.(*HeartbeatRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HeartbeatReply)
		return ctx.Result(200, reply)
	}
}

func _Card_CardLogout0_HTTP_Handler(srv CardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LogoutRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCardCardLogout)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CardLogout(ctx, req.(*LogoutRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LogoutReply)
		return ctx.Result(200, reply)
	}
}

func _Card_CreateCard0_HTTP_Handler(srv CardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateCardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCardCreateCard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateCard(ctx, req.(*CreateCardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CardReply)
		return ctx.Result(200, reply)
	}
}

func _Card_FrozenCard0_HTTP_Handler(srv CardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in FrozenCardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCardFrozenCard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.FrozenCard(ctx, req.(*FrozenCardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*FrozenCardReply)
		return ctx.Result(200, reply)
	}
}

func _Card_ThawCard0_HTTP_Handler(srv CardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ThawCardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCardThawCard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ThawCard(ctx, req.(*ThawCardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ThawCardReply)
		return ctx.Result(200, reply)
	}
}

func _Card_DelCard0_HTTP_Handler(srv CardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DelCardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCardDelCard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DelCard(ctx, req.(*DelCardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DelCardReply)
		return ctx.Result(200, reply)
	}
}

func _Card_CloneCard0_HTTP_Handler(srv CardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CloneCardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCardCloneCard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CloneCard(ctx, req.(*CloneCardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CardReply)
		return ctx.Result(200, reply)
	}
}

func _Card_DurationCard0_HTTP_Handler(srv CardHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DurationCardRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCardDurationCard)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DurationCard(ctx, req.(*DurationCardRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CardReply)
		return ctx.Result(200, reply)
	}
}

type CardHTTPClient interface {
	CardHeartbeat(ctx context.Context, req *HeartbeatRequest, opts ...http.CallOption) (rsp *HeartbeatReply, err error)
	CardLogin(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginReply, err error)
	CardLogout(ctx context.Context, req *LogoutRequest, opts ...http.CallOption) (rsp *LogoutReply, err error)
	CloneCard(ctx context.Context, req *CloneCardRequest, opts ...http.CallOption) (rsp *CardReply, err error)
	CreateCard(ctx context.Context, req *CreateCardRequest, opts ...http.CallOption) (rsp *CardReply, err error)
	DelCard(ctx context.Context, req *DelCardRequest, opts ...http.CallOption) (rsp *DelCardReply, err error)
	DurationCard(ctx context.Context, req *DurationCardRequest, opts ...http.CallOption) (rsp *CardReply, err error)
	FrozenCard(ctx context.Context, req *FrozenCardRequest, opts ...http.CallOption) (rsp *FrozenCardReply, err error)
	ThawCard(ctx context.Context, req *ThawCardRequest, opts ...http.CallOption) (rsp *ThawCardReply, err error)
}

type CardHTTPClientImpl struct {
	cc *http.Client
}

func NewCardHTTPClient(client *http.Client) CardHTTPClient {
	return &CardHTTPClientImpl{client}
}

func (c *CardHTTPClientImpl) CardHeartbeat(ctx context.Context, in *HeartbeatRequest, opts ...http.CallOption) (*HeartbeatReply, error) {
	var out HeartbeatReply
	pattern := "/card/heartbeat"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCardCardHeartbeat))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CardHTTPClientImpl) CardLogin(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginReply, error) {
	var out LoginReply
	pattern := "/card/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCardCardLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CardHTTPClientImpl) CardLogout(ctx context.Context, in *LogoutRequest, opts ...http.CallOption) (*LogoutReply, error) {
	var out LogoutReply
	pattern := "/card/logout"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCardCardLogout))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CardHTTPClientImpl) CloneCard(ctx context.Context, in *CloneCardRequest, opts ...http.CallOption) (*CardReply, error) {
	var out CardReply
	pattern := "/card/clone"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCardCloneCard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CardHTTPClientImpl) CreateCard(ctx context.Context, in *CreateCardRequest, opts ...http.CallOption) (*CardReply, error) {
	var out CardReply
	pattern := "/card/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCardCreateCard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CardHTTPClientImpl) DelCard(ctx context.Context, in *DelCardRequest, opts ...http.CallOption) (*DelCardReply, error) {
	var out DelCardReply
	pattern := "/card/del"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCardDelCard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CardHTTPClientImpl) DurationCard(ctx context.Context, in *DurationCardRequest, opts ...http.CallOption) (*CardReply, error) {
	var out CardReply
	pattern := "/card/duration"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCardDurationCard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CardHTTPClientImpl) FrozenCard(ctx context.Context, in *FrozenCardRequest, opts ...http.CallOption) (*FrozenCardReply, error) {
	var out FrozenCardReply
	pattern := "/card/frozen"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCardFrozenCard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CardHTTPClientImpl) ThawCard(ctx context.Context, in *ThawCardRequest, opts ...http.CallOption) (*ThawCardReply, error) {
	var out ThawCardReply
	pattern := "/card/thaw"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCardThawCard))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}