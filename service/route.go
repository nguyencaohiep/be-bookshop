package service

import (
	"soc1-service/pkg/router"
	"soc1-service/service/index"
	"soc1-service/service/soc1"

	"github.com/go-chi/chi/middleware"
)

// LoadRoutes to Load Routes to Router
func LoadRoutes() {

	router.Router.Use(middleware.RealIP)

	//* Set Endpoint for admin
	router.Router.Get(router.RouterBasePath+"/", index.GetIndex)
	router.Router.Get(router.RouterBasePath+"/health", index.GetHealth)
	router.Router.Mount(router.RouterBasePath+"/soc1", soc1.SOCServiceSubRoute)
}
