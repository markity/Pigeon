package model

import "pigeon/kitex_gen/service/base"

type ApplyModel struct {
	// TODO: 这里暂时用1, 2, 3...表示一个群, 后续需要改进
	Id      int64  `gorm:"column:id;primaryKey"`
	OwnerId string `gorm:"column:owner_id;type:varchar(256);uniqueIndex:idx_ownerid_groupid"`
	GroupId int64  `gorm:"column:group_id;uniqueIndex:idx_ownerid_groupid"`
	// 这里给unique, 1062错误, TODO: 后续需要改进
	// 申请自增id, 一个用户最多一条applyModel记录, 多次申请则递增
	ApplyVersion int64 `gorm:"column:apply_version"`
	// 申请原因, 同qq的入群理由
	ApplyMsg     string           `gorm:"column:apply_msg"`
	CreatedAt    int64            `gorm:"column:created_at"`
	UpdatedAt    int64            `gorm:"column:updated_at"`
	GroupOwnerId string           `gorm:"column:group_owner_id;type:varchar(256);index:idx_group_owner_id"`
	Status       base.ApplyStatus `gorm:"column:status"`
}

func (ApplyModel) TableName() string {
	return "apply"
}
