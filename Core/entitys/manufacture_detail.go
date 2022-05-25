package entitys

type ManufactureDetail struct {
	Id          uint   `gorm:"primary_key:auto_increment" json:"id"`
	ModelNumber string `gorm:"type:varchar(255)" json:"    model_number"`
	ReleaseDate string `gorm:"type:varchar(255)" json:"    release_date"`
}
