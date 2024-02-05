package web

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

type OrderRestHandler struct {
	ListOrdersUseCase *usecase.ListOrdersUseCase
}

func NewOrderRestHandler(listOrdersUseCase *usecase.ListOrdersUseCase) *OrderRestHandler {
	return &OrderRestHandler{
		ListOrdersUseCase: listOrdersUseCase,
	}
}

func (h *OrderRestHandler) ListOrdersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling ListOrders request")
	orders, err := h.ListOrdersUseCase.Execute()
	if err != nil {
		log.Printf("Error listing orders: %s", err.Error())
		http.Error(w, "Failed to list orders", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(orders)
	if err != nil {
		log.Printf("Error marshaling response: %s", err.Error())
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
