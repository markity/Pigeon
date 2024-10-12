package db

import (
	"pigeon/im-relation/db/model"

	"gorm.io/gorm"
)

func GetAllRelationsByUsername(txn *gorm.DB, username string) ([]*model.RelationModel, error) {
	var relations []*model.RelationModel
	err := txn.Model(&model.RelationModel{}).Where("owner_id = ?", username).Find(&relations).Error
	if err != nil {
		return nil, err
	}

	return relations, nil
}
