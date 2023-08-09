package dao

import (
	"soc1-service/pkg/db"
)

type AccountDAO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (account *AccountDAO) InsertAccount() error {
	query := `INSERT INTO public.account (email, password) VALUES ( $1, $2);`
	_, err := db.PSQL.Exec(query, account.Email, account.Password)
	return err
}

func (account *AccountDAO) CheckAccountExist() (bool, error) {
	var exist bool
	query := `SELECT exists(select 1 FROM account WHERE email = $1);`
	err := db.PSQL.QueryRow(query, account.Email).Scan(&exist)
	return exist, err
}

func (account *AccountDAO) GetAccountByEmail() error {
	query := `SELECT * FROM account WHERE email = $1;`
	err := db.PSQL.QueryRow(query, account.Email).Scan(&account.Email, &account.Password)
	return err
}
