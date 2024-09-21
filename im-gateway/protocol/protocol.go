package protocol

import "encoding/json"

type PacketType int

const (
	PacketTypeHeartbeat = iota

	PakcetTypeC2SLogin
	PacketTypeS2CLoginResp

	PacketTypeC2SLogout
	PacketTypeS2CLogoutResp
)

type HeartbeatPacket struct{}

type C2SLoginPacket struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type S2CLoginRespPacket struct {
	Username  string  `json:"username"`   // echo username
	Success   bool    `json:"success"`    // 是否登录成功
	SessionId *string `json:"session_id"` // 只有当success == true时才有值
}

type C2SLogoutPacket struct{}
type S2CLogoutRespPacket struct {
	Success  bool    `json:"success"`  // success == false只有一种可能, 就是当前状态为未登录
	Username *string `json:"username"` // 只有当success == true时才有值
}

type jsonHeader struct {
	PacketType string      `json:"packet_type"`
	Data       interface{} `json:"data"`
}

func dataToPacketTypeInString(data interface{}) (string, bool) {
	switch data.(type) {
	case *HeartbeatPacket:
		return "heartbeat", true
	case *C2SLoginPacket:
		return "login", true
	case *C2SLogoutPacket:
		return "logout", true
	case *S2CLoginRespPacket:
		return "login-resp", true
	case *S2CLogoutRespPacket:
		return "logout-resp", true
	}

	return "", false
}

// 编码包
func MustEncodePacket(data interface{}) []byte {
	packType, ok := dataToPacketTypeInString(data)
	if !ok {
		panic("packet type")
	}

	hd := jsonHeader{
		PacketType: packType,
		Data:       data,
	}

	bs, err := json.Marshal(hd)
	if err != nil {
		panic(err)
	}
	return bs
}

func ParseC2SPacket(data []byte) (interface{}, bool) {
	var header jsonHeader
	if err := json.Unmarshal(data, &header); err != nil {
		return nil, false
	}

	var err error

	switch header.PacketType {
	case "login":
		header.Data = new(C2SLoginPacket)
		err = json.Unmarshal(data, &header)
	case "logout":
		header.Data = new(C2SLoginPacket)
		err = json.Unmarshal(data, &header)
	case "heartbeat":
		header.Data = new(HeartbeatPacket)
		err = json.Unmarshal(data, &header)
	}
	if err != nil {
		return nil, false
	}
	return header.Data, true
}

func ParseS2CPacket(data []byte) (interface{}, bool) {
	var header jsonHeader
	if err := json.Unmarshal(data, &header); err != nil {
		return nil, false
	}

	var err error

	switch header.PacketType {
	case "login-resp":
		header.Data = new(S2CLoginRespPacket)
		err = json.Unmarshal(data, &header)
	case "logout-resp":
		header.Data = new(S2CLogoutRespPacket)
		err = json.Unmarshal(data, &header)
	case "heartbeat":
		header.Data = new(HeartbeatPacket)
		err = json.Unmarshal(data, &header)
	}

	if err != nil {
		return nil, false
	}
	return header.Data, true
}
