package dao

import (
	"soc1-service/pkg/db"
	"soc1-service/pkg/utils"
)

type RepoOrder struct {
	Email  string     `json:"email"`
	Orders []OrderDAO `json:"orders"`
}

type OrderDAO struct {
	Id         int32        `json:"id"`
	Status     bool         `json:"status"`
	Email      string       `json:"email"`
	CreateDate string       `json:"createDate"`
	Info       []ProductDAO `json:"info"`
}

func (order *OrderDAO) CreateOrder() error {
	query := `INSERT INTO public.orders (email, createdate, status) VALUES($1, $2, $3) returning id;`
	err := db.PSQL.QueryRow(query, order.Email, utils.TimeNowString(), false).Scan(&order.Id)
	return err
}

func (order *OrderDAO) UpdateStatus() error {
	query := `UPDATE public.orders SET status=true where id= $1;`
	_, err := db.PSQL.Exec(query, order.Id)
	return err
}

func (repo *RepoOrder) GetOrders() error {
	query := `SELECT * FROM public.orders where email= $1;`
	rows, err := db.PSQL.Query(query, repo.Email)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var order = &OrderDAO{}
		err := rows.Scan(&order.Id, &order.Email, &order.CreateDate, &order.Status)
		if err != nil {
			return err
		}

		repoInfo := &RepoInfo{
			IdOrder: order.Id,
		}

		err = repoInfo.GetInfos()
		if err != nil {
			return err
		}

		for _, info := range repoInfo.Infos {
			product := Products[info.IdProduct]
			product.Amount = info.Amount
			order.Info = append(order.Info, product)
		}
		repo.Orders = append(repo.Orders, *order)
	}

	return nil
}
