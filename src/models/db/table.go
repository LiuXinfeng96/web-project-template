package db

type InterestsOrder struct {
	Id                 string `gorm:"primaryKey"`
	OrderId            string `gorm:"index:orderid_suborderid_index,priority:1"`
	SuborderId         string `gorm:"index:orderid_suborderid_index,priority:2"`
	HashSuborderId     string `gorm:"index:hash_suborderid_index"`
	ProvinceId         string
	OrderStatus        string
	StatusDesc         string
	ProvinceRelationId string
	CreateAt           int
	UpdateAt           int    `gorm:"index:update_at_index"`
	Context            string `gorm:"type:longtext"`
}

func (i *InterestsOrder) TableName() string {
	return "interests_orders"
}
