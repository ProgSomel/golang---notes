package user

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type reqCreateUser struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	IsShopOwner bool `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request){
	var req reqCreateUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil{
		fmt.Println(err)
		http.Error(w, "Plz give me a valid json", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userRepo.Create(repo.User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		Password: req.Password,
		IsShopOwner: req.IsShopOwner,

	})
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
	}
	util.SendData(w, createdUser, http.StatusCreated)
	
}