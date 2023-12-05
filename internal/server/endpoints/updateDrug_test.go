package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/szmulinho/drugstore/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUpdateDrug(t *testing.T) {
	err := godotenv.Load(".env.test")
	if err != nil {
		t.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + user + " dbname=" + name + " sslmode=require password=" + password + " port=" + port

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	defer db.DB()

	db.AutoMigrate(&model.Drug{})

	h := &handlers{db: db}

	testDrug := model.Drug{
	}
	db.Create(&testDrug)

	updatedDrug := model.Drug{
		Name:  "Test Drug",
		Price: 20,
	}

	updatedDrugJSON, err := json.Marshal(updatedDrug)
	if err != nil {
		t.Fatal(err)
	}
	request, err := http.NewRequest("PUT", fmt.Sprintf("/drugs/%v", testDrug.DrugID), bytes.NewBuffer(updatedDrugJSON))
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/drugs/{id}", h.UpdateDrug).Methods("PUT")

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var updatedDrugFromDB model.Drug
	db.First(&updatedDrugFromDB, testDrug.DrugID)

	assert.Equal(t, updatedDrug.Name, updatedDrugFromDB.Name)
	assert.Equal(t, updatedDrug.Price, updatedDrugFromDB.Price)
}
