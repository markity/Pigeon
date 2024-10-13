// Code generated by Kitex v0.9.1. DO NOT EDIT.

package imchatevloop

import (
	server "github.com/cloudwego/kitex/server"
	imchatevloop "pigeon/kitex_gen/service/imchatevloop"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler imchatevloop.IMChatEvloop, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}