package models

import (
	"fmt"
	"testing"
	"web-project-model/src/models/db"
	"web-project-model/src/services"
	"web-project-model/src/utils"
)

func ModelsInit(t *testing.T) {
	utils.SetConfig(utils.GetConfigEnv())
	db.DBInit()
}

func TestInsertInterestsOrder(t *testing.T) {
	ModelsInit(t)

	inOrder := &db.InterestsOrder{
		Id:                 services.Getuuid(),
		OrderId:            "lxf",
		SuborderId:         "TEST",
		HashSuborderId:     "aaabbbccc",
		ProvinceId:         "test",
		OrderStatus:        "pc",
		StatusDesc:         "more",
		ProvinceRelationId: "more",
		Context:            "000000001111111",
	}
	err := InsertInterestsOrder(inOrder)
	if err != nil {
		fmt.Errorf("[DB] insert interests order error: %s", err.Error())
	}
	var inOrders []*db.InterestsOrder
	inOrder1 := &db.InterestsOrder{
		Id:                 services.Getuuid(),
		OrderId:            "lxf",
		SuborderId:         "MULTEST",
		HashSuborderId:     "aaabbbccc",
		ProvinceId:         "test",
		OrderStatus:        "pc",
		StatusDesc:         "more",
		ProvinceRelationId: "more",
		Context:            "000000001111111",
	}
	inOrder2 := &db.InterestsOrder{
		Id:                 services.Getuuid(),
		OrderId:            "lxf",
		SuborderId:         "MULTEST",
		HashSuborderId:     "aaabbbccc",
		ProvinceId:         "test",
		OrderStatus:        "pc",
		StatusDesc:         "more",
		ProvinceRelationId: "more",
		Context:            "000000001111111",
	}
	inOrders = append(inOrders, inOrder1)
	inOrders = append(inOrders, inOrder2)
	err = MulInsertInterestsOrders(inOrders)
	if err != nil {
		fmt.Errorf("[DB] insert interests order error: %s", err.Error())
	}
}
