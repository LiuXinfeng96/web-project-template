package models

import (
	"fmt"
	"web-project-model/src/models/db"

	"gorm.io/gorm"
)

func InsertInterestsOrder(inOrder *db.InterestsOrder) error {
	if err := db.DB.Debug().Create(inOrder).Error; err != nil {
		return fmt.Errorf("[DB] create interests order failed: %s", err.Error())
	}
	return nil
}

func MulInsertInterestsOrders(inOrders []*db.InterestsOrder) error {
	len := len(inOrders)
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(inOrders, len).Error; err != nil {
			return fmt.Errorf("[DB] create multiple interests order failed: %s", err.Error())
		}
		return nil
	})
	return err
}
