package endpoints

import (
	"encoding/json"
	"github.com/szmulinho/drugstore/internal/model"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestGetAllDrugs(t *testing.T) {
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

	req, err := http.NewRequest("GET", "/drugs", nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()

	h.GetAllDrugs(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var drugs []model.Drug
	err = json.Unmarshal(w.Body.Bytes(), &drugs)
	assert.NoError(t, err)

}
