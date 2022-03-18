package dto

type AdministradorDTO struct {
	Id         uint   `json:"id" form:"id" `
	Name       string `json:"name" form:"name" binding:"required"`
	Criticidad string `json:"criticidad" form:"criticidad" binding:"required"`
	Owner      string `json:"owner" form:"owner" binding:"required"`
}

type LoginDTO struct {
	Name  string `json:"name" form:"name" binding:"required"`
	Owner string `json:"owner" form:"owner" binding:"required"`
}
