package main

import (
	"encoding/json"
	"net/http"

	"github.com/amauryeuzebio/goexpert/M08-API/configs"
	"github.com/amauryeuzebio/goexpert/M08-API/internal/dto"
	"github.com/amauryeuzebio/goexpert/M08-API/internal/entity"
	"github.com/amauryeuzebio/goexpert/M08-API/internal/infra/database"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := NewProductHandler(productDB)

	http.HandleFunc("/products", productHandler.CreateProduct)

	http.ListenAndServe(":3000", nil)

}

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	err = h.ProductDB.Create(p)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		return
	}

	w.WriteHeader(http.StatusCreated)
}
