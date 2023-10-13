package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/restjcr/restapi-gorm-chi/model"
	"github.com/restjcr/restapi-gorm-chi/service"
	"log"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: service,
	}
}

func (h *ProductHandler) GetAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("In the future this response will be a full list of products!"))
	products, err := h.service.GetAllProducts()

	if err != nil {
		http.Error(w, "Cannot retrieve list of products from db: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(products)

	if err != nil {
		log.Println("Cannot encode the products slice")
	}

}

func (h *ProductHandler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		log.Println("Cannot convert string productId to int")
		return
	}

	product, err := h.service.GetProduct(productId)

	if err != nil {
		//log.Println("Cannot retrieve product by id given")
		w.Write([]byte("Cannot retrieve product by id given"))
		return
	}

	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(product)

	if err != nil {
		log.Println("Cannot encode the product to json")
	}

}

func (h *ProductHandler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	var newProduct model.Product

	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		http.Error(w, "Error al decodificar el JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.CreateProduct(newProduct)

	if err != nil {
		w.Write([]byte("Cannot insert that record"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Record succesful registered to the database"))

}

func (h *ProductHandler) UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Cannot read product id sent"+err.Error(), http.StatusBadRequest)
		return
	}

	var actualProduct model.Product

	if err := json.NewDecoder(r.Body).Decode(&actualProduct); err != nil {
		http.Error(w, "Error al decodificar el JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	var updatedProduct *model.Product

	updatedProduct, err = h.service.UpdateProduct(actualProduct, productId)

	if err != nil {
		http.Error(w, "Cannot update the product: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(updatedProduct)

	if err != nil {
		log.Println("Cannot encode the product to json")
	}

}

func (h *ProductHandler) DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	productId, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		http.Error(w, "Cannot read product id sent: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.DeleteProduct(productId)

	if err != nil {
		http.Error(w, "Cannot delete that product: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Deleted product"))

}
