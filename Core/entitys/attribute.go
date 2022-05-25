package entitys

type Attribute struct {
	Id    uint   `gorm:"primary_key:auto_increment" json:"id"`
	Brand string `gorm:"type:varchar(255)" json:"brand"`
	Color string `gorm:"type:varchar(255)" json:"color"`
}
