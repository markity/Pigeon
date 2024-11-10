package model

import "pigeon/kitex_gen/service/base"

type ApplyModel struct {
	Id      int64  `gorm:"column:id;primaryKey"`
	OwnerId string `gorm:"column:owner_id;type:varchar(256);uniqueIndex:idx_ownerid_groupid"`
	GroupId string `gorm:"column:group_id;type:varchar(32);uniqueIndex:idx_ownerid_groupid"`
	// 申请自增id, 一个用户最多一条applyModel记录, 多次申请则递增
	ApplyVersion int64 `gorm:"column:apply_version"`
	// 申请原因, 同qq的入群理由
	ApplyMsg  string `gorm:"column:apply_msg"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
	// 这个字段用来推送给群主apply, 其实是group的子属性
	GroupOwnerId string           `gorm:"column:group_owner_id;type:varchar(256);index:idx_group_owner_id"`
	Status       base.ApplyStatus `gorm:"column:status"`
}

func (ApplyModel) TableName() string {
	return "apply"
}
