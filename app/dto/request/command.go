package request

type CommandRequestDTO struct {
	command string `json:"command" example:"iptables -L -n -v"`
}