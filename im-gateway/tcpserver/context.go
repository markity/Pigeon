package tcpserver

import (
	"sync/atomic"

	eventloop "github.com/markity/go-reactor/pkg/event_loop"
)

var evloopContextKey = "_ev_loop_ctx"

type EvloopContext struct {
	RelayCli     interface{}
	AuthRouteCli interface{}
	ConnMetrics  *atomic.Uint64
}

func SetUpEvLoopContext(loop eventloop.EventLoop, ctx *EvloopContext) {
	loop.SetContext(evloopContextKey, ctx)
}

func MustLoadEvLoopContext(loop eventloop.EventLoop) *EvloopContext {
	c, ok := loop.GetContext(evloopContextKey)
	if !ok {
		panic("get context")
	}
	return c.(*EvloopContext)
}
