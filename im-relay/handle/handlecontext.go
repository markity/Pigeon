package handle

import (
	chatevloopconfig "pigeon/common/chatevloop-config"
	"pigeon/kitex_gen/service/imrelation/imrelation"
)

type HandleContext struct {
	RelationCli  imrelation.Client
	EvCfgWatcher *chatevloopconfig.ChatevWatcher
}
