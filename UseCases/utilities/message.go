package utilities

type MessageRequired struct {
	Name string
}

func (msg MessageRequired) RequiredId() string {
	return "El id es requerido"
}

func (msg MessageRequired) RequiredName() string {
	return "El nombre es requerido"
}

func (msg MessageRequired) RequiredCriticidad() string {
	return "La criticidad es requerido"
}
func (msg MessageRequired) RequiredOwner() string {
	return "El owner es requerido"
}
