package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type ReqUpdateProduct struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgURL string `json:"imageURL"`
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request){
	productiID := r.PathValue("id")

	pId, err := strconv.Atoi(productiID)
	if err != nil{
		util.SendError(w, http.StatusBadRequest, "invalid product id")
		return
	}

	var req ReqUpdateProduct
	
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&req)
	if err != nil{
		util.SendError(w, http.StatusBadRequest, "Please give me valid json")
		return
	}
	
	product, err := h.productRepo.Update(repo.Product{
		ID: pId,
		Title: req.Title,
		Description: req.Description,
		Price: req.Price,
		ImgURL: req.ImgURL,
	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.SendData(w, product, http.StatusOK)
}