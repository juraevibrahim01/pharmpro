package reposirory

import (
	"database/sql"
	"log"

	"github.com/juraevibrahim01/pharmpro/internal/models"
	"github.com/juraevibrahim01/pharmpro/pkg"
)

type User_repository struct {
	postgres *pkg.Postgres
}

func User_new_repository(postgres *pkg.Postgres) *User_repository {
	return &User_repository{postgres: postgres}
}

func (u *User_repository) User_create(email, name, password *string) error {
	query := `
	insert into users(email, name, password)
	values($1, $2, $3)
	`
	_, err := u.postgres.DB.Exec(query, email, name, password)
	if err != nil {
		log.Print("Ошибка при insert данных: ", err)
		return err
	}
	return nil
}

func (u *User_repository) User_check_exist_user(email *string) error {
	var res_db string

	query := `
		select name
		from users
		where email = $1;
	`
	row := u.postgres.DB.QueryRow(query, email)
	err := row.Scan(&res_db)
	if err == sql.ErrNoRows {
		log.Print("Ошибка: Запрос не нашел ни одно пользователя с почтой: ", *email)
		return nil
	}
	if err != nil {
		log.Print("Ошибка при сканировании", err)
		return err
	}
	return models.User_err_exists_user
}
