package main

import (
	"fmt"
	bizprotocol "pigeon/common/biz_protocol"
	"pigeon/common/protocol"
)

func main() {
	pullRelationsPkt := protocol.C2SBizMessagePacket{
		BizType: (&bizprotocol.BizPullRelations{}).String(),
		Data:    &bizprotocol.BizPullRelations{},
	}
	fmt.Println(string(protocol.MustEncodePacket(&pullRelationsPkt)))
}
