// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package evloopio

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	base "pigeon/kitex_gen/service/base"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *AlterGroupMemberRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlterGroupMemberRequest[number], err)
}

func (x *AlterGroupMemberRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v base.RelationEntry
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Relation = &v
	return offset, nil
}

func (x *AlterGroupMemberResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlterGroupMemberResponse[number], err)
}

func (x *AlterGroupMemberResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Code = AlterGroupMemberResponse_AlterGroupMemberResponseCode(v)
	return offset, nil
}

func (x *AlterGroupMemberResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.RelationVersion, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AlterGroupMemberResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.CurrentSeqId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AlterGroupMemberResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.ChangeAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SubscribeGroupRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SubscribeGroupRequest[number], err)
}

func (x *SubscribeGroupRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v base.SessionEntry
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Session = &v
	return offset, nil
}

func (x *SubscribeGroupRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.EchoCode, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SubscribeGroupResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SubscribeGroupResponse[number], err)
}

func (x *SubscribeGroupResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Code = SubscribeGroupResponse_SubscribeGroupCode(v)
	return offset, nil
}

func (x *SubscribeGroupResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.RelationId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SubscribeGroupResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.MaxSeqId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SendMessageRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SendMessageRequest[number], err)
}

func (x *SendMessageRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v base.SessionEntry
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Session = &v
	return offset, nil
}

func (x *SendMessageRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.EchoCode, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendMessageRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.MessageData, offset, err = fastpb.ReadBytes(buf, _type)
	return offset, err
}

func (x *SendMessageRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CheckIdempotent, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *SendMessageRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.IdempotentKey, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SendMessageResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SendMessageResponse[number], err)
}

func (x *SendMessageResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v int32
	v, offset, err = fastpb.ReadInt32(buf, _type)
	if err != nil {
		return offset, err
	}
	x.Code = SendMessageResponse_SendMessageCode(v)
	return offset, nil
}

func (x *SendMessageResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.RelationVersion, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SendMessageResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.MaxSeqId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SendMessageResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.MessageSeq, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UniversalGroupEvloopInput) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UniversalGroupEvloopInput[number], err)
}

func (x *UniversalGroupEvloopInput) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var ov UniversalGroupEvloopInput_AlterGroupMember
	x.Input = &ov
	var v AlterGroupMemberRequest
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	ov.AlterGroupMember = &v
	return offset, nil
}

func (x *UniversalGroupEvloopInput) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var ov UniversalGroupEvloopInput_SubscribeGroup
	x.Input = &ov
	var v SubscribeGroupRequest
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	ov.SubscribeGroup = &v
	return offset, nil
}

func (x *UniversalGroupEvloopInput) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var ov UniversalGroupEvloopInput_SendMessage
	x.Input = &ov
	var v SendMessageRequest
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	ov.SendMessage = &v
	return offset, nil
}

func (x *UniversalGroupEvloopOutput) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UniversalGroupEvloopOutput[number], err)
}

func (x *UniversalGroupEvloopOutput) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var ov UniversalGroupEvloopOutput_AlterGroupMember
	x.Output = &ov
	var v AlterGroupMemberResponse
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	ov.AlterGroupMember = &v
	return offset, nil
}

func (x *UniversalGroupEvloopOutput) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var ov UniversalGroupEvloopOutput_SubscribeGroup
	x.Output = &ov
	var v SubscribeGroupResponse
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	ov.SubscribeGroup = &v
	return offset, nil
}

func (x *UniversalGroupEvloopOutput) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var ov UniversalGroupEvloopOutput_SendMessage
	x.Output = &ov
	var v SendMessageResponse
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	ov.SendMessage = &v
	return offset, nil
}

func (x *AlterGroupMemberRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *AlterGroupMemberRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Relation == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetRelation())
	return offset
}

func (x *AlterGroupMemberResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *AlterGroupMemberResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, int32(x.GetCode()))
	return offset
}

func (x *AlterGroupMemberResponse) fastWriteField2(buf []byte) (offset int) {
	if x.RelationVersion == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetRelationVersion())
	return offset
}

func (x *AlterGroupMemberResponse) fastWriteField3(buf []byte) (offset int) {
	if x.CurrentSeqId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetCurrentSeqId())
	return offset
}

func (x *AlterGroupMemberResponse) fastWriteField4(buf []byte) (offset int) {
	if x.ChangeAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetChangeAt())
	return offset
}

