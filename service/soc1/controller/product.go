package controller

import (
	"net/http"
	"soc1-service/pkg/log"
	"soc1-service/pkg/router"
	"soc1-service/service/soc1/dao"
)

func SearchProduct(w http.ResponseWriter, r *http.Request) {
	productName := r.URL.Query().Get("productName")
	repo := &dao.RepoProduct{
		Name: productName,
	}

	err := repo.GetProdcuts()
	if err != nil {
		log.Println(log.LogLevelDebug, "CreateAccount: account.CheckAccountExist()", err)
		router.ResponseInternalError(w, "Error", err)
		return
	}

	router.ResponseCreatedWithData(w, "B.200", "Get products successfully", repo)
}
