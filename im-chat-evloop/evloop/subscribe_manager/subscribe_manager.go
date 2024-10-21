package subscribemanager

import (
	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imchatevloop"

	"github.com/golang/protobuf/proto"
)

type Sub struct {
	SubRelationVersion int64
	Entry              *base.SessionEntry
}

// 非线程安全的订阅管理器, 只能在eventloop线程中使用
type SubscribeManager struct {
	// map<userId, map<session_id, session_entry>>
	subscribers map[string](map[string]*Sub)
}

func NewSubscribeManager() *SubscribeManager {
	return &SubscribeManager{subscribers: make(map[string](map[string]*Sub))}
}

func NewSubscrbieManagerFromMigrage(resp *imchatevloop.DoMigrateResp) *SubscribeManager {
	subscribes := make(map[string](map[string]*Sub))
	for k, v := range resp.Subscribers {
		subscribes[k] = make(map[string]*Sub)
		for _, sub := range v.Entries {
			subscribes[k][sub.Session.SessionId] = &Sub{
				SubRelationVersion: sub.OnSubRelationVersion,
				Entry:              sub.Session,
			}
		}
	}
	return &SubscribeManager{
		subscribers: subscribes,
	}
}

func (sm *SubscribeManager) GetSubscibers() map[string](map[string]*Sub) {
	return sm.subscribers
}

func (sm *SubscribeManager) SessionSub(userId string, sessionId string, relationVersion int64, sessionEntry *base.SessionEntry) {
	userSubs := sm.subscribers[userId]
	if userSubs == nil {
		userSubs = make(map[string]*Sub)
	}
	userSubs[sessionId] = &Sub{
		SubRelationVersion: relationVersion,
		Entry:              sessionEntry,
	}
	sm.subscribers[userId] = userSubs
}

func (sm *SubscribeManager) SessionUnsub(userId string, sessionId string) {
	userSubs := sm.subscribers[userId]
	if userSubs != nil {
		delete(userSubs, sessionId)
	}
}

func (sm *SubscribeManager) RemoveOldSubs(useId string, relationVersion int64) {
	userSubs := sm.subscribers[useId]
	for sessionId, sub := range userSubs {
		if relationVersion > sub.SubRelationVersion {
			delete(userSubs, sessionId)
		}
	}
}

func (sm *SubscribeManager) SnapshotAllSubs() []*base.SessionEntry {
	var allSubs []*base.SessionEntry
	for _, userSubs := range sm.subscribers {
		for _, sessionEntry := range userSubs {
			allSubs = append(allSubs, proto.Clone(sessionEntry.Entry).(*base.SessionEntry))
		}
	}
	return allSubs
}
