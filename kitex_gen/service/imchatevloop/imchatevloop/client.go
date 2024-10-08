// Code generated by Kitex v0.11.3. DO NOT EDIT.

package imchatevloop

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	imchatevloop "pigeon/kitex_gen/service/imchatevloop"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateGroup(ctx context.Context, Req *imchatevloop.CreateGroupRequest, callOptions ...callopt.Option) (r *imchatevloop.CreateGroupResponse, err error)
	AlterGroupMember(ctx context.Context, Req *imchatevloop.AlterGroupMemberRequest, callOptions ...callopt.Option) (r *imchatevloop.AlterGroupMemberResponse, err error)
	SubscribeGroup(ctx context.Context, Req *imchatevloop.SubscribeGroupRequest, callOptions ...callopt.Option) (r *imchatevloop.SubscribeGroupResponse, err error)
	SendMessage(ctx context.Context, Req *imchatevloop.SendMessageRequest, callOptions ...callopt.Option) (r *imchatevloop.SendMessageResponse, err error)
	DisbandGroup(ctx context.Context, Req *imchatevloop.DisbandGroupRequest, callOptions ...callopt.Option) (r *imchatevloop.DisbandGroupResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kIMChatEvloopClient{
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

type kIMChatEvloopClient struct {
	*kClient
}

func (p *kIMChatEvloopClient) CreateGroup(ctx context.Context, Req *imchatevloop.CreateGroupRequest, callOptions ...callopt.Option) (r *imchatevloop.CreateGroupResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateGroup(ctx, Req)
}

func (p *kIMChatEvloopClient) AlterGroupMember(ctx context.Context, Req *imchatevloop.AlterGroupMemberRequest, callOptions ...callopt.Option) (r *imchatevloop.AlterGroupMemberResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AlterGroupMember(ctx, Req)
}

func (p *kIMChatEvloopClient) SubscribeGroup(ctx context.Context, Req *imchatevloop.SubscribeGroupRequest, callOptions ...callopt.Option) (r *imchatevloop.SubscribeGroupResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SubscribeGroup(ctx, Req)
}

func (p *kIMChatEvloopClient) SendMessage(ctx context.Context, Req *imchatevloop.SendMessageRequest, callOptions ...callopt.Option) (r *imchatevloop.SendMessageResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendMessage(ctx, Req)
}

func (p *kIMChatEvloopClient) DisbandGroup(ctx context.Context, Req *imchatevloop.DisbandGroupRequest, callOptions ...callopt.Option) (r *imchatevloop.DisbandGroupResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DisbandGroup(ctx, Req)
}