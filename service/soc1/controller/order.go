package controller

import (
	"encoding/json"
	"net/http"
	"soc1-service/pkg/log"
	"soc1-service/pkg/router"
	"soc1-service/service/soc1/dao"
	"strconv"
)

type InfoOrder struct {
	Email string          `json:"email"`
	Infos []dao.InfoOrder `json:"infos"`
}

func Order(w http.ResponseWriter, r *http.Request) {
	var infoOrder = &InfoOrder{}
	err := json.NewDecoder(r.Body).Decode(infoOrder)
	if err != nil {
		log.Println(log.LogLevelDebug, "Order: Decode(infoOrder)", err)
		router.ResponseInternalError(w, "decode infoOrder failed", err)
		return
	}

	order := &dao.OrderDAO{
		Email: infoOrder.Email,
	}

	err = order.CreateOrder()
	if err != nil {
		log.Println(log.LogLevelDebug, "Order: order.CreateOrder()", err)
		router.ResponseInternalError(w, "Create Order failed", err)
		return
	}

	for _, i := range infoOrder.Infos {
		info := &dao.InfoOrder{
			IdOrder:   order.Id,
			IdProduct: i.IdProduct,
			Amount:    i.Amount,
		}

		err = info.AddInfo()
		if err != nil {
			log.Println(log.LogLevelDebug, "Order: order.CreateOrder()", err)
			router.ResponseInternalError(w, "Create Order failed", err)
			return
		}

		product := &dao.ProductDAO{
			Id:     i.IdProduct,
			Amount: i.Amount,
		}

		err = product.UpdateAmount()
		if err != nil {
			log.Println(log.LogLevelDebug, "Order: product.UpdateAmount()", err)
			router.ResponseInternalError(w, "Create Order failed", err)
			return
		}
	}

	router.ResponseSuccess(w, "B.200", "Order successfully")
}

func FindOrders(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")

	repo := &dao.RepoOrder{
		Email: email,
	}

	err := repo.GetOrders()
	if err != nil {
		log.Println(log.LogLevelDebug, "FindOrders: repo.GetOrders()", err)
		router.ResponseInternalError(w, "Create Order failed", err)
		return
	}

	router.ResponseSuccessWithData(w, "B.200", "Get orders successfully", repo)
}

func UpdateStatus(w http.ResponseWriter, r *http.Request) {
	idOrder := r.URL.Query().Get("idOrder")

	id, err := strconv.Atoi(idOrder)
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateStatus: strconv.Atoi(idOrder)", err)
		router.ResponseInternalError(w, "Error", err)
		return
	}

	order := &dao.OrderDAO{
		Id: int32(id),
	}

	err = order.UpdateStatus()
	if err != nil {
		log.Println(log.LogLevelDebug, "UpdateStatus: order.UpdateStatus()", err)
		router.ResponseInternalError(w, "Error", err)
		return
	}

	router.ResponseSuccess(w, "B.200", "Update status successfully")
}
