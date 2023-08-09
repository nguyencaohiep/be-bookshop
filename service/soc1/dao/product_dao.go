package dao

import (
	"soc1-service/pkg/db"
	"strings"
)

var Products map[int32]ProductDAO

type RepoProduct struct {
	Name     string       `json:"name"`
	Products []ProductDAO `json:"products"`
}

type ProductDAO struct {
	Id     int32   `json:"id"`
	Name   string  `json:"name"`
	Amount int32   `json:"amount"`
	Image  string  `json:"image"`
	Price  float32 `json:"price"`
}

func init() {
	Products = map[int32]ProductDAO{}
	query := `SELECT * FROM product;`
	rows, err := db.PSQL.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var product = &ProductDAO{}
		err := rows.Scan(&product.Id, &product.Name, &product.Amount, &product.Image, &product.Price)
		if err != nil {
			return
		}
		Products[product.Id] = *product

	}
}

func (repo *RepoProduct) GetProdcuts() error {
	query := `SELECT * FROM product where lower(name) like $$%` + strings.ToLower(repo.Name) + `%$$  ;`
	rows, err := db.PSQL.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var product = &ProductDAO{}
		err := rows.Scan(&product.Id, &product.Name, &product.Amount, &product.Image, &product.Price)
		if err != nil {
			return err
		}
		repo.Products = append(repo.Products, *product)

	}
	return nil
}

func (product *ProductDAO) UpdateAmount() error {
	query := `UPDATE public.product SET amount= amount - $1 where id = $2;`
	_, err := db.PSQL.Exec(query, product.Amount, product.Id)
	return err
}
