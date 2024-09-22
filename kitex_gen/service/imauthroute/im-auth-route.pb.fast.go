// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package imauthroute

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *SessionEntry) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 5:
		offset, err = x.fastReadField5(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SessionEntry[number], err)
}

func (x *SessionEntry) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.LoginAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SessionEntry) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SessionEntry) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.SessionId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SessionEntry) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.DeviceDesc, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SessionEntry) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.AdvertiseAddrPort, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LoginReq[number], err)
}

func (x *LoginReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Password, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.DeviceDesc, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LoginResp[number], err)
}

func (x *LoginResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *LoginResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.SessionId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LoginResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *LoginResp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	var v SessionEntry
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Sessions = append(x.Sessions, &v)
	return offset, nil
}

func (x *LogoutReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LogoutReq[number], err)
}

func (x *LogoutReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SessionId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *LogoutResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_LogoutResp[number], err)
}

func (x *LogoutResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *ForceOfflineReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ForceOfflineReq[number], err)
}

func (x *ForceOfflineReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SelfSessionId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ForceOfflineReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.RemoteSessionId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *ForceOfflineResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_ForceOfflineResp[number], err)
}

func (x *ForceOfflineResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *ForceOfflineResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *ForceOfflineResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v SessionEntry
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Sessions = append(x.Sessions, &v)
	return offset, nil
}

func (x *QuerySessionRouteReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_QuerySessionRouteReq[number], err)
}

func (x *QuerySessionRouteReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.SessionId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *QuerySessionRouteResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_QuerySessionRouteResp[number], err)
}

func (x *QuerySessionRouteResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *QuerySessionRouteResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v SessionEntry
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Route = &v
	return offset, nil
}

func (x *QueryUserRouteReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_QueryUserRouteReq[number], err)
}

func (x *QueryUserRouteReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Username, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *QueryUserRouteResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_QueryUserRouteResp[number], err)
}

func (x *QueryUserRouteResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *QueryUserRouteResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v SessionEntry
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Routes = append(x.Routes, &v)
	return offset, nil
}

func (x *SessionEntry) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	return offset
}

func (x *SessionEntry) fastWriteField1(buf []byte) (offset int) {
	if x.LoginAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetLoginAt())
	return offset
}

func (x *SessionEntry) fastWriteField2(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetUsername())
	return offset
}

func (x *SessionEntry) fastWriteField3(buf []byte) (offset int) {
	if x.SessionId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetSessionId())
	return offset
}

func (x *SessionEntry) fastWriteField4(buf []byte) (offset int) {
	if x.DeviceDesc == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetDeviceDesc())
	return offset
}

func (x *SessionEntry) fastWriteField5(buf []byte) (offset int) {
	if x.AdvertiseAddrPort == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetAdvertiseAddrPort())
	return offset
}

func (x *LoginReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *LoginReq) fastWriteField1(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUsername())
	return offset
}

func (x *LoginReq) fastWriteField2(buf []byte) (offset int) {
	if x.Password == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetPassword())
	return offset
}

func (x *LoginReq) fastWriteField3(buf []byte) (offset int) {
	if x.DeviceDesc == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetDeviceDesc())
	return offset
}

func (x *LoginResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *LoginResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *LoginResp) fastWriteField2(buf []byte) (offset int) {
	if x.SessionId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetSessionId())
	return offset
}

func (x *LoginResp) fastWriteField3(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetVersion())
	return offset
}

func (x *LoginResp) fastWriteField4(buf []byte) (offset int) {
	if x.Sessions == nil {
		return offset
	}
	for i := range x.GetSessions() {
		offset += fastpb.WriteMessage(buf[offset:], 4, x.GetSessions()[i])
	}
	return offset
}

func (x *LogoutReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *LogoutReq) fastWriteField1(buf []byte) (offset int) {
	if x.SessionId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetSessionId())
	return offset
}

func (x *LogoutResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *LogoutResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *ForceOfflineReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *ForceOfflineReq) fastWriteField1(buf []byte) (offset int) {
	if x.SelfSessionId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetSelfSessionId())
	return offset
}

func (x *ForceOfflineReq) fastWriteField2(buf []byte) (offset int) {
	if x.RemoteSessionId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetRemoteSessionId())
	return offset
}

func (x *ForceOfflineResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *ForceOfflineResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *ForceOfflineResp) fastWriteField2(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetVersion())
	return offset
}

func (x *ForceOfflineResp) fastWriteField3(buf []byte) (offset int) {
	if x.Sessions == nil {
		return offset
	}
	for i := range x.GetSessions() {
		offset += fastpb.WriteMessage(buf[offset:], 3, x.GetSessions()[i])
	}
	return offset
}

func (x *QuerySessionRouteReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *QuerySessionRouteReq) fastWriteField1(buf []byte) (offset int) {
	if x.SessionId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetSessionId())
	return offset
}

func (x *QuerySessionRouteResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *QuerySessionRouteResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *QuerySessionRouteResp) fastWriteField2(buf []byte) (offset int) {
	if x.Route == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetRoute())
	return offset
}

