package evloop

import (
	subscribemanager "pigeon/im-chat-evloop/evloop/subscribe_manager"
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/evloopio"
)

type evInput struct {
	input  interface{}
	output chan interface{}
}

type evInputMove struct{}

type evInputStop struct{}

type evInputUniversal struct {
	Input *evloopio.UniversalGroupEvloopInput
}

// 拿到evloop的全部状态, 用于迁移
type EvOutputMove struct {
	err       error // queueMessageError, 可能是stop
	ChatId    string
	OwnerId   string
	CreatedAt int64
	SeqId     int64
	Relations map[string]*base.RelationEntry
	Subs      map[string](map[string]*subscribemanager.Sub)
}

type EvOutputStop struct {
	err error // queueMessageError, 可能是running
}

type EvOutputUniversal struct {
	err    error // queueMessageError, 可能是stop或moving
	Output *evloopio.UniversalGroupEvloopOutput
}

type NewChatEvLoopInput struct {
	ChatId string
	// 群主id
	OwnerId string
}
