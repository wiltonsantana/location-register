package entities

// Location represents the user domain model
type Location struct {
	Width       float32       `json:"width" validate:"required"`
	Length      float32       `json:"length" validate:"required"`
	AccessPoint []AccessPoint `json:"accessPoint,omitempty"`
}
