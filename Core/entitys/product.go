package entitys

type Product struct {
	Id                  uint              `gorm:"primary_key:auto_increment" json:"id"`
	Sku                 string            `gorm:"type:varchar(255)" json:"sku"`
	Name                string            `gorm:"type:varchar(255)" json:"name"`
	Description         string            `gorm:"type:varchar(255)" json:"description"`
	Department          string            `gorm:"type:varchar(255)" json:"department"`
	Quantity            int64             `gorm:"type:int(50)" json:"quantity"`
	ShippingDetailId    uint              `gorm:"NULL" json:"shipping_detailId"`
	ShippingDetail      ShippingDetail    `gorm:"foreignkey:ShippingDetailId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"shipping_detail"`
	ManufactureDetailId uint              `gorm:"NULL" json:"manufacture_detailId"`
	ManufactureDetail   ManufactureDetail `gorm:"foreignkey:ManufactureDetailId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"manufacture_detail"`
	AttributeId         uint              `gorm:"NULL" json:"attributeId"`
	Attribute           Attribute         `gorm:"foreignkey:AttributeId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"attribute"`
}
