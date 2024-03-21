package service

import (
	"bytes"
	"fmt"
	"k8s-log-translation/app/dto"
	"os/exec"
)

func K8sGPTService() string{
	var response = ""
	cmd := exec.Command("k8sgpt", "analyze")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr 
	err := cmd.Run()
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	response = outStr

	var commandResponse = dto.CommandResponseDTO{}
	commandResponse.Message = response
	commandResponse.Status = true
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	if err != nil {
		commandResponse.Message = err.Error()
		commandResponse.Status = false
	}
	
	return response
}