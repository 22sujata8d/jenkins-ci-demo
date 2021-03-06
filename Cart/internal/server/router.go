package server

import (
	"net/http"

	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/literals"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/internal/server/httphandlers"
	"github.com/swiggy-2022-bootcamp/cdp-team2/cart/util"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	*mux.Router
}

func NewRouter() *Router {
	return &Router{mux.NewRouter()}
}

func (r *Router) InitializeRouter(routerConfig *util.RouterConfig) {
	r.InitializeRoutes(routerConfig)
}

func (r *Router) InitializeRoutes(routerConfig *util.RouterConfig) {
	s := (*r).PathPrefix(routerConfig.WebServerConfig.RoutePrefix).Subrouter()
	// Swagger
	s.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	s.HandleFunc(literals.HealthCheckEndpoint,
		httphandlers.HealthCheckHandler(routerConfig)).
		Methods(http.MethodGet).
		Name(literals.HealthCheckAPIName)

	s.HandleFunc(literals.GetCartEndpoint,
		httphandlers.GetCartHandler(routerConfig)).
		Methods(http.MethodGet).
		Name(literals.GetCartAPIName)

	s.HandleFunc(literals.AddCartItemEndpoint,
		httphandlers.AddCartItemHandler(routerConfig)).
		Methods(http.MethodPost).
		Name(literals.AddCartAPIName)

	s.HandleFunc(literals.UpdateCartItemEndpoint,
		httphandlers.UpdateCartItemHandler(routerConfig)).
		Methods(http.MethodPut).
		Name(literals.UpdateCartItemAPIName)

	s.HandleFunc(literals.DeleteCartItemEndpoint,
		httphandlers.DeleteCartItemHandler(routerConfig)).
		Methods(http.MethodDelete).
		Name(literals.DeleteCartItemAPIName)

}
