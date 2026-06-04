package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/juraevibrahim01/pharmpro/internal/models"
	"github.com/juraevibrahim01/pharmpro/internal/service"
)

type Product_handler struct {
	service *service.Product_service
}

func Product_new_service(service *service.Product_service) *Product_handler {
	return &Product_handler{service: service}
}

func (p *Product_handler) Product_insert(w http.ResponseWriter, r *http.Request) {

	// Response | Request
	var req models.Product_Req_login
	var res models.Product_Res_login

	// парсим данные
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Ошибка сервера", 500)
		return
	}

	// Проверка: Все поля возвращются
	if req.AccessToken == "" || req.Amount == 0 || req.Amount_sale == 0 || req.Country_product == "" || req.Firm_name == "" || req.Markup == 0 || req.Product_name == "" || req.Quantity_in_pack == 0 || req.Unit_id == 0 {
		res.Status = "error"
		res.Error = "Поля не валидны"
		w.WriteHeader(400)
		// response
		json.NewEncoder(w).Encode(res)
		return
	}

	_, err = service.ValidateToken(req.AccessToken, "")
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

	// слой сервис
	err = p.service.Product_insert(&req.Product_name, &req.Firm_name, &req.Country_product, &req.Amount, &req.Amount_sale, &req.Markup, &req.Quantity_in_pack, &req.Unit_id)
	if err != nil {
		res.Status = "error"
		res.Error = "Ошибка сервера"
		w.WriteHeader(500)
		// response
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Status = "success"
	res.Description = "Товар уcпешно добавлен"
	w.WriteHeader(201)
	// response
	json.NewEncoder(w).Encode(res)
}
