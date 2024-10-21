// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package imchatevloop

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

func (x *CreateGroupResponse) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.CreatedAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UniversalGroupEvloopRequestReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UniversalGroupEvloopRequestReq[number], err)
}

func (x *UniversalGroupEvloopRequestReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UniversalGroupEvloopRequestReq) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *UniversalGroupEvloopRequestReq) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v evloopio.UniversalGroupEvloopInput
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Input = &v
	return offset, nil
}

func (x *UniversalGroupEvloopRequestResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_UniversalGroupEvloopRequestResp[number], err)
}

func (x *UniversalGroupEvloopRequestResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Success, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *UniversalGroupEvloopRequestResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Version, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *UniversalGroupEvloopRequestResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	var v evloopio.UniversalGroupEvloopOutput
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Output = &v
	return offset, nil
}

func (x *DoMigrateReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoMigrateReq[number], err)
}

func (x *DoMigrateReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DoMigrateResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoMigrateResp[number], err)
}

func (x *DoMigrateResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Ok, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *DoMigrateResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DoMigrateResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.OwnerId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DoMigrateResp) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CreatedAt, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DoMigrateResp) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.SeqId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DoMigrateResp) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	if x.Relations == nil {
		x.Relations = make(map[string]*DoMigrateResp_RelationInfo)
	}
	var key string
	var value *DoMigrateResp_RelationInfo
	offset, err = fastpb.ReadMapEntry(buf, _type,
		func(buf []byte, _type int8) (offset int, err error) {
			key, offset, err = fastpb.ReadString(buf, _type)
			return offset, err
		},
		func(buf []byte, _type int8) (offset int, err error) {
			var v DoMigrateResp_RelationInfo
			offset, err = fastpb.ReadMessage(buf, _type, &v)
			if err != nil {
				return offset, err
			}
			value = &v
			return offset, nil
		})
	if err != nil {
		return offset, err
	}
	x.Relations[key] = value
	return offset, nil
}

func (x *DoMigrateResp) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	if x.Subscribers == nil {
		x.Subscribers = make(map[string]*DoMigrateResp_UserSubscribeEntry)
	}
	var key string
	var value *DoMigrateResp_UserSubscribeEntry
	offset, err = fastpb.ReadMapEntry(buf, _type,
		func(buf []byte, _type int8) (offset int, err error) {
			key, offset, err = fastpb.ReadString(buf, _type)
			return offset, err
		},
		func(buf []byte, _type int8) (offset int, err error) {
			var v DoMigrateResp_UserSubscribeEntry
			offset, err = fastpb.ReadMessage(buf, _type, &v)
			if err != nil {
				return offset, err
			}
			value = &v
			return offset, nil
		})
	if err != nil {
		return offset, err
	}
	x.Subscribers[key] = value
	return offset, nil
}

func (x *MigrateDoneReq) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MigrateDoneReq[number], err)
}

func (x *MigrateDoneReq) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.GroupId, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *MigrateDoneResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_MigrateDoneResp[number], err)
}

func (x *MigrateDoneResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Ok, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *DoMigrateResp_RelationInfo) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoMigrateResp_RelationInfo[number], err)
}

func (x *DoMigrateResp_RelationInfo) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v base.RelationEntry
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Relation = &v
	return offset, nil
}

func (x *DoMigrateResp_UserSubscribeEntry) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoMigrateResp_UserSubscribeEntry[number], err)
}

func (x *DoMigrateResp_UserSubscribeEntry) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v DoMigrateResp_UserSubscribeEntry_SubscribeEntry
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Entries = append(x.Entries, &v)
	return offset, nil
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
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
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DoMigrateResp_UserSubscribeEntry_SubscribeEntry[number], err)
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.OnSubRelationVersion, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v base.SessionEntry
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Session = &v
	return offset, nil
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
	offset += x.fastWriteField3(buf[offset:])
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

func (x *CreateGroupResponse) fastWriteField3(buf []byte) (offset int) {
	if x.CreatedAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.GetCreatedAt())
	return offset
}

func (x *UniversalGroupEvloopRequestReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *UniversalGroupEvloopRequestReq) fastWriteField1(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetVersion())
	return offset
}

func (x *UniversalGroupEvloopRequestReq) fastWriteField2(buf []byte) (offset int) {
	if x.GroupId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetGroupId())
	return offset
}

func (x *UniversalGroupEvloopRequestReq) fastWriteField3(buf []byte) (offset int) {
	if x.Input == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetInput())
	return offset
}

func (x *UniversalGroupEvloopRequestResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *UniversalGroupEvloopRequestResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Success {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetSuccess())
	return offset
}

