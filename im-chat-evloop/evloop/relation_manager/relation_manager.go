package relationmanager

import (
	"errors"
	"time"

	"pigeon/kitex_gen/service/base"
	"pigeon/kitex_gen/service/imchatevloop"
)

// 管理relation, 非线程安全, 供evevloop使用
type RelationManager struct {
	relations map[string]*base.RelationEntry
}

func (rm *RelationManager) GetRelations() map[string]*base.RelationEntry {
	return rm.relations
}

func NewRelationManager(ownerRelation *base.RelationEntry) *RelationManager {
	relations := make(map[string]*base.RelationEntry)
	relations[ownerRelation.UserId] = ownerRelation
	return &RelationManager{
		relations: relations,
	}
}

func NewRelationManagerFromMigrage(resp *imchatevloop.DoMigrateResp) *RelationManager {
	relations := make(map[string]*base.RelationEntry)
	for k, v := range resp.Relations {
		relations[k] = v.Relation
	}
	return &RelationManager{
		relations: relations,
	}
}

// 每次更新relation, 都得把老的subscriber删除
func (rm *RelationManager) UpdateRelation(relation *base.RelationEntry) error {
	origin := rm.relations[relation.UserId]
	// relation变化不符合逻辑, 此处直接返回错误, 让调用方得知风险
	if origin != nil && (origin.RelationVersion > relation.RelationVersion ||
		abs(origin.RelationVersion, relation.RelationVersion) > 1) {
		return errors.New("relation version change not match")
	}
	rm.relations[relation.UserId] = relation
	relation.ChangeAt = time.Now().UnixMilli()
	return nil
}

func (rm *RelationManager) CanSubscribe(userId string) (subRelationVersion int64, canSubScribe bool) {
	relation := rm.relations[userId]
	ok := relation != nil && relation.Status != base.RelationStatus_RELATION_STATUS_NOT_IN_GROUP
	if !ok {
		return 0, false
	}

	version := relation.RelationVersion
	return version, true
}

func abs(a int64, b int64) int64 {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}
