package handle

import (
	chatevloopconfig "pigeon/common/chatevloop_config"
	"pigeon/im-relay/bizpush"
	"pigeon/kitex_gen/service/imrelation/imrelation"

	"gorm.io/gorm"
)

type HandleContext struct {
	RelationCli  imrelation.Client
	EvCfgWatcher *chatevloopconfig.ChatevWatcher
	BPush        *bizpush.BizPusher
	DB           *gorm.DB
}
