// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package imgateway

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	base "pigeon/kitex_gen/service/base"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *PushMessageReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_PushMessageReq[number], err)
}

func (x *PushMessageReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SessionId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PushMessageReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.PacketType, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PushMessageReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Data, offset, err = fastpb.ReadBytes(buf, _type)
	return offset, err
}

func (x *PushMessageReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.EchoCode, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *PushMessageResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *OtherDeviceKickReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_OtherDeviceKickReq[number], err)
}

func (x *OtherDeviceKickReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.FromSession, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OtherDeviceKickReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.FromSessionDesc, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OtherDeviceKickReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ToSession, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *OtherDeviceKickResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *BroadcastDeviceInfoReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BroadcastDeviceInfoReq[number], err)
}

func (x *BroadcastDeviceInfoReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SessionId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BroadcastDeviceInfoReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *BroadcastDeviceInfoReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v base.SessionEntry
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Sessions = append(x.Sessions, &v)
	return offset, nil
}

func (x *BroadcastDeviceInfoResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
}

func (x *PushMessageReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *PushMessageReq) fastWriteField1(buf []byte) (offset int) {
	if x.SessionId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetSessionId())
	return offset
}

func (x *PushMessageReq) fastWriteField2(buf []byte) (offset int) {
	if x.PacketType == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPacketType())
	return offset
}

func (x *PushMessageReq) fastWriteField3(buf []byte) (offset int) {
	if len(x.Data) == 0 {
		return offset
	}
	offset += fastpb.WriteBytes(buf[offset:], 3, x.GetData())
	return offset
}

func (x *PushMessageReq) fastWriteField4(buf []byte) (offset int) {
	if x.EchoCode == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetEchoCode())
	return offset
}

func (x *PushMessageResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *OtherDeviceKickReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *OtherDeviceKickReq) fastWriteField1(buf []byte) (offset int) {
	if x.FromSession == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetFromSession())
	return offset
}

func (x *OtherDeviceKickReq) fastWriteField2(buf []byte) (offset int) {
	if x.FromSessionDesc == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetFromSessionDesc())
	return offset
}

func (x *OtherDeviceKickReq) fastWriteField3(buf []byte) (offset int) {
	if x.ToSession == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetToSession())
	return offset
}

func (x *OtherDeviceKickResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *BroadcastDeviceInfoReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *BroadcastDeviceInfoReq) fastWriteField1(buf []byte) (offset int) {
	if x.SessionId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetSessionId())
	return offset
}

func (x *BroadcastDeviceInfoReq) fastWriteField2(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetVersion())
	return offset
}

func (x *BroadcastDeviceInfoReq) fastWriteField3(buf []byte) (offset int) {
	if x.Sessions == nil {
		return offset
	}
	for i := range x.GetSessions() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetSessions()[i])
	}
	return offset
}

func (x *BroadcastDeviceInfoResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *PushMessageReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *PushMessageReq) sizeField1() (n int) {
	if x.SessionId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetSessionId())
	return n
}

func (x *PushMessageReq) sizeField2() (n int) {
	if x.PacketType == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPacketType())
	return n
}

func (x *PushMessageReq) sizeField3() (n int) {
	if len(x.Data) == 0 {
		return n
	}
	n += fastpb.SizeBytes(3, x.GetData())
	return n
}

func (x *PushMessageReq) sizeField4() (n int) {
	if x.EchoCode == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetEchoCode())
	return n
}

func (x *PushMessageResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *OtherDeviceKickReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *OtherDeviceKickReq) sizeField1() (n int) {
	if x.FromSession == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetFromSession())
	return n
}

func (x *OtherDeviceKickReq) sizeField2() (n int) {
	if x.FromSessionDesc == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetFromSessionDesc())
	return n
}

func (x *OtherDeviceKickReq) sizeField3() (n int) {
	if x.ToSession == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetToSession())
	return n
}

func (x *OtherDeviceKickResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *BroadcastDeviceInfoReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *BroadcastDeviceInfoReq) sizeField1() (n int) {
	if x.SessionId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetSessionId())
	return n
}

func (x *BroadcastDeviceInfoReq) sizeField2() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetVersion())
	return n
}

func (x *BroadcastDeviceInfoReq) sizeField3() (n int) {
	if x.Sessions == nil {
		return n
	}
	for i := range x.GetSessions() {
		n += fastpb.SizeMessage(3, x.GetSessions()[i])
	}
	return n
}

func (x *BroadcastDeviceInfoResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

var fieldIDToName_PushMessageReq = map[int32]string{
	1: "SessionId",
	2: "PacketType",
	3: "Data",
	4: "EchoCode",
}

var fieldIDToName_PushMessageResp = map[int32]string{}

var fieldIDToName_OtherDeviceKickReq = map[int32]string{
	1: "FromSession",
	2: "FromSessionDesc",
	3: "ToSession",
}

var fieldIDToName_OtherDeviceKickResp = map[int32]string{}

var fieldIDToName_BroadcastDeviceInfoReq = map[int32]string{
	1: "SessionId",
	2: "Version",
	3: "Sessions",
}

var fieldIDToName_BroadcastDeviceInfoResp = map[int32]string{}

var _ = base.File_base_base_proto
