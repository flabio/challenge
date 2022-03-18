package utilities

type MessageRequired struct {
	Name string
}

func (msg MessageRequired) RequiredId() string {
	return "The id is required"
}

func (msg MessageRequired) RequiredName() string {
	return "The name is required"
}

func (msg MessageRequired) RequiredCriticidad() string {
	return "The criticidad is required"
}
func (msg MessageRequired) RequiredOwner() string {
	return "The owner is required"
}
