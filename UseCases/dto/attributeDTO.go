package dto

type AttributeDTO struct {
	Id    uint   `json:"id" form:"id" `
	Brand string `json:"brand" form:"brand" binding:"required"`
	Color string `json:"color" form:"color" binding:"required"`
}
