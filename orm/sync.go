package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type Sync struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	SyncDate time.Time `json:"sync_date" query:"sync_date" gorm:"type:time;"`
	SyncBy   string    `json:"sync_by" query:"sync_by" gorm:"type:varchar(50);"`
	SyncName string    `json:"sync_name" query:"sync_name" gorm:"type:varchar(50);"`
	SyncType string    `json:"sync_type" query:"sync_type" gorm:"type:varchar(5);"`
}