func (x *UniversalGroupEvloopRequestResp) fastWriteField2(buf []byte) (offset int) {
	if x.Version == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.GetVersion())
	return offset
}

func (x *UniversalGroupEvloopRequestResp) fastWriteField3(buf []byte) (offset int) {
	if x.Output == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 3, x.GetOutput())
	return offset
}

func (x *DoMigrateReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DoMigrateReq) fastWriteField1(buf []byte) (offset int) {
	if x.GroupId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetGroupId())
	return offset
}

func (x *DoMigrateResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	return offset
}

func (x *DoMigrateResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Ok {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetOk())
	return offset
}

func (x *DoMigrateResp) fastWriteField2(buf []byte) (offset int) {
	if x.GroupId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.GetGroupId())
	return offset
}

func (x *DoMigrateResp) fastWriteField3(buf []byte) (offset int) {
	if x.OwnerId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.GetOwnerId())
	return offset
}

func (x *DoMigrateResp) fastWriteField4(buf []byte) (offset int) {
	if x.CreatedAt == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.GetCreatedAt())
	return offset
}

func (x *DoMigrateResp) fastWriteField5(buf []byte) (offset int) {
	if x.SeqId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.GetSeqId())
	return offset
}

func (x *DoMigrateResp) fastWriteField6(buf []byte) (offset int) {
	if x.Relations == nil {
		return offset
	}
	for k, v := range x.GetRelations() {
		offset += fastpb.WriteMapEntry(buf[offset:], 6,
			func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
				offset := 0
				offset += fastpb.WriteString(buf[offset:], numTagOrKey, k)
				offset += fastpb.WriteMessage(buf[offset:], numIdxOrVal, v)
				return offset
			})
	}
	return offset
}

func (x *DoMigrateResp) fastWriteField7(buf []byte) (offset int) {
	if x.Subscribers == nil {
		return offset
	}
	for k, v := range x.GetSubscribers() {
		offset += fastpb.WriteMapEntry(buf[offset:], 7,
			func(buf []byte, numTagOrKey, numIdxOrVal int32) int {
				offset := 0
				offset += fastpb.WriteString(buf[offset:], numTagOrKey, k)
				offset += fastpb.WriteMessage(buf[offset:], numIdxOrVal, v)
				return offset
			})
	}
	return offset
}

func (x *MigrateDoneReq) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *MigrateDoneReq) fastWriteField1(buf []byte) (offset int) {
	if x.GroupId == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.GetGroupId())
	return offset
}

func (x *MigrateDoneResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *MigrateDoneResp) fastWriteField1(buf []byte) (offset int) {
	if !x.Ok {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 1, x.GetOk())
	return offset
}

func (x *DoMigrateResp_RelationInfo) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DoMigrateResp_RelationInfo) fastWriteField1(buf []byte) (offset int) {
	if x.Relation == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.GetRelation())
	return offset
}

func (x *DoMigrateResp_UserSubscribeEntry) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	return offset
}

func (x *DoMigrateResp_UserSubscribeEntry) fastWriteField1(buf []byte) (offset int) {
	if x.Entries == nil {
		return offset
	}
	for i := range x.GetEntries() {
		offset += fastpb.WriteMessage(buf[offset:], 1, x.GetEntries()[i])
	}
	return offset
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) fastWriteField1(buf []byte) (offset int) {
	if x.OnSubRelationVersion == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.GetOnSubRelationVersion())
	return offset
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) fastWriteField2(buf []byte) (offset int) {
	if x.Session == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.GetSession())
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
	n += x.sizeField3()
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

func (x *CreateGroupResponse) sizeField3() (n int) {
	if x.CreatedAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.GetCreatedAt())
	return n
}

func (x *UniversalGroupEvloopRequestReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *UniversalGroupEvloopRequestReq) sizeField1() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetVersion())
	return n
}

func (x *UniversalGroupEvloopRequestReq) sizeField2() (n int) {
	if x.GroupId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetGroupId())
	return n
}

func (x *UniversalGroupEvloopRequestReq) sizeField3() (n int) {
	if x.Input == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetInput())
	return n
}

func (x *UniversalGroupEvloopRequestResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *UniversalGroupEvloopRequestResp) sizeField1() (n int) {
	if !x.Success {
		return n
	}
	n += fastpb.SizeBool(1, x.GetSuccess())
	return n
}

func (x *UniversalGroupEvloopRequestResp) sizeField2() (n int) {
	if x.Version == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.GetVersion())
	return n
}

func (x *UniversalGroupEvloopRequestResp) sizeField3() (n int) {
	if x.Output == nil {
		return n
	}
	n += fastpb.SizeMessage(3, x.GetOutput())
	return n
}

