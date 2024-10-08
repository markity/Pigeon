// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package imchatevloop

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
	base "pigeon/kitex_gen/service/base"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *CreateGroupRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateGroupRequest[number], err)
}

func (x *CreateGroupRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateGroupRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateGroupRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.GroupOwnerId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateGroupResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateGroupResponse[number], err)
}

func (x *CreateGroupResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *CreateGroupResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AlterGroupMemberRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_AlterGroupMemberRequest[number], err)
}

func (x *AlterGroupMemberRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AlterGroupMemberRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *AlterGroupMemberRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.IsAdd, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *AlterGroupMemberRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.MemberId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *AlterGroupMemberRequest) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.RelationId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
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
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *AlterGroupMemberResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_SubscribeGroupRequest[number], err)
}

func (x *SubscribeGroupRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SubscribeGroupRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SubscribeGroupRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.GwAdvertiseAddrPort, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *SubscribeGroupRequest) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.SessionId, offset, err = fastpb.ReadString(buf, _type)
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
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SubscribeGroupResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.RelationId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SubscribeGroupResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
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
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SendMessageRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadInt64(buf, _type)
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
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SendMessageResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.RelationId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SendMessageResponse) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.MaxSeqId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *SendMessageResponse) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.MessageSeq, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DisbandGroupRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DisbandGroupRequest[number], err)
}

func (x *DisbandGroupRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DisbandGroupRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DisbandGroupResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DisbandGroupResponse[number], err)
}

func (x *DisbandGroupResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *DisbandGroupResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateGroupRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CreateGroupRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetVersion())
	return offset
}

func (x *CreateGroupRequest) fastWriteField2(buf []byte) (offset int) {
	if x.GroupId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetGroupId())
	return offset
}

func (x *CreateGroupRequest) fastWriteField3(buf []byte) (offset int) {
	if x.GroupOwnerId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetGroupOwnerId())
	return offset
}

func (x *CreateGroupResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CreateGroupResponse) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *CreateGroupResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetVersion())
	return offset
}

func (x *AlterGroupMemberRequest) FastWrite(buf []byte) (offset int) {
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

func (x *AlterGroupMemberRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetVersion())
	return offset
}

func (x *AlterGroupMemberRequest) fastWriteField2(buf []byte) (offset int) {
	if x.GroupId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetGroupId())
	return offset
}

func (x *AlterGroupMemberRequest) fastWriteField3(buf []byte) (offset int) {
	if !x.IsAdd {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 3, x.GetIsAdd())
	return offset
}

func (x *AlterGroupMemberRequest) fastWriteField4(buf []byte) (offset int) {
	if x.MemberId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetMemberId())
	return offset
}

func (x *AlterGroupMemberRequest) fastWriteField5(buf []byte) (offset int) {
	if x.RelationId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetRelationId())
	return offset
}

func (x *AlterGroupMemberResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *AlterGroupMemberResponse) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *AlterGroupMemberResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetVersion())
	return offset
}

func (x *SubscribeGroupRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	return offset
}

func (x *SubscribeGroupRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetVersion())
	return offset
}

func (x *SubscribeGroupRequest) fastWriteField2(buf []byte) (offset int) {
	if x.GroupId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetGroupId())
	return offset
}

func (x *SubscribeGroupRequest) fastWriteField3(buf []byte) (offset int) {
	if x.GwAdvertiseAddrPort == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetGwAdvertiseAddrPort())
	return offset
}

func (x *SubscribeGroupRequest) fastWriteField4(buf []byte) (offset int) {
	if x.SessionId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.GetSessionId())
	return offset
}

func (x *SubscribeGroupResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
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
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetVersion())
	return offset
}

func (x *SubscribeGroupResponse) fastWriteField3(buf []byte) (offset int) {
	if x.RelationId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetRelationId())
	return offset
}

func (x *SubscribeGroupResponse) fastWriteField4(buf []byte) (offset int) {
	if x.MaxSeqId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetMaxSeqId())
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
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetVersion())
	return offset
}

func (x *SendMessageRequest) fastWriteField2(buf []byte) (offset int) {
	if x.GroupId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetGroupId())
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
	offset += x.fastWriteField5(buf[offset:])
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
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetVersion())
	return offset
}

func (x *SendMessageResponse) fastWriteField3(buf []byte) (offset int) {
	if x.RelationId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetRelationId())
	return offset
}

func (x *SendMessageResponse) fastWriteField4(buf []byte) (offset int) {
	if x.MaxSeqId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetMaxSeqId())
	return offset
}

func (x *SendMessageResponse) fastWriteField5(buf []byte) (offset int) {
	if x.MessageSeq == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetMessageSeq())
	return offset
}

func (x *DisbandGroupRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DisbandGroupRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetVersion())
	return offset
}

func (x *DisbandGroupRequest) fastWriteField2(buf []byte) (offset int) {
	if x.GroupId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetGroupId())
	return offset
}

func (x *DisbandGroupResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DisbandGroupResponse) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *DisbandGroupResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetVersion())
	return offset
}

