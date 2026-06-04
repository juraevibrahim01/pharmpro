package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/juraevibrahim01/pharmpro/internal/models"
	"github.com/juraevibrahim01/pharmpro/internal/service"
)

type User_handler struct {
	service *service.User_service
}

func User_new_handler(service *service.User_service) *User_handler {
	return &User_handler{service: service}
}

func (u *User_handler) User_create(w http.ResponseWriter, r *http.Request) {

	// Response | Request
	var req models.User_create_req
	var res models.User_create_res

	// парсим данные
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка сервера", 500)
		return
	}

	// Проверка: Все поля возвращются
	if req.AccessToken == "" || req.Password == "" || req.Name == "" {
		res.Status = "error"
		res.Error = "Поля не валидны"
		w.WriteHeader(400)
		// response
		json.NewEncoder(w).Encode(res)
		return
	}

	claims, err := service.ValidateToken(req.AccessToken, "")

	if err != nil {
		res.Error = "error"

		if errors.Is(err, models.ErrTokenExpired) {
			res.Description = models.ErrTokenExpired.Error()
			w.WriteHeader(400)
		} else if err.Error() == models.ErrTokenInvalid.Error() {
			res.Description = models.ErrTokenInvalid.Error()
			w.WriteHeader(400)
		} else {
			res.Description = "Ошибка сервера"
			w.WriteHeader(500)
		}
		// response
		json.NewEncoder(w).Encode(res)
		return
	}

	// передаем на слой service
	err = u.service.User_create(&claims.Email, &req.Name, &req.Password)
	// Если ошибка
	if err == models.User_err_exists_user {
		res.Status = "error"
		res.Error = "Такой пользователь уже сушествует"
		w.WriteHeader(400)
		// response
		json.NewEncoder(w).Encode(res)
		return
	} else if err != nil {
		res.Status = "error"
		res.Error = "Ошибка сервера"
		w.WriteHeader(500)
		// response
		json.NewEncoder(w).Encode(res)
		return
	}

	// если всё ОК
	res.Status = "Success"
	res.Description = "Пользователь успешно создан"
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(res)
}
