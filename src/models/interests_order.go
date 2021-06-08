package models

import (
	"fmt"
	"web-project-model/src/models/db"
)

func InsertInterestsOrder(inOrder *db.InterestsOrder) error {
	if err := db.DB.Debug().Create(inOrder).Error; err != nil {
		return fmt.Errorf("[DB] create interests order failed: %s", err.Error())
	}
	return nil
}