func (x *CreateGroupRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CreateGroupRequest) sizeField1() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetVersion())
	return n
}

func (x *CreateGroupRequest) sizeField2() (n int) {
	if x.GroupId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetGroupId())
	return n
}

func (x *CreateGroupRequest) sizeField3() (n int) {
	if x.GroupOwnerId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetGroupOwnerId())
	return n
}

func (x *CreateGroupResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CreateGroupResponse) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *CreateGroupResponse) sizeField2() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetVersion())
	return n
}

func (x *AlterGroupMemberRequest) Size() (n int) {
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

func (x *AlterGroupMemberRequest) sizeField1() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetVersion())
	return n
}

func (x *AlterGroupMemberRequest) sizeField2() (n int) {
	if x.GroupId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetGroupId())
	return n
}

func (x *AlterGroupMemberRequest) sizeField3() (n int) {
	if !x.IsAdd {
		return n
	}
	n += fastpb.SizeBool(3, x.GetIsAdd())
	return n
}

func (x *AlterGroupMemberRequest) sizeField4() (n int) {
	if x.MemberId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetMemberId())
	return n
}

func (x *AlterGroupMemberRequest) sizeField5() (n int) {
	if x.RelationId == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetRelationId())
	return n
}

func (x *AlterGroupMemberResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *AlterGroupMemberResponse) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *AlterGroupMemberResponse) sizeField2() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetVersion())
	return n
}

func (x *SubscribeGroupRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	return n
}

func (x *SubscribeGroupRequest) sizeField1() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetVersion())
	return n
}

func (x *SubscribeGroupRequest) sizeField2() (n int) {
	if x.GroupId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetGroupId())
	return n
}

func (x *SubscribeGroupRequest) sizeField3() (n int) {
	if x.GwAdvertiseAddrPort == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetGwAdvertiseAddrPort())
	return n
}

func (x *SubscribeGroupRequest) sizeField4() (n int) {
	if x.SessionId == "" {
		return n
	}
	n += fastpb.SizeString(4, x.GetSessionId())
	return n
}

func (x *SubscribeGroupResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
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
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetVersion())
	return n
}

func (x *SubscribeGroupResponse) sizeField3() (n int) {
	if x.RelationId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetRelationId())
	return n
}

func (x *SubscribeGroupResponse) sizeField4() (n int) {
	if x.MaxSeqId == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetMaxSeqId())
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
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetVersion())
	return n
}

func (x *SendMessageRequest) sizeField2() (n int) {
	if x.GroupId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetGroupId())
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
	n += x.sizeField5()
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
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetVersion())
	return n
}

func (x *SendMessageResponse) sizeField3() (n int) {
	if x.RelationId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetRelationId())
	return n
}

func (x *SendMessageResponse) sizeField4() (n int) {
	if x.MaxSeqId == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetMaxSeqId())
	return n
}

func (x *SendMessageResponse) sizeField5() (n int) {
	if x.MessageSeq == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetMessageSeq())
	return n
}

func (x *DisbandGroupRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DisbandGroupRequest) sizeField1() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetVersion())
	return n
}

func (x *DisbandGroupRequest) sizeField2() (n int) {
	if x.GroupId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetGroupId())
	return n
}

func (x *DisbandGroupResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DisbandGroupResponse) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *DisbandGroupResponse) sizeField2() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetVersion())
	return n
}

var fieldIDToName_CreateGroupRequest = map[int32]string{
	1: "Version",
	2: "GroupId",
	3: "GroupOwnerId",
}

var fieldIDToName_CreateGroupResponse = map[int32]string{
	1: "Success",
	2: "Version",
}

var fieldIDToName_AlterGroupMemberRequest = map[int32]string{
	1: "Version",
	2: "GroupId",
	3: "IsAdd",
	4: "MemberId",
	5: "RelationId",
}

var fieldIDToName_AlterGroupMemberResponse = map[int32]string{
	1: "Success",
	2: "Version",
}

var fieldIDToName_SubscribeGroupRequest = map[int32]string{
	1: "Version",
	2: "GroupId",
	3: "GwAdvertiseAddrPort",
	4: "SessionId",
}

var fieldIDToName_SubscribeGroupResponse = map[int32]string{
	1: "Code",
	2: "Version",
	3: "RelationId",
	4: "MaxSeqId",
}

var fieldIDToName_SendMessageRequest = map[int32]string{
	1: "Version",
	2: "GroupId",
	3: "MessageData",
	4: "CheckIdempotent",
	5: "IdempotentKey",
}

var fieldIDToName_SendMessageResponse = map[int32]string{
	1: "Code",
	2: "Version",
	3: "RelationId",
	4: "MaxSeqId",
	5: "MessageSeq",
}

var fieldIDToName_DisbandGroupRequest = map[int32]string{
	1: "Version",
	2: "GroupId",
}

var fieldIDToName_DisbandGroupResponse = map[int32]string{
	1: "Success",
	2: "Version",
}

var _ = base.File_base_base_proto
