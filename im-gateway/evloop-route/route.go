package evlooproute

import "sync"

// sessionId(string)->eventloop的映射, 负责把消息路由到对应的eventloop
var RouteMap sync.Map