func (x *SubscribeGroupRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *SubscribeGroupRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Session == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetSession())
	return offset
}

func (x *SubscribeGroupRequest) fastWriteField2(buf []byte) (offset int) {
	if x.EchoCode == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetEchoCode())
	return offset
}

func (x *SubscribeGroupResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *SubscribeGroupResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, int32(x.GetCode()))
	return offset
}

func (x *SubscribeGroupResponse) fastWriteField2(buf []byte) (offset int) {
	if x.RelationId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetRelationId())
	return offset
}

func (x *SubscribeGroupResponse) fastWriteField3(buf []byte) (offset int) {
	if x.MaxSeqId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetMaxSeqId())
	return offset
}

func (x *SendMessageRequest) FastWrite(buf []byte) (offset int) {
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

func (x *SendMessageRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Session == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetSession())
	return offset
}

func (x *SendMessageRequest) fastWriteField2(buf []byte) (offset int) {
	if x.EchoCode == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetEchoCode())
	return offset
}

func (x *SendMessageRequest) fastWriteField3(buf []byte) (offset int) {
	if len(x.MessageData) == 0 {
		return offset
	}
	offset += fastpb.WriteBytes(buf[offset:], 3, x.GetMessageData())
	return offset
}

func (x *SendMessageRequest) fastWriteField4(buf []byte) (offset int) {
	if !x.CheckIdempotent {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 4, x.GetCheckIdempotent())
	return offset
}

func (x *SendMessageRequest) fastWriteField5(buf []byte) (offset int) {
	if x.IdempotentKey == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 5, x.GetIdempotentKey())
	return offset
}

func (x *SendMessageResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *SendMessageResponse) fastWriteField1(buf []byte) (offset int) {
	if x.Code == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, int32(x.GetCode()))
	return offset
}

func (x *SendMessageResponse) fastWriteField2(buf []byte) (offset int) {
	if x.RelationVersion == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetRelationVersion())
	return offset
}

func (x *SendMessageResponse) fastWriteField3(buf []byte) (offset int) {
	if x.MaxSeqId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetMaxSeqId())
	return offset
}

func (x *SendMessageResponse) fastWriteField4(buf []byte) (offset int) {
	if x.MessageSeq == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetMessageSeq())
	return offset
}

func (x *UniversalGroupEvloopInput) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *UniversalGroupEvloopInput) fastWriteField1(buf []byte) (offset int) {
	if x.GetAlterGroupMember() == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetAlterGroupMember())
	return offset
}

func (x *UniversalGroupEvloopInput) fastWriteField2(buf []byte) (offset int) {
	if x.GetSubscribeGroup() == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetSubscribeGroup())
	return offset
}

func (x *UniversalGroupEvloopInput) fastWriteField3(buf []byte) (offset int) {
	if x.GetSendMessage() == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetSendMessage())
	return offset
}

func (x *UniversalGroupEvloopOutput) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *UniversalGroupEvloopOutput) fastWriteField1(buf []byte) (offset int) {
	if x.GetAlterGroupMember() == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetAlterGroupMember())
	return offset
}

func (x *UniversalGroupEvloopOutput) fastWriteField2(buf []byte) (offset int) {
	if x.GetSubscribeGroup() == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetSubscribeGroup())
	return offset
}

func (x *UniversalGroupEvloopOutput) fastWriteField3(buf []byte) (offset int) {
	if x.GetSendMessage() == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetSendMessage())
	return offset
}

func (x *AlterGroupMemberRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *AlterGroupMemberRequest) sizeField1() (n int) {
	if x.Relation == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetRelation())
	return n
}

func (x *AlterGroupMemberResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *AlterGroupMemberResponse) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, int32(x.GetCode()))
	return n
}

func (x *AlterGroupMemberResponse) sizeField2() (n int) {
	if x.RelationVersion == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetRelationVersion())
	return n
}

func (x *AlterGroupMemberResponse) sizeField3() (n int) {
	if x.CurrentSeqId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetCurrentSeqId())
	return n
}

func (x *AlterGroupMemberResponse) sizeField4() (n int) {
	if x.ChangeAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetChangeAt())
	return n
}

func (x *SubscribeGroupRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *SubscribeGroupRequest) sizeField1() (n int) {
	if x.Session == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetSession())
	return n
}

