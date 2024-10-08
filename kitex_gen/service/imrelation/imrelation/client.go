// Code generated by Kitex v0.11.3. DO NOT EDIT.

package imrelation

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	imrelation "pigeon/kitex_gen/service/imrelation"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	CreateGroup(ctx context.Context, Req *imrelation.CreateGroupReq, callOptions ...callopt.Option) (r *imrelation.CreateGroupResp, err error)
	GetGroupInfo(ctx context.Context, Req *imrelation.GetGroupInfoReq, callOptions ...callopt.Option) (r *imrelation.GroupInfo, err error)
	FetchAllRelations(ctx context.Context, Req *imrelation.FetchAllRelationsReq, callOptions ...callopt.Option) (r *imrelation.FetchAllRelationsResp, err error)
	FetchAllApplications(ctx context.Context, Req *imrelation.FetchAllApplicationsReq, callOptions ...callopt.Option) (r *imrelation.FetchAllApplicationsResp, err error)
	ApplyGroup(ctx context.Context, Req *imrelation.ApplyGroupReq, callOptions ...callopt.Option) (r *imrelation.ApplyGroupReqResp, err error)
	HandleApply(ctx context.Context, Req *imrelation.HandleApplyReq, callOptions ...callopt.Option) (r *imrelation.HandleApplyResp, err error)
	QuitGroup(ctx context.Context, Req *imrelation.QuitGroupReq, callOptions ...callopt.Option) (r *imrelation.QuitGroupResp, err error)
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
	return &kIMRelationClient{
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

type kIMRelationClient struct {
	*kClient
}

func (p *kIMRelationClient) CreateGroup(ctx context.Context, Req *imrelation.CreateGroupReq, callOptions ...callopt.Option) (r *imrelation.CreateGroupResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateGroup(ctx, Req)
}

func (p *kIMRelationClient) GetGroupInfo(ctx context.Context, Req *imrelation.GetGroupInfoReq, callOptions ...callopt.Option) (r *imrelation.GroupInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetGroupInfo(ctx, Req)
}

func (p *kIMRelationClient) FetchAllRelations(ctx context.Context, Req *imrelation.FetchAllRelationsReq, callOptions ...callopt.Option) (r *imrelation.FetchAllRelationsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FetchAllRelations(ctx, Req)
}

func (p *kIMRelationClient) FetchAllApplications(ctx context.Context, Req *imrelation.FetchAllApplicationsReq, callOptions ...callopt.Option) (r *imrelation.FetchAllApplicationsResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FetchAllApplications(ctx, Req)
}

func (p *kIMRelationClient) ApplyGroup(ctx context.Context, Req *imrelation.ApplyGroupReq, callOptions ...callopt.Option) (r *imrelation.ApplyGroupReqResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.ApplyGroup(ctx, Req)
}

func (p *kIMRelationClient) HandleApply(ctx context.Context, Req *imrelation.HandleApplyReq, callOptions ...callopt.Option) (r *imrelation.HandleApplyResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.HandleApply(ctx, Req)
}

func (p *kIMRelationClient) QuitGroup(ctx context.Context, Req *imrelation.QuitGroupReq, callOptions ...callopt.Option) (r *imrelation.QuitGroupResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QuitGroup(ctx, Req)
}