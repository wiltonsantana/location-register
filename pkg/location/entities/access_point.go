package entities

// AccessPoint represents the user domain model
type AccessPoint struct {
	Name string  `json:"name" validate:"required"`
	Mac  string  `json:"mac" validate:"required"`
	X    float32 `json:"x" validate:"required"`
	Y    float32 `json:"y" validate:"required"`
}
