package response

type CommandResponseDTO struct{
	message string `json:"message" example:"success"`
	status bool `json:"status" example:"true"`
	
}