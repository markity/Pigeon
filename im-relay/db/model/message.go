package model

type MessageModel struct {
	Id      int64  `gorm:"column:id;primaryKey"`
	OwnerId string `gorm:"column:owner_id;type:varchar(64)"`
	GroupId string `gorm:"column:group_id;type:varchar(32)"`
	SeqId   int64  `gorm:"column:seq_id"`
	MsgId   string `gorm:"column:msg_id;type:varchar(64)"`
	// username-客户端自定义生成
	IdempotentKey string `gorm:"column:idempotent_key;type:varchar(64);index"`
	Data          string `gorm:"column:data"`
	CreatedAt     int64  `gorm:"created_at"`
}

func (MessageModel) TableName() string {
	return "message"
}
