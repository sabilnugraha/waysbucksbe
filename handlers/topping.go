package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	dto "waysbucks/dto/result"
	toppingsdto "waysbucks/dto/toppings"
	"waysbucks/models"
	"waysbucks/repositories"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type toppingHandler struct {
	ToppingRepository repositories.ToppingRepository
}

func HandlerTopping(ToppingRepository repositories.ToppingRepository) *toppingHandler {
	return &toppingHandler{ToppingRepository}
}

func convertResponseTopping(u models.Topping) toppingsdto.ToppingResponse {
	return toppingsdto.ToppingResponse{
		ID:    u.ID,
		Title: u.Title,
		Price: u.Price,
		Image: u.Image,
	}
}

func (h *toppingHandler) FindToppings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	toppings, err := h.ToppingRepository.FindToppings()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	// Create Embed Path File on Image property here ...
	// for i, p := range toppings {
	// 	toppings[i].Image = path_file + p.Image
	// }

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: toppings}
	json.NewEncoder(w).Encode(response)
}

func (h *toppingHandler) GetTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	topping, err := h.ToppingRepository.GetTopping(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	// Create Embed Path File on Image property here ...
	topping.Image = path_file + topping.Image

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: topping}
	json.NewEncoder(w).Encode(response)
}

func (h *toppingHandler) CreateTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dataContex := r.Context().Value("dataFile") // add this code
	filepath := dataContex.(string)             // add this code

	price, _ := strconv.Atoi(r.FormValue("price"))

	request := toppingsdto.CreateToppingRequest{
		Title: r.FormValue("title"),
		Price: price,
		Image: filepath,
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	var ctx = context.Background()
	var CLOUD_NAME = os.Getenv("CLOUD_NAME")
	var API_KEY = os.Getenv("API_KEY")
	var API_SECRET = os.Getenv("API_SECRET")

	// Add your Cloudinary credentials ...
	cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

	// Upload file to Cloudinary ...
	resp, err := cld.Upload.Upload(ctx, filepath, uploader.UploadParams{Folder: "waysbuck"})

	if err != nil {
		fmt.Println(err.Error())
	}

	// data form pattern submit to pattern entity db user
	topping := models.Topping{
		Title: request.Title,
		Price: request.Price,
		Image: resp.SecureURL,
	}

	topping, err = h.ToppingRepository.CreateTopping(topping)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	topping, _ = h.ToppingRepository.GetTopping(topping.ID)

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: topping}
	json.NewEncoder(w).Encode(response)
}

func (h *toppingHandler) UpdateTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	price, _ := strconv.Atoi(r.FormValue("price"))
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	dataContex := r.Context().Value("dataFile") // add this code
	filename := dataContex.(string)             // add this code

	request := toppingsdto.CreateToppingRequest{
		Title: r.FormValue("title"),
		Price: price,
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	topping, _ := h.ToppingRepository.GetTopping(id)

	topping.Title = request.Title
	topping.Price = request.Price

	if filename != "false" {
		topping.Image = filename
	}

	data, err := h.ToppingRepository.UpdateTopping(topping)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}

func (h *toppingHandler) DeleteTopping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.ToppingRepository.GetTopping(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.ToppingRepository.DeleteTopping(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "success", Data: data}
	json.NewEncoder(w).Encode(response)
}
