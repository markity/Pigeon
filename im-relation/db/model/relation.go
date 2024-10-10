package model

type RelationStatus int

const (
	RelationStatusUnUsed     RelationStatus = iota
	ReleationStatusOwner     RelationStatus = 1
	RelationStatusInGroup    RelationStatus = 2
	RelationStatusNotInGroup RelationStatus = 3
)

type RelationModel struct {
	// TODO: 这里暂时用1, 2, 3...表示一个群, 后续需要改进
	Id        int64          `gorm:"column:id,primaryKey"`
	OwnerId   string         `gorm:"column:owner_id"`
	GroupId   int64          `gorm:"column:group_id"`
	Status    RelationStatus `gorm:"column:status"`
	CreatedAt int64          `gorm:"column:created_at"`
}

func (RelationModel) TableName() string {
	return "relation"
}