func (x *SubscribeGroupRequest) sizeField2() (n int) {
	if x.EchoCode == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetEchoCode())
	return n
}

func (x *SubscribeGroupResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *SubscribeGroupResponse) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, int32(x.GetCode()))
	return n
}

func (x *SubscribeGroupResponse) sizeField2() (n int) {
	if x.RelationId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetRelationId())
	return n
}

func (x *SubscribeGroupResponse) sizeField3() (n int) {
	if x.MaxSeqId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetMaxSeqId())
	return n
}

func (x *SendMessageRequest) Size() (n int) {
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

func (x *SendMessageRequest) sizeField1() (n int) {
	if x.Session == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetSession())
	return n
}

func (x *SendMessageRequest) sizeField2() (n int) {
	if x.EchoCode == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetEchoCode())
	return n
}

func (x *SendMessageRequest) sizeField3() (n int) {
	if len(x.MessageData) == 0 {
		return n
	}
	n += fastpb.SizeBytes(3, x.GetMessageData())
	return n
}

func (x *SendMessageRequest) sizeField4() (n int) {
	if !x.CheckIdempotent {
		return n
	}
	n += fastpb.SizeBool(4, x.GetCheckIdempotent())
	return n
}

func (x *SendMessageRequest) sizeField5() (n int) {
	if x.IdempotentKey == "" {
		return n
	}
	n += fastpb.SizeString(5, x.GetIdempotentKey())
	return n
}

func (x *SendMessageResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *SendMessageResponse) sizeField1() (n int) {
	if x.Code == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, int32(x.GetCode()))
	return n
}

func (x *SendMessageResponse) sizeField2() (n int) {
	if x.RelationVersion == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetRelationVersion())
	return n
}

func (x *SendMessageResponse) sizeField3() (n int) {
	if x.MaxSeqId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetMaxSeqId())
	return n
}

func (x *SendMessageResponse) sizeField4() (n int) {
	if x.MessageSeq == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetMessageSeq())
	return n
}

func (x *UniversalGroupEvloopInput) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *UniversalGroupEvloopInput) sizeField1() (n int) {
	if x.GetAlterGroupMember() == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetAlterGroupMember())
	return n
}

func (x *UniversalGroupEvloopInput) sizeField2() (n int) {
	if x.GetSubscribeGroup() == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetSubscribeGroup())
	return n
}

func (x *UniversalGroupEvloopInput) sizeField3() (n int) {
	if x.GetSendMessage() == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetSendMessage())
	return n
}

func (x *UniversalGroupEvloopOutput) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *UniversalGroupEvloopOutput) sizeField1() (n int) {
	if x.GetAlterGroupMember() == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetAlterGroupMember())
	return n
}

func (x *UniversalGroupEvloopOutput) sizeField2() (n int) {
	if x.GetSubscribeGroup() == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetSubscribeGroup())
	return n
}

func (x *UniversalGroupEvloopOutput) sizeField3() (n int) {
	if x.GetSendMessage() == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetSendMessage())
	return n
}

var fieldIDToName_AlterGroupMemberRequest = map[int32]string{
	1: "Relation",
}

var fieldIDToName_AlterGroupMemberResponse = map[int32]string{
	1: "Code",
	2: "RelationVersion",
	3: "CurrentSeqId",
	4: "ChangeAt",
}

var fieldIDToName_SubscribeGroupRequest = map[int32]string{
	1: "Session",
	2: "EchoCode",
}

var fieldIDToName_SubscribeGroupResponse = map[int32]string{
	1: "Code",
	2: "RelationId",
	3: "MaxSeqId",
}

var fieldIDToName_SendMessageRequest = map[int32]string{
	1: "Session",
	2: "EchoCode",
	3: "MessageData",
	4: "CheckIdempotent",
	5: "IdempotentKey",
}

var fieldIDToName_SendMessageResponse = map[int32]string{
	1: "Code",
	2: "RelationVersion",
	3: "MaxSeqId",
	4: "MessageSeq",
}

var fieldIDToName_UniversalGroupEvloopInput = map[int32]string{
	1: "AlterGroupMember",
	2: "SubscribeGroup",
	3: "SendMessage",
}

var fieldIDToName_UniversalGroupEvloopOutput = map[int32]string{
	1: "AlterGroupMember",
	2: "SubscribeGroup",
	3: "SendMessage",
}

var _ = base.File_base_base_proto
