package handle

import (
	chatevloopconfig "pigeon/common/chatevloop_config"
	"pigeon/kitex_gen/service/imrelation/imrelation"
)

type HandleContext struct {
	RelationCli  imrelation.Client
	EvCfgWatcher *chatevloopconfig.ChatevWatcher
}
