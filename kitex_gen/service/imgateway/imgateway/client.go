// Code generated by Kitex v0.9.1. DO NOT EDIT.

package imgateway

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	imgateway "pigeon/kitex_gen/service/imgateway"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PushMessage(ctx context.Context, Req *imgateway.PushMessageReq, callOptions ...callopt.Option) (r *imgateway.PushMessageResp, err error)
	OtherDeviceKick(ctx context.Context, Req *imgateway.OtherDeviceKickReq, callOptions ...callopt.Option) (r *imgateway.OtherDeviceKickResp, err error)
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
	return &kIMGatewayClient{
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

type kIMGatewayClient struct {
	*kClient
}

func (p *kIMGatewayClient) PushMessage(ctx context.Context, Req *imgateway.PushMessageReq, callOptions ...callopt.Option) (r *imgateway.PushMessageResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PushMessage(ctx, Req)
}

func (p *kIMGatewayClient) OtherDeviceKick(ctx context.Context, Req *imgateway.OtherDeviceKickReq, callOptions ...callopt.Option) (r *imgateway.OtherDeviceKickResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.OtherDeviceKick(ctx, Req)
}
