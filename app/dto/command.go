package dto

type CommandResponseDTO struct {
	Message string `json:"message" example:"success"`
	Status  bool   `json:"status" example:"true"`
}