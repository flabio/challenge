package dto

type ShippingDetailDTO struct {
	Id     uint   `json:"id" form:"id" `
	Weight string `json:"weight" form:"weight" binding:"required"`
	Width  string `json:"width" form:"width" binding:"required"`
	Height string `json:"height" form:"height" binding:"required"`
	Depth  int64  `json:"depth" form:"depth" binding:"required"`
}
