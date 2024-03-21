package service

import (
	"fmt"
	"k8s-log-translation/app/database"
	"k8s-log-translation/app/model"
)

var UserFields = []string{"id", "account", "email"}

func SelectOneUsers(id int64) (*model.User, error) {
	userOne:=&model.User{}
	err := database.SqlSession.Select(UserFields).Where("id=?", id).First(&userOne).Error
	if err != nil {
		return nil, err
	} else {
		return userOne, nil
	}
}

func RegisterOneUser(account string, password string, email string) error {
	if !CheckOneUser(account) {
		return fmt.Errorf("User exists.")
	}
	user := model.User{
		Account: account,
		Password: password,
		Email: email,
	}
	insertErr := database.SqlSession.Model(&model.User{}).Create(&user).Error
	return insertErr
}

func CheckOneUser(account string) bool {
	result := false
	var user model.User

	dbResult := database.SqlSession.Where("account = ?", account).Find(&user)
	if dbResult.Error != nil {
		fmt.Printf("Get User Info Failed:%v\n", dbResult.Error)
	} else {
		result = true
	}
	return result
}