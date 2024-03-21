package controller

import (
	"k8s-log-translation/app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)
type CommandController struct {}

func ExecuteCommandController() CommandController {
	return CommandController{}
}

// Execute K8sGPT Command@Summary
// @Tags command
// @version 1.0
// @produce application/json
// @Success 200 string successful return value
// @Router /v1/command/k8sgpt [get]
func (controller CommandController) K8sGPTCommand(c *gin.Context){

	var response = service.K8sGPTService()
	if response == ""{
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": -1,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"status": 1,
			"response": response,
		})

	}
}