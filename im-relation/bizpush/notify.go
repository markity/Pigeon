package bizpush

import (
	"pigeon/kitex_gen/service/base"
)

type ApplyGroupNotifyInput struct {
	// 群拥有者
	OwnerId string

	// 某个人, 关于某个群的申请信息+版本号
	Username     string
	GroupId      string
	ApplyMsg     string
	ApplyVersion int64
	ApplyAt      int64
}

func (bp *BizPusher) ApplyGroupNotify(input *ApplyGroupNotifyInput) {
	m := map[string]interface{}{
		"user_id":       input.Username,
		"group_id":      input.GroupId,
		"apply_version": input.ApplyVersion,
		"apply_msg":     input.ApplyMsg,
		"apply_at":      input.ApplyAt,
	}

	bp.pushMan.PushToUserByMap(input.OwnerId, "push-apply-notify", "", m)
}

type HandleApplyNotifyInput struct {
	OwnerId string

	Username     string
	GroupId      string
	ApplyVersion int64
	ApplyMsg     string
	ApplyStatus  base.ApplyStatus
	ApplyAt      int64
	HandleAt     int64
}

// 用户处理了apply, 需要把新版本的apply推送给group owner
func (bp *BizPusher) HandleApplyNotify(input *HandleApplyNotifyInput) {
	m := map[string]interface{}{
		"user_id":       input.Username,
		"group_id":      input.GroupId,
		"apply_version": input.ApplyVersion,
		"apply_msg":     input.ApplyMsg,
		"apply_status":  input.ApplyStatus,
		"apply_at":      input.ApplyAt,
		"handle_at":     input.HandleAt,
	}
	bp.pushMan.PushToUserByMap(input.OwnerId, "push-handle-apply-notify", "", m)
}

// 比如a成功加入了群, 那么a加入群的消息将下发给a的全部设备
type RelationChangeNotifyInput struct {
	Username   string
	GroupId    string
	Version    int64
	Status     base.RelationStatus
	UpdatedAt  int64
	ChangeType base.RelationChangeType
}

func (bp *BizPusher) RelationChangeNotify(input *RelationChangeNotifyInput) {
	m := map[string]interface{}{
		"username":             input.Username,
		"group_id":             input.GroupId,
		"relation_version":     input.Version,
		"relation_status":      input.Status,
		"relation_change_type": input.ChangeType,
		"updated_at":           input.UpdatedAt,
	}

	bp.pushMan.PushToUserByMap(input.Username, "push-relation-change-notify", "", m)
}

// TODO: 暂时不支持退出群, 先做好主要功能
// type QuitGroupRespInput struct {
// 	EchoCode        string
// 	Code            imrelation.QuitGroupResp_QuitGroupRespCode
// 	RelationVersion int64
// }

// func QuitGroupResp(session *base.SessionEntry, input *QuitGroupRespInput) {
// 	cli := api.NewGatewayClientFromAdAddr(session.GwAdvertiseAddrport)
// 	for {
// 		_, err := cli.PushMessage(context.Background(), &imgateway.PushMessageReq{
// 			SessionId: session.SessionId,
// 			EchoCode:  input.EchoCode,
// 			PushType:  "push-quit-group-resp",
// 			Data: mustMarshal(map[string]interface{}{
// 				"code":             input.Code,
// 				"relation_version": input.RelationVersion,
// 			}),
// 		})
// 		if err != nil {
// 			log.Printf("push QuitGroupResp: %v\n", err)
// 			time.Sleep(time.Millisecond * 50)
// 			continue
// 		}
// 		break
// 	}
// }
