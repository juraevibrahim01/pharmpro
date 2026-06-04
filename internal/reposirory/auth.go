package reposirory

import (
	"database/sql"
	"log"

	"github.com/juraevibrahim01/pharmpro/pkg"
)

type Auth_reposirory struct {
	postgres *pkg.Postgres
}

func Auth_new_repository(postgres *pkg.Postgres) *Auth_reposirory {
	return &Auth_reposirory{postgres: postgres}
}

// Проверка на идентификацию
func (r *Auth_reposirory) Reposirory_identification(email *string) (int, error) {

	var res_db int

	query := `
		select id
		from users
		where email = $1;
	`
	row := r.postgres.DB.QueryRow(query, email)

	err := row.Scan(&res_db)
	if err == sql.ErrNoRows {
		log.Print("Ошибка: Запрос не нашел ни одно пользователя с почтой: ", *email)
		return 0, err
	}
	if err != nil {
		log.Print("Ошибка при сканировании", err)
		return 0, err
	}

	return res_db, nil
}

// Проверка что пороли совподают
func (r *Auth_reposirory) Reposirory_check_password(id *int) (string, error) {
	var res_db string

	query := `
		select password
		from users
		where id = $1;
	`
	row := r.postgres.DB.QueryRow(query, id)
	err := row.Scan(&res_db)
	if err == sql.ErrNoRows {
		log.Print("Ошибка: Пароли не совпадают, пользователя с идентификатором: ", *id)
		return "", err
	}
	if err != nil {
		log.Print("Ошибка при сканировании", err)
		return "", err
	}
	return res_db, nil
}
