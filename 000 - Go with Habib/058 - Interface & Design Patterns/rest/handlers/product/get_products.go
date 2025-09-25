package product

import (
	"net/http"
	"ecommerce/util"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request){ 
	list, err := h.productRepo.List()
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	util.SendData(w, list, http.StatusOK)
}