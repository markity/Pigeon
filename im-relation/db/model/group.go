package model

type GroupModel struct {
	// TODO: 这里暂时用1, 2, 3...表示一个群, 后续需要改进
	Id        int64  `gorm:"column:id;primaryKey"`
	GroupId   string `gorm:"column:group_id;type:varchar(32);unique"`
	OwnerId   string `gorm:"column:owner_id"`
	CreatedAt int64  `gorm:"column:created_at"`
	// Disbaned   bool   `gorm:"column:disbaned"`
	// DisbanedAt int64  `gorm:"column:disbaned_at"`
}

func (GroupModel) TableName() string {
	return "grp"
}
