package protocol

import (
	"encoding/json"
	"errors"
	"fmt"
)

type PacketType int

// const (
// 	PacketTypeHeartbeat PacketType = iota

// 	PakcetTypeC2SLogin
// 	PacketTypeS2CLoginResp

// 	PacketTypeC2SLogout
// 	PacketTypeS2CLogoutResp

// 	PacketTypeC2SKickOhterDevice
// 	PackDataTypeS2CKickOhterDeviceResp

// 	// 广播包, 告诉其它设备当前用户的设备在线信息
// 	PacketTypeS2CBroadcastDeviceInfo
// 	// 其它设备踢下线的通知包
// 	PacketTypeS2COhterDeviceKick
// 	// 业务push包
// 	PacketTypeS2CPush
// )

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

type C2SQueryStatusPacket struct {
	WithEchoCode
}

type S2CQueryStatusRespPacket struct {
	WithEchoCode
	Status    string                `json:"status"`
	Username  string                `json:"username"`
	SessionId string                `json:"session_id"`
	Version   int64                 `json:"version"`
	Sessions  []*DeviceSessionEntry `json:"sessions"`
}

type C2SLoginPacket struct {
	WithEchoCode
	Username   string `json:"username"`
	Password   string `json:"password"`
	DeviceDesc string `json:"device_desc"`
}

type LoginRespCode int

const (
	LoginRespCodeUnKnown LoginRespCode = iota
	LoginRespCodeSuccess
	LoginRespCodeAuthError
	LoginRespCodeDeviceNumLimit
	LoginRespCodeAlreadyLogin
)

type S2CLoginRespPacket struct {
	WithEchoCode
	Code      LoginRespCode `json:"code"`       // code
	SessionId string        `json:"session_id"` // 只有当success == true时才有值
	// 特殊语意，deviceNumLimit, alreadyLogin也有会下面的值
	Version  int64                 `json:"version"`     // 在线信息版本号
	Sessions []*DeviceSessionEntry `json:"all_devices"` // 所有在线的设备
}

type C2SLogoutPacket struct {
	WithEchoCode
}

type S2CLogoutRespPacket struct {
	WithEchoCode
	Success bool `json:"success"` // success == false只有一种可能, 就是当前状态为未登录
	// 未登录状态下, 下面的值为空
	Version  int64                 `json:"version"`
	Sessions []*DeviceSessionEntry `json:"sessions"` // 所有在线的设备
}

// 此包用来踢下线其它设备, 不能踢下线自身
type C2SKickOtherDevicePacket struct {
	WithEchoCode
	SessionId string `json:"session_id"`
}

type S2CKickOhterDeviceRespPacket struct {
	WithEchoCode
	KickOK bool `json:"kick_ok"`
	// 无论是否踢下线成功, 都会返回当前在线设备信息
	Version  int64                 `json:"version"`
	Sessions []*DeviceSessionEntry `json:"sessions"`
}

type DeviceSessionEntry struct {
	SessionId  string `json:"session_id"`
	LoginAt    int64  `json:"login_at"`
	DeviceDesc string `json:"device_desc"`
}

// 广播其它设备的信息, 做多设备管理
type S2CDeviceInfoBroadcastPacket struct {
	Version  int64                 `json:"version"`
	Sessions []*DeviceSessionEntry `json:"sessions"`
}

// push消息
type S2CPushMessagePacket struct {
	WithEchoCode
	PushType string      `json:"push_type"`
	Data     interface{} `json:"data"`
}

// 被踢下线的通知
type S2COtherDeviceKickNotifyPacket struct {
	FromSessionId   string `json:"from_session_id"`
	FromSessionDesc string `json:"from_session_desc"`
}

type C2SBizMessagePacket struct {
	WithEchoCode
	BizType string      `json:"biz_type"`
	Data    interface{} `json:"data"`
}

type jsonHeader struct {
	PacketType string      `json:"packet_type"`
	PushType   string      `json:"push_type,omitempty"`
	BizType    string      `json:"biz_type,omitempty"`
	Data       interface{} `json:"data"`
	// 用户帮助客户端辅助定位请求的response
	EchoCode string `json:"echo_code"`
}

