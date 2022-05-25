package entitys

type Pricing struct {
	Id    uint   `gorm:"primary_key:auto_increment" json:"id"`
	Price string `gorm:"type:varchar(255)" json:"price"`
}
