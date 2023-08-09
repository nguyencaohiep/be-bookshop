package soc1

import (
	"soc1-service/service/soc1/controller"

	"github.com/go-chi/chi"
)

var SOCServiceSubRoute = chi.NewRouter()

// Init package with sub-router for account service
func init() {
	SOCServiceSubRoute.Group(func(_ chi.Router) {
		//account
		SOCServiceSubRoute.Post("/signup", controller.CreateAccount)
		SOCServiceSubRoute.Post("/signin", controller.SignIn)

		// product
		SOCServiceSubRoute.Get("/product/search", controller.SearchProduct)

		// order
		SOCServiceSubRoute.Post("/order/add", controller.Order)
		SOCServiceSubRoute.Get("/order/find", controller.FindOrders)
		SOCServiceSubRoute.Patch("/order/update", controller.UpdateStatus)
	})
}
