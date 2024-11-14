package pushprotocol

import (
	"encoding/json"
	"errors"
	"pigeon/common/protocol"
)

func ParsePush(pkt *protocol.S2CPushMessagePacket) (data interface{}, echoCode string, err error) {
	var v1 ApplyGroupNotify
	var v2 HandleApplyNotify
	var v3 RelationChangeNotify
	var v4 SeqNotify
	var v5 SubResp
	var v6 SendMessageResp
	var v7 CreateGroupResp
	var v8 FetchAllRelationsResp
	var v9 FetchAllApplicationsResp
	var v10 ApplyGroupResp
	var v11 HandleApplyResp
	switch pkt.PushType {
	case v1.String():
		err := json.Unmarshal(pkt.Data.([]byte), &v1)
		return &v1, pkt.EchoCode(), err
	case v2.String():
		err := json.Unmarshal(pkt.Data.([]byte), &v2)
		return &v2, pkt.EchoCode(), err
	case v3.String():
		err := json.Unmarshal(pkt.Data.([]byte), &v3)
		return &v3, pkt.EchoCode(), err
	case v4.String():
		err := json.Unmarshal(pkt.Data.([]byte), &v4)
		return &v4, pkt.EchoCode(), err
	case v5.String():
		err := json.Unmarshal(pkt.Data.([]byte), &v5)
		return &v5, pkt.EchoCode(), err
	case v6.String():
		err := json.Unmarshal(pkt.Data.([]byte), &v6)
		return &v6, pkt.EchoCode(), err
	case v7.String():
		err := json.Unmarshal(pkt.Data.([]byte), &v7)
		return &v7, pkt.EchoCode(), err
	case v8.String():
		err := json.Unmarshal(pkt.Data.([]byte), &v8)
		return &v8, pkt.EchoCode(), err
	case v9.String():
		err := json.Unmarshal(pkt.Data.([]byte), &v9)
		return &v9, pkt.EchoCode(), err
	case v10.String():
		err := json.Unmarshal(pkt.Data.([]byte), &v10)
		return &v10, pkt.EchoCode(), err
	case v11.String():
		err := json.Unmarshal(pkt.Data.([]byte), &v11)
		return &v11, pkt.EchoCode(), err
	default:
		return nil, pkt.EchoCode(), errors.New("unknown pkg")
	}
}
