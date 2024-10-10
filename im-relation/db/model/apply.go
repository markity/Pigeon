package model

type ApplyModel struct {
	// TODO: 这里暂时用1, 2, 3...表示一个群, 后续需要改进
	Id           int64  `gorm:"column:id,primaryKey"`
	GroupId      int64  `gorm:"column:group_id"`
	OwnerId      string `gorm:"column:owner_id"`
	ApplyCounter int64  `gorm:"column:apply_counter"`
	CreatedAt    int64  `gorm:"column:created_at"`
	UpdatedAt    int64  `gorm:"column:updated_at"`
}

func (ApplyModel) TableName() string {
	return "apply"
}
