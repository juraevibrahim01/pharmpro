package reposirory

import (
	"log"

	"github.com/juraevibrahim01/pharmpro/pkg"
)

type Product_reposirory struct {
	postgres *pkg.Postgres
}

func Product_new_repository(postgres *pkg.Postgres) *Product_reposirory {
	return &Product_reposirory{postgres: postgres}
}

func (p *Product_reposirory) Product_insert(product_name, firm_name, country_product *string, amount, amount_sale *float64, markup *float32, quantity_in_pack *int32, unit_id *int8) error {
	query := `
		insert into products
			(name, firm, country, amount, amount_sale, markup, quantity_in_pack, unit_id)
			values
			($1, $2, $3, $4, $5, $6, $7, $8)
	`
	_, err := p.postgres.DB.Exec(query, product_name, firm_name, country_product, amount, amount_sale, markup, quantity_in_pack, unit_id)
	if err != nil {
		log.Println("func (p *Product_reposirory) Product_insert errors: ", err)
		return err
	}
	return nil
}
