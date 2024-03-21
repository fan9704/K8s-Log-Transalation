package controller

import (
	"k8s-log-translation/app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/i18n/gi18n"
)

type UsersController struct {}

func NewUsersController() UsersController {
	return UsersController{}
}

func QueryUsersController() UsersController {
	return UsersController{}
}

type Register struct {
	Account string `json:"account" binding:"required" example:"account"`
	Password string `json:"password" binding:"required" example:"password"`
	Email string `json:"email" binding:"required" example:"test123@gmail.com"`
}
// CreateUser @Summary
// @Tags user
// @version 1.0
// @produce application/json
// @param language header string true "language"
// @param register body Register true "register"
// @Success 200 string successful return value
// @Router /v1/users [post]
func (u UsersController) CreateUser (c *gin.Context){

	t := gi18n.New()
	lan := c.Request.Header.Get("language")
	if lan == "" {
		lan = "en"
	}
	t.SetLanguage(lan)

	var form Register
	bindErr := c.BindJSON(&form)

	if bindErr == nil {
		err := service.RegisterOneUser(form.Account, form.Password, form.Email)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 1,
				"msg": "success Register",
				"data": nil,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": -1,
				"msg": "Register Failed" + err.Error(),
				"data": nil,
			})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg": "Failed to parse register data" + bindErr.Error(),
			"data": nil,
		})
	}
}
// GetUser GetUser @Summary
// @Tags user
// @version 1.0
// @produce application/json
// @param id path int true "id" default(1)
// @Success 200 string string successful return data
// @Router /v1/users/{id} [get]
func (u UsersController) GetUser (c *gin.Context){
	id := c.Params.ByName("id")
	userId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": -1,
			"msg": "Failed to parse params" + err.Error(),
			"data": nil,
		})
	}

	userOne, err := service.SelectOneUsers(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": -1,
			"msg": "User not found" + err.Error(),
			"data": nil,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":  "Successfully get user data",
			"user": &userOne,
		})
	}
}