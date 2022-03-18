package entitys

type Administrador struct {
	Id         uint   `gorm:"primary_key:auto_increment" json:"id"`
	Name       string `gorm:"type:varchar(255)" json:"name"`
	Criticidad string `gorm:"type:varchar(255)" json:"criticidad"`
	Owner      string `gorm:"type:varchar(255)" json:"owner"`
	Token      string `gorm:"-" json:"token,omitempty"`
}
