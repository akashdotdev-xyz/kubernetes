package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/akashdotdev-xyz/order-service/internal/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: service,
	}
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()

	response := map[string]any{
		"pod":    hostname,
		"orders": h.service.GetOrders(),
	}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(response)
}
