package dao

import (
	"soc1-service/pkg/db"
)

type RepoInfo struct {
	IdOrder int32
	Infos   []InfoOrder
}

type InfoOrder struct {
	IdOrder   int32 `json:"idOrder"`
	IdProduct int32 `json:"idProduct"`
	Amount    int32 `json:"amount"`
}

func (info *InfoOrder) AddInfo() error {
	query := `INSERT INTO public.info_order (id_order, id_product, amount) VALUES($1, $2, $3);`
	_, err := db.PSQL.Exec(query, info.IdOrder, info.IdProduct, info.Amount)
	return err
}

func (repo *RepoInfo) GetInfos() error {
	query := `SELECT * FROM public.info_order where id_order= $1;`
	rows, err := db.PSQL.Query(query, repo.IdOrder)
	if err != nil {
		return err
	}

	defer rows.Close()

	for rows.Next() {
		var info = &InfoOrder{}
		err := rows.Scan(&info.IdOrder, &info.IdProduct, &info.Amount)
		if err != nil {
			return err
		}
		repo.Infos = append(repo.Infos, *info)
	}
	return nil
}
