package task

import (
	"net/http"

	"github.com/gorilla/mux"

	ctrl "backend_test/task/controllers"
)

func PaymentHandler(muxRoutes *mux.Router) *mux.Router {
	payment := muxRoutes.PathPrefix("/api/tasks").Subrouter()

	payment.Path("/").Handler(http.HandlerFunc(ctrl.GetPaymentAPI)).Methods("GET")
	payment.Path("/").Handler(http.HandlerFunc(ctrl.SavePaymentAPI)).Methods("POST")
	payment.Path("/list").Handler(http.HandlerFunc(ctrl.GetAllPaymentAPI)).Methods("GET")

	return muxRoutes
}
