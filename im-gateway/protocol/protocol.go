package protocol

import (
	"encoding/json"
)

type PacketType int

const (
	PacketTypeHeartbeat PacketType = iota

	PakcetTypeC2SLogin
	PacketTypeS2CLoginResp

	PacketTypeC2SLogout
	PacketTypeS2CLogoutResp

	PacketTypeC2SKickOhterDevice
	PackDataTypeS2CKickOhterDeviceResp

	// 广播包, 告诉其它设备当前用户的设备在线信息
	PacketTypeS2CBroadcastDeviceInfo
	// 其它设备踢下线的通知包
	PacketTypeS2COhterDeviceKick
)

type WithEchoCode struct {
	code string
}

func (w *WithEchoCode) EchoCode() string {
	return w.code
}

func (w *WithEchoCode) SetEchoCode(code string) {
	w.code = code
}

type WithEchoCoder interface {
	EchoCode() string
	SetEchoCode(code string)
}

type HeartbeatPacket struct {
	WithEchoCode
}

type C2SLoginPacket struct {
	WithEchoCode
	Username   string `json:"username"`
	Password   string `json:"password"`
	DeviceDesc string `json:"device_desc"`
}
type S2CLoginRespPacket struct {
	WithEchoCode
	Username           string                `json:"username"`    // echo username
	Success            bool                  `json:"success"`     // 是否登录成功
	SessionId          *string               `json:"session_id"`  // 只有当success == true时才有值
	DeviceSessionEntry []*DeviceSessionEntry `json:"all_devices"` // 所有在线的设备
}

type C2SLogoutPacket struct {
	WithEchoCode
}
type S2CLogoutRespPacket struct {
	WithEchoCode
	Success  bool    `json:"success"`  // success == false只有一种可能, 就是当前状态为未登录
	Username *string `json:"username"` // 只有当success == true时才有值
}

// 此包用来踢下线其它设备, 不能踢下线自身
type C2SKickOhterDevicePacket struct {
	WithEchoCode
	SessionId string `json:"session_id"`
}

type S2CKickOhterDeviceRespPacket struct {
	WithEchoCode
	KickOK     bool                  `json:"kick_ok"`
	Version    int64                 `json:"version"`
	NewDevices []*DeviceSessionEntry `json:"new_devices"`
}

type DeviceSessionEntry struct {
	WithEchoCode
	SessionId  string `json:"session_id"`
	LoginAt    int64  `json:"login_at"`
	DeviceDesc string `json:"device_desc"`
}

// 广播其它设备的信息, 做多设备管理
type S2CDeviceInfoBroadcastPacket struct {
	Version int64                 `json:"version"`
	Devices []*DeviceSessionEntry `json:"devices"`
}

// push消息
type S2CPushMessagePacket struct {
	Data interface{} `json:"data"`
}

// 被踢下线的通知
type S2COtherDeviceKickNotify struct {
	FromSessionId   string `json:"from_session_id"`
	FromSessionDesc string `json:"from_session_desc"`
}

type jsonHeader struct {
	PacketType string      `json:"packet_type"`
	Data       interface{} `json:"data"`
	// 用户帮助客户端辅助定位请求的response
	EchoCode string `json:"echo_code"`
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
	case *C2SKickOhterDevicePacket:
		return "kick-other", true
	case *S2CKickOhterDeviceRespPacket:
		return "kick-other-resp", true
	case *S2CDeviceInfoBroadcastPacket:
		return "device-info", true
	case *S2CPushMessagePacket:
		return "push-msg", true
	case *S2COtherDeviceKickNotify:
		return "other-kick-notify", true
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
		EchoCode:   data.(WithEchoCoder).EchoCode(),
	}

	// 特化这种情况
	if packType == "push-msg" {
		hd.Data = data.(*S2CPushMessagePacket).Data
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
		header.Data.(WithEchoCoder).SetEchoCode(header.EchoCode)
		err = json.Unmarshal(data, &header)
		if !IsUsernameValid(header.Data.(*C2SLoginPacket).Username) || !IsPasswordValid(header.Data.(*C2SLoginPacket).Password) {
			return nil, false
		}
	case "logout":
		header.Data = new(C2SLogoutPacket)
		header.Data.(WithEchoCoder).SetEchoCode(header.EchoCode)
		err = json.Unmarshal(data, &header)
	case "heartbeat":
		header.Data = new(HeartbeatPacket)
		header.Data.(WithEchoCoder).SetEchoCode(header.EchoCode)
		err = json.Unmarshal(data, &header)
	case "kick-other":
		header.Data = new(C2SKickOhterDevicePacket)
		header.Data.(WithEchoCoder).SetEchoCode(header.EchoCode)
		err = json.Unmarshal(data, &header)
	default:
		return nil, false
		//panic "packet type
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
		header.Data.(WithEchoCoder).SetEchoCode(header.EchoCode)
		err = json.Unmarshal(data, &header)
	case "logout-resp":
		header.Data = new(S2CLogoutRespPacket)
		header.Data.(WithEchoCoder).SetEchoCode(header.EchoCode)
		err = json.Unmarshal(data, &header)
	case "heartbeat":
		header.Data = new(HeartbeatPacket)
		header.Data.(WithEchoCoder).SetEchoCode(header.EchoCode)
		err = json.Unmarshal(data, &header)
	case "kick-other-resp":
		header.Data = new(S2CKickOhterDeviceRespPacket)
		header.Data.(WithEchoCoder).SetEchoCode(header.EchoCode)
		err = json.Unmarshal(data, &header)
	case "push-msg":
		panic("unsupport")
	}

	if err != nil {
		return nil, false
	}
	return header.Data, true
}
