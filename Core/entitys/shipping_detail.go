package entitys

type ShippingDetail struct {
	Id     uint   `gorm:"primary_key:auto_increment" json:"id"`
	Weight string `gorm:"type:varchar(255)" json:"weight"`
	Width  string `gorm:"type:varchar(255)" json:"width"`
	Height string `gorm:"type:varchar(255)" json:"height"`
	Depth  int64  `gorm:"type:int(20)" json:"depth"`
}
