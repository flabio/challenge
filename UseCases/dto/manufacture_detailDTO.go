package dto

type ManufactureDetailDTO struct {
	Id          uint   `json:"id" form:"id" `
	ModelNumber string `json:"model_number" form:"model_number" binding:"required"`
	ReleaseDate string `json:"release_date" form:"release_date" binding:"required"`
}
