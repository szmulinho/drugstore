package endpoints

import (
	"encoding/json"
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

func TestGetDrugByName(t *testing.T) {
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
		Name: "Test Drug",
	}

	model.Drugs = []model.Drug{testDrug}

	request, err := http.NewRequest("GET", "/drugs/"+testDrug.Name, nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/drugs/{name}", h.GetDrugByName).Methods("GET")

	router.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)

	var responseDrugs []model.Drug
	err = json.Unmarshal(recorder.Body.Bytes(), &responseDrugs)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(responseDrugs))
	assert.Equal(t, testDrug.Name, responseDrugs[0].Name)
}
