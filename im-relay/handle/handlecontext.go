package handle

import "pigeon/kitex_gen/service/imrelation/imrelation"

type HandleContext struct {
	RelationCli imrelation.Client
}
