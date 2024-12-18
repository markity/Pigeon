// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package imrelay

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	base "pigeon/kitex_gen/service/base"
	evloopio "pigeon/kitex_gen/service/evloopio"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *BizMessageReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BizMessageReq[number], err)
}

func (x *BizMessageReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v base.SessionEntry
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Session = &v
	return offset, nil
}

func (x *BizMessageReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Biz, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BizMessageReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.EchoCode, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BizMessageReq) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.Data, offset, err = fastpb.ReadBytes(buf, _type)
	return offset, err
}

func (x *BizMessageResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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

func (x *CreateChatEventLoopReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateChatEventLoopReq[number], err)
}

func (x *CreateChatEventLoopReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateChatEventLoopReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.OwnerId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateChatEventLoopResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateChatEventLoopResp[number], err)
}

func (x *CreateChatEventLoopResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *RedirectToChatEventLoopReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RedirectToChatEventLoopReq[number], err)
}

func (x *RedirectToChatEventLoopReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *RedirectToChatEventLoopReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v evloopio.UniversalGroupEvloopInput
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Input = &v
	return offset, nil
}

func (x *RedirectToChatEventLoopResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_RedirectToChatEventLoopResp[number], err)
}

func (x *RedirectToChatEventLoopResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *RedirectToChatEventLoopResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v evloopio.UniversalGroupEvloopOutput
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Output = &v
	return offset, nil
}

func (x *GetLastVersionConfigReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetLastVersionConfigReq[number], err)
}

func (x *GetLastVersionConfigReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *GetLastVersionConfigReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *GetLastVersionConfigResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_GetLastVersionConfigResp[number], err)
}

func (x *GetLastVersionConfigResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.EvloopServerAddrPort, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BizMessageReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *BizMessageReq) fastWriteField1(buf []byte) (offset int) {
	if x.Session == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetSession())
	return offset
}

func (x *BizMessageReq) fastWriteField2(buf []byte) (offset int) {
	if x.Biz == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetBiz())
	return offset
}

func (x *BizMessageReq) fastWriteField3(buf []byte) (offset int) {
	if x.EchoCode == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetEchoCode())
	return offset
}

func (x *BizMessageReq) fastWriteField4(buf []byte) (offset int) {
	if len(x.Data) == 0 {
		return offset
	}
	offset += fastpb.WriteBytes(buf[offset:], 4, x.GetData())
	return offset
}

func (x *BizMessageResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	return offset
}

func (x *CreateChatEventLoopReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CreateChatEventLoopReq) fastWriteField1(buf []byte) (offset int) {
	if x.GroupId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetGroupId())
	return offset
}

func (x *CreateChatEventLoopReq) fastWriteField2(buf []byte) (offset int) {
	if x.OwnerId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetOwnerId())
	return offset
}

func (x *CreateChatEventLoopResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *CreateChatEventLoopResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *RedirectToChatEventLoopReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RedirectToChatEventLoopReq) fastWriteField1(buf []byte) (offset int) {
	if x.GroupId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetGroupId())
	return offset
}

func (x *RedirectToChatEventLoopReq) fastWriteField2(buf []byte) (offset int) {
	if x.Input == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetInput())
	return offset
}

func (x *RedirectToChatEventLoopResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *RedirectToChatEventLoopResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *RedirectToChatEventLoopResp) fastWriteField2(buf []byte) (offset int) {
	if x.Output == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetOutput())
	return offset
}

func (x *GetLastVersionConfigReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *GetLastVersionConfigReq) fastWriteField1(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetVersion())
	return offset
}

func (x *GetLastVersionConfigReq) fastWriteField2(buf []byte) (offset int) {
	if x.GroupId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetGroupId())
	return offset
}

func (x *GetLastVersionConfigResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *GetLastVersionConfigResp) fastWriteField1(buf []byte) (offset int) {
	if x.EvloopServerAddrPort == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetEvloopServerAddrPort())
	return offset
}

func (x *BizMessageReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *BizMessageReq) sizeField1() (n int) {
	if x.Session == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetSession())
	return n
}

func (x *BizMessageReq) sizeField2() (n int) {
	if x.Biz == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetBiz())
	return n
}

func (x *BizMessageReq) sizeField3() (n int) {
	if x.EchoCode == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetEchoCode())
	return n
}

func (x *BizMessageReq) sizeField4() (n int) {
	if len(x.Data) == 0 {
		return n
	}
	n += fastpb.SizeBytes(4, x.GetData())
	return n
}

func (x *BizMessageResp) Size() (n int) {
	if x == nil {
		return n
	}
	return n
}

func (x *CreateChatEventLoopReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CreateChatEventLoopReq) sizeField1() (n int) {
	if x.GroupId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetGroupId())
	return n
}

func (x *CreateChatEventLoopReq) sizeField2() (n int) {
	if x.OwnerId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetOwnerId())
	return n
}

func (x *CreateChatEventLoopResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *CreateChatEventLoopResp) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *RedirectToChatEventLoopReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RedirectToChatEventLoopReq) sizeField1() (n int) {
	if x.GroupId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetGroupId())
	return n
}

func (x *RedirectToChatEventLoopReq) sizeField2() (n int) {
	if x.Input == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetInput())
	return n
}

func (x *RedirectToChatEventLoopResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *RedirectToChatEventLoopResp) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *RedirectToChatEventLoopResp) sizeField2() (n int) {
	if x.Output == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetOutput())
	return n
}

func (x *GetLastVersionConfigReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *GetLastVersionConfigReq) sizeField1() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetVersion())
	return n
}

func (x *GetLastVersionConfigReq) sizeField2() (n int) {
	if x.GroupId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetGroupId())
	return n
}

func (x *GetLastVersionConfigResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *GetLastVersionConfigResp) sizeField1() (n int) {
	if x.EvloopServerAddrPort == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetEvloopServerAddrPort())
	return n
}

var fieldIDToName_BizMessageReq = map[int32]string{
	1: "Session",
	2: "Biz",
	3: "EchoCode",
	4: "Data",
}

var fieldIDToName_BizMessageResp = map[int32]string{}

var fieldIDToName_CreateChatEventLoopReq = map[int32]string{
	1: "GroupId",
	2: "OwnerId",
}

var fieldIDToName_CreateChatEventLoopResp = map[int32]string{
	1: "Success",
}

var fieldIDToName_RedirectToChatEventLoopReq = map[int32]string{
	1: "GroupId",
	2: "Input",
}

var fieldIDToName_RedirectToChatEventLoopResp = map[int32]string{
	1: "Success",
	2: "Output",
}

var fieldIDToName_GetLastVersionConfigReq = map[int32]string{
	1: "Version",
	2: "GroupId",
}

var fieldIDToName_GetLastVersionConfigResp = map[int32]string{
	1: "EvloopServerAddrPort",
}

var _ = base.File_base_base_proto
var _ = evloopio.File_base_evloopio_proto