func dataToPacketTypeInString(data interface{}) (string, bool) {
	switch data.(type) {
	case *HeartbeatPacket:
		return "heartbeat", true
	case *C2SQueryStatusPacket:
		return "query-status", true
	case *S2CQueryStatusRespPacket:
		return "query-status-resp", true
	case *C2SLoginPacket:
		return "login", true
	case *C2SLogoutPacket:
		return "logout", true
	case *S2CLoginRespPacket:
		return "login-resp", true
	case *S2CLogoutRespPacket:
		return "logout-resp", true
	case *C2SKickOtherDevicePacket:
		return "kick-other", true
	case *S2CKickOhterDeviceRespPacket:
		return "kick-other-resp", true
	case *S2CDeviceInfoBroadcastPacket:
		return "device-info", true
	case *S2CPushMessagePacket:
		return "push-msg", true
	case *S2COtherDeviceKickNotifyPacket:
		return "other-kick-notify", true
	case *C2SBizMessagePacket:
		return "biz-msg", true
	}

	return "", false
}

// 编码包
func MustEncodePacket(data interface{}, echoCode ...string) []byte {
	packType, ok := dataToPacketTypeInString(data)
	if !ok {
		panic("packet type")
	}

	var hd jsonHeader
	if len(echoCode) == 0 {
		var ec string
		ecr, ok := data.(WithEchoCoder)
		if ok {
			ec = ecr.EchoCode()
		}
		hd = jsonHeader{
			PacketType: packType,
			Data:       data,
			EchoCode:   ec,
		}
	} else {
		hd = jsonHeader{
			PacketType: packType,
			Data:       data,
			EchoCode:   echoCode[0],
		}
	}

	// 特化这种情况
	if packType == "push-msg" {
		hd.Data = data.(*S2CPushMessagePacket).Data
		hd.PushType = data.(*S2CPushMessagePacket).PushType
	}

	if packType == "biz-msg" {
		hd.Data = data.(*C2SBizMessagePacket).Data
		hd.BizType = data.(*C2SBizMessagePacket).BizType
	}

	bs, err := json.Marshal(hd)
	if err != nil {
		panic(err)
	}
	return bs
}

func ParseC2SPacket(data []byte) (interface{}, error) {
	var header jsonHeader
	if err := json.Unmarshal(data, &header); err != nil {
		return nil, err
	}

	var err error

	switch header.PacketType {
	case "login":
		header.Data = new(C2SLoginPacket)
		header.Data.(WithEchoCoder).SetEchoCode(header.EchoCode)
		err = json.Unmarshal(data, &header)
		if !IsUsernameValid(header.Data.(*C2SLoginPacket).Username) || !IsPasswordValid(header.Data.(*C2SLoginPacket).Password) {
			return nil, errors.New("login packet validate failed: " + fmt.Sprint(header.Data.(*C2SLoginPacket)))
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
		header.Data = new(C2SKickOtherDevicePacket)
		header.Data.(WithEchoCoder).SetEchoCode(header.EchoCode)
		err = json.Unmarshal(data, &header)
	case "query-status":
		header.Data = new(C2SQueryStatusPacket)
		header.Data.(WithEchoCoder).SetEchoCode(header.EchoCode)
		err = json.Unmarshal(data, &header)
	case "biz-msg":
		header.Data = map[string]interface{}{}
		err = json.Unmarshal(data, &header)
		if err != nil {
			return nil, err
		}
		bs, err := json.Marshal(header.Data)
		if err != nil {
			return nil, err
		}
		header.Data = bs
		p := &C2SBizMessagePacket{
			BizType: header.BizType,
			Data:    bs,
		}
		p.SetEchoCode(header.EchoCode)
		return p, nil
	default:
		return nil, errors.New("packet type not found")
	}
	if err != nil {
		return nil, err
	}
	return header.Data, nil
}

func ParseS2CPacket(data []byte) (interface{}, error) {
	var header jsonHeader
	if err := json.Unmarshal(data, &header); err != nil {
		return nil, err
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
	case "device-info":
		header.Data = new(S2CDeviceInfoBroadcastPacket)
		err = json.Unmarshal(data, &header)
	case "query-status-resp":
		header.Data = new(S2CQueryStatusRespPacket)
		header.Data.(WithEchoCoder).SetEchoCode(header.EchoCode)
		err = json.Unmarshal(data, &header)
		// 特殊逻辑
	case "other-kick-notify":
		header.Data = new(S2COtherDeviceKickNotifyPacket)
		err = json.Unmarshal(data, &header)
	case "push-msg":
		header.Data = map[string]interface{}{}
		err = json.Unmarshal(data, &header)
		if err != nil {
			return nil, err
		}
		bs, err := json.Marshal(header.Data)
		if err != nil {
			return nil, err
		}
		header.Data = bs
		p := &S2CPushMessagePacket{
			PushType: header.PushType,
			Data:     bs,
		}
		p.SetEchoCode(header.EchoCode)
		return p, nil
	default:
		return nil, errors.New("unsupoort")
	}

	if err != nil {
		return nil, err
	}
	return header.Data, nil
}
