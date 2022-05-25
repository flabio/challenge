package dto

type ProductDTO struct {
	Id                  uint   `json:"id" form:"id" `
	Sku                 string `json:"sku" form:"sku" binding:"required"`
	Name                string `json:"name" form:"name" binding:"required"`
	Description         string `json:"description" form:"description" binding:"required"`
	Department          string `json:"department" form:"department" binding:"required"`
	Quantity            int64  `json:"quantity" form:"quantity" binding:"required"`
	ShippingDetailId    uint   `json:"shipping_detail_id" form:"shipping_detail_id" binding:"required"`
	ManufactureDetailId uint   `json:"manufacture_detail_id" form:"manufacture_detail_id" binding:"required"`
	AttributeId         uint   `json:"attribute_id" form:"attribute_id" binding:"required"`
}
