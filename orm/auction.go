package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type Auction struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	ApproveDate time.Time `json:"approve_date" query:"approve_date" gorm:"type:time;"`
	ApproveBy   string    `json:"approve_by" query:"approve_by" gorm:"type:varchar(50);"`
	VehicleNo   string    `json:"vehicle_no" query:"vehicle_no" gorm:"type:varchar(20);"`
	AuctionNo   string    `json:"auction_no" query:"auction_no" gorm:"type:varchar(20);"`
	IsApprove   bool      `json:"is_approve" query:"is_approve" gorm:"type:bool;"`
}
