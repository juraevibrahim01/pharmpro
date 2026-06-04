package models

import "errors"

type User_create_res struct {
	Status      string `json:"status"`
	Description string `json:"description"`
	Error       string `json:"error"`
}

type User_create_req struct {
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
	Password    string `json:"password"`
}

var User_err_exists_user = errors.New("Такой пользователь сушествует")
