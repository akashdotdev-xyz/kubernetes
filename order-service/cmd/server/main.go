package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/akashdotdev-xyz/order-service/internal/handler"
	"github.com/akashdotdev-xyz/order-service/internal/service"
)

func main() {
	router := chi.NewRouter()

	orderService := service.NewOrderService()
	orderHandler := handler.NewOrderHandler(orderService)

	router.Get("/health", handler.Health)
	router.Get("/orders", orderHandler.GetOrders)

	log.Println("Server listening on :8080")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