func (x *DoMigrateReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DoMigrateReq) sizeField1() (n int) {
	if x.GroupId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetGroupId())
	return n
}

func (x *DoMigrateResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	return n
}

func (x *DoMigrateResp) sizeField1() (n int) {
	if !x.Ok {
		return n
	}
	n += fastpb.SizeBool(1, x.GetOk())
	return n
}

func (x *DoMigrateResp) sizeField2() (n int) {
	if x.GroupId == "" {
		return n
	}
	n += fastpb.SizeString(2, x.GetGroupId())
	return n
}

func (x *DoMigrateResp) sizeField3() (n int) {
	if x.OwnerId == "" {
		return n
	}
	n += fastpb.SizeString(3, x.GetOwnerId())
	return n
}

func (x *DoMigrateResp) sizeField4() (n int) {
	if x.CreatedAt == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.GetCreatedAt())
	return n
}

func (x *DoMigrateResp) sizeField5() (n int) {
	if x.SeqId == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.GetSeqId())
	return n
}

func (x *DoMigrateResp) sizeField6() (n int) {
	if x.Relations == nil {
		return n
	}
	for k, v := range x.GetRelations() {
		n += fastpb.SizeMapEntry(6,
			func(numTagOrKey, numIdxOrVal int32) int {
				n := 0
				n += fastpb.SizeString(numTagOrKey, k)
				n += fastpb.SizeMessage(numIdxOrVal, v)
				return n
			})
	}
	return n
}

func (x *DoMigrateResp) sizeField7() (n int) {
	if x.Subscribers == nil {
		return n
	}
	for k, v := range x.GetSubscribers() {
		n += fastpb.SizeMapEntry(7,
			func(numTagOrKey, numIdxOrVal int32) int {
				n := 0
				n += fastpb.SizeString(numTagOrKey, k)
				n += fastpb.SizeMessage(numIdxOrVal, v)
				return n
			})
	}
	return n
}

func (x *MigrateDoneReq) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *MigrateDoneReq) sizeField1() (n int) {
	if x.GroupId == "" {
		return n
	}
	n += fastpb.SizeString(1, x.GetGroupId())
	return n
}

func (x *MigrateDoneResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *MigrateDoneResp) sizeField1() (n int) {
	if !x.Ok {
		return n
	}
	n += fastpb.SizeBool(1, x.GetOk())
	return n
}

func (x *DoMigrateResp_RelationInfo) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DoMigrateResp_RelationInfo) sizeField1() (n int) {
	if x.Relation == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.GetRelation())
	return n
}

func (x *DoMigrateResp_UserSubscribeEntry) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	return n
}

func (x *DoMigrateResp_UserSubscribeEntry) sizeField1() (n int) {
	if x.Entries == nil {
		return n
	}
	for i := range x.GetEntries() {
		n += fastpb.SizeMessage(1, x.GetEntries()[i])
	}
	return n
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) sizeField1() (n int) {
	if x.OnSubRelationVersion == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.GetOnSubRelationVersion())
	return n
}

func (x *DoMigrateResp_UserSubscribeEntry_SubscribeEntry) sizeField2() (n int) {
	if x.Session == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.GetSession())
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
	3: "CreatedAt",
}

var fieldIDToName_UniversalGroupEvloopRequestReq = map[int32]string{
	1: "Version",
	2: "GroupId",
	3: "Input",
}

var fieldIDToName_UniversalGroupEvloopRequestResp = map[int32]string{
	1: "Success",
	2: "Version",
	3: "Output",
}

var fieldIDToName_DoMigrateReq = map[int32]string{
	1: "GroupId",
}

var fieldIDToName_DoMigrateResp = map[int32]string{
	1: "Ok",
	2: "GroupId",
	3: "OwnerId",
	4: "CreatedAt",
	5: "SeqId",
	6: "Relations",
	7: "Subscribers",
}

var fieldIDToName_MigrateDoneReq = map[int32]string{
	1: "GroupId",
}

var fieldIDToName_MigrateDoneResp = map[int32]string{
	1: "Ok",
}

var fieldIDToName_DoMigrateResp_RelationInfo = map[int32]string{
	1: "Relation",
}

var fieldIDToName_DoMigrateResp_UserSubscribeEntry = map[int32]string{
	1: "Entries",
}

var fieldIDToName_DoMigrateResp_UserSubscribeEntry_SubscribeEntry = map[int32]string{
	1: "OnSubRelationVersion",
	2: "Session",
}

var _ = base.File_base_base_proto
var _ = evloopio.File_base_evloopio_proto
