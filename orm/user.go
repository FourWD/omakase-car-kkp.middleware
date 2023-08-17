package orm

import "github.com/FourWD/middleware/orm"

type User struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	orm.GormModel

	UserTypeID string `json:"user_type_id" query:"user_type_id" gorm:"type:varchar(36);"`

	Username     string `json:"username" query:"username" gorm:"type:varchar(20);"`
	Password     string `json:"password" query:"password" gorm:"type:varchar(20);"`
	Firstname    string `json:"firstname" query:"firstname" gorm:"type:varchar(100);"`
	Lastname     string `json:"lastname" query:"lastname" orm:"type:varchar(100);"`
	FileAvatarID string `json:"file_avartar_id" query:"file_avartar_id" gorm:"type:varchar(36)"`
	Mobile       string `json:"mobile" query:"mobile" gorm:"type:varchar(20);"`
	Email        string `json:"email" query:"email" gorm:"type:varchar(20);"`
	//eiei
}