func (x *QueryUserRouteReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *QueryUserRouteReq) fastWriteField1(buf []byte) (offset int) {
	if x.Username == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetUsername())
	return offset
}

func (x *QueryUserRouteResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *QueryUserRouteResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *QueryUserRouteResp) fastWriteField2(buf []byte) (offset int) {
	if x.Routes == nil {
		return offset
	}
	for i := range x.GetRoutes() {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.GetRoutes()[i])
	}
	return offset
}

func (x *SessionEntry) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	return n
}

func (x *SessionEntry) sizeField1() (n int) {
	if x.LoginAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetLoginAt())
	return n
}

func (x *SessionEntry) sizeField2() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetUsername())
	return n
}

func (x *SessionEntry) sizeField3() (n int) {
	if x.SessionId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetSessionId())
	return n
}

func (x *SessionEntry) sizeField4() (n int) {
	if x.DeviceDesc == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetDeviceDesc())
	return n
}

func (x *SessionEntry) sizeField5() (n int) {
	if x.AdvertiseAddrPort == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetAdvertiseAddrPort())
	return n
}

func (x *LoginReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *LoginReq) sizeField1() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUsername())
	return n
}

func (x *LoginReq) sizeField2() (n int) {
	if x.Password == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetPassword())
	return n
}

func (x *LoginReq) sizeField3() (n int) {
	if x.DeviceDesc == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetDeviceDesc())
	return n
}

func (x *LoginResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *LoginResp) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *LoginResp) sizeField2() (n int) {
	if x.SessionId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetSessionId())
	return n
}

func (x *LoginResp) sizeField3() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetVersion())
	return n
}

func (x *LoginResp) sizeField4() (n int) {
	if x.Sessions == nil {
		return n
	}
	for i := range x.GetSessions() {
		n += fastpb.SizeMessage(4, x.GetSessions()[i])
	}
	return n
}

func (x *LogoutReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *LogoutReq) sizeField1() (n int) {
	if x.SessionId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetSessionId())
	return n
}

func (x *LogoutResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *LogoutResp) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *ForceOfflineReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *ForceOfflineReq) sizeField1() (n int) {
	if x.SelfSessionId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetSelfSessionId())
	return n
}

func (x *ForceOfflineReq) sizeField2() (n int) {
	if x.RemoteSessionId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetRemoteSessionId())
	return n
}

func (x *ForceOfflineResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *ForceOfflineResp) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *ForceOfflineResp) sizeField2() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetVersion())
	return n
}

func (x *ForceOfflineResp) sizeField3() (n int) {
	if x.Sessions == nil {
		return n
	}
	for i := range x.GetSessions() {
		n += fastpb.SizeMessage(3, x.GetSessions()[i])
	}
	return n
}

func (x *QuerySessionRouteReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *QuerySessionRouteReq) sizeField1() (n int) {
	if x.SessionId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetSessionId())
	return n
}

func (x *QuerySessionRouteResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *QuerySessionRouteResp) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *QuerySessionRouteResp) sizeField2() (n int) {
	if x.Route == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetRoute())
	return n
}

func (x *QueryUserRouteReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *QueryUserRouteReq) sizeField1() (n int) {
	if x.Username == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetUsername())
	return n
}

func (x *QueryUserRouteResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *QueryUserRouteResp) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *QueryUserRouteResp) sizeField2() (n int) {
	if x.Routes == nil {
		return n
	}
	for i := range x.GetRoutes() {
		n += fastpb.SizeMessage(2, x.GetRoutes()[i])
	}
	return n
}

var fieldIDToName_SessionEntry = map[int32]string{
	1: "LoginAt",
	2: "Username",
	3: "SessionId",
	4: "DeviceDesc",
	5: "AdvertiseAddrPort",
}

var fieldIDToName_LoginReq = map[int32]string{
	1: "Username",
	2: "Password",
	3: "DeviceDesc",
}

var fieldIDToName_LoginResp = map[int32]string{
	1: "Success",
	2: "SessionId",
	3: "Version",
	4: "Sessions",
}

var fieldIDToName_LogoutReq = map[int32]string{
	1: "SessionId",
}

var fieldIDToName_LogoutResp = map[int32]string{
	1: "Success",
}

var fieldIDToName_ForceOfflineReq = map[int32]string{
	1: "SelfSessionId",
	2: "RemoteSessionId",
}

var fieldIDToName_ForceOfflineResp = map[int32]string{
	1: "Success",
	2: "Version",
	3: "Sessions",
}

var fieldIDToName_QuerySessionRouteReq = map[int32]string{
	1: "SessionId",
}

var fieldIDToName_QuerySessionRouteResp = map[int32]string{
	1: "Success",
	2: "Route",
}

var fieldIDToName_QueryUserRouteReq = map[int32]string{
	1: "Username",
}

var fieldIDToName_QueryUserRouteResp = map[int32]string{
	1: "Success",
	2: "Routes",
}
