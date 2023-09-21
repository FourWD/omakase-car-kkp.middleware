package orm

import "github.com/FourWD/middleware/orm"

type CarType struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel
	TypeID   string `json:"typeid" query:"typeid" gorm:"type:varchar(5);"`
	TypeName string `json:"typename" query:"typename" gorm:"type:varchar(20);"`
	//eiei
}
