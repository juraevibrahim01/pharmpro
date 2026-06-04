package service

import (
	"log"

	"github.com/juraevibrahim01/pharmpro/internal/reposirory"
)

type Product_service struct {
	reposirory *reposirory.Product_reposirory
}

func Product_new_service(reposirory *reposirory.Product_reposirory) *Product_service {
	return &Product_service{reposirory: reposirory}
}

func (p *Product_service) Product_insert(product_name, firm_name, country_product *string, amount, amount_sale *float64, markup *float32, quantity_in_pack *int32, unit_id *int8) error {
	err := p.reposirory.Product_insert(product_name, firm_name, country_product, amount, amount_sale, markup, quantity_in_pack, unit_id)
	if err != nil {
		log.Println("func (p *Product_service) Product_insert", err)
		return err
	}
	return nil
}
