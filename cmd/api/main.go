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
	auth_repository := reposirory.NewRepository(db)
	auth_service := service.NewService(auth_repository)
	auth_handler := handler.New_auth_handler(auth_service)

	// ---------------------------------- apis --------------------------------------
	http.HandleFunc("/login", auth_handler.Login)

	// url
	http.ListenAndServe(":8081", nil)
}
