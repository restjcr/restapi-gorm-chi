package tests

import (
	"github.com/go-chi/chi/v5"
	"github.com/restjcr/restapi-gorm-chi/database"
	"github.com/restjcr/restapi-gorm-chi/model"
	"testing"
)

func TestGetProductHandler(t *testing.T) {
	r := chi.NewRouter()

	db := database.NewConnectionForTest()
	err := db.AutoMigrate(&model.Product{})

	if err != nil {
		t.Fatal("Cannot automigrate product model")
	}

}
