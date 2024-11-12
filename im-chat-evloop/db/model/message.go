package model

type MessageModel struct {
	Id        int64  `gorm:"column:id;primaryKey"`
	OwnerId   string `gorm:"column:owner_id;type:varchar(64)"`
	GroupId   string `gorm:"column:group_id;type:varchar(32)"`
	SeqId     int64  `gorm:"column:seq_id"`
	Data      string `gorm:"column:data"`
	CreatedAt int64  `gorm:"created_at"`
}

func (MessageModel) TableName() string {
	return "message"
}
