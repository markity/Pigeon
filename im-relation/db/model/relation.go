package model

import (
	"pigeon/kitex_gen/service/imrelation"
)

type RelationModel struct {
	// TODO: 这里暂时用1, 2, 3...表示一个群, 后续需要改进
	Id int64 `gorm:"column:id,primaryKey"`
	// ownerId和groupId 建立联合唯一索引
	OwnerId string                                  `gorm:"column:owner_id;unique_index:idx_ownerid_groupid"`
	GroupId int64                                   `gorm:"column:group_id,unique_index:idx_ownerid_groupid"`
	Status  imrelation.RelationEntry_RelationStatus `gorm:"column:status"`
	// relation_counter 用于记录关系version, 越大的就是更新的关系, 一个用户关于一个群只有一条RelationModel
	RelationCounter int64 `gorm:"column:relation_counter"`
	CreatedAt       int64 `gorm:"column:created_at"`
	UpdatedAt       int64 `gorm:"column:updated_at"`
}

func (RelationModel) TableName() string {
	return "relation"
}
