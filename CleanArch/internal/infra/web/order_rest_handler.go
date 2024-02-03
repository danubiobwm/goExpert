package web

import (
	"encoding/json"
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
	orders, err := h.ListOrdersUseCase.Execute()
	if err != nil {
		http.Error(w, "Failed to list orders", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(orders)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
