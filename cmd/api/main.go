package main

import (
	"log"
	"net/http"

	"github.com/juraevibrahim01/pharmpro/internal/handler"
	"github.com/juraevibrahim01/pharmpro/internal/reposirory"
	"github.com/juraevibrahim01/pharmpro/internal/service"
	"github.com/juraevibrahim01/pharmpro/pkg"
)

func main() {
	db, err := pkg.InitPostgres()
	if err != nil {
		log.Print("Ошибка сервера при соединении бд: ", err)
		return
	}
	defer db.DB.Close()

	// ---------------------------------- auth ------------------------------------
	auth_repository := reposirory.Auth_new_repository(db)
	auth_service := service.Auth_new_service(auth_repository)
	auth_handler := handler.Auth_new_handler(auth_service)

	// ---------------------------------- user ------------------------------------
	user_repository := reposirory.User_new_repository(db)
	user_service := service.User_new_service(user_repository)
	user_handler := handler.User_new_handler(user_service)

	// ---------------------------------- apis --------------------------------------
	// Маршрутизатор
	mux := http.NewServeMux()

	mux.HandleFunc("POST /login", auth_handler.Login)
	mux.HandleFunc("POST /users", user_handler.User_create)

	// url
	http.ListenAndServe(":8081", nil)
}
