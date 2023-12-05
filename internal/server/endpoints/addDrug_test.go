package endpoints

import (
	"bytes"
	"encoding/json"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szmulinho/drugstore/internal/model"
)

func TestAddDrug(t *testing.T) {
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

	newDrug := model.Drug{
		DrugID:      1234,
		Name:        "Test Drug",
		Type:        "Test Type",
		Image:       "Test url",
		Price:       1,
		Description: "Test description",
	}

	newDrugJSON, err := json.Marshal(newDrug)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/drug", bytes.NewBuffer(newDrugJSON))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	h.AddDrug(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var responseDrug model.Drug
	err = json.Unmarshal(w.Body.Bytes(), &responseDrug)
	assert.NoError(t, err)

	assert.Equal(t, newDrug.DrugID, responseDrug.DrugID)
	assert.Equal(t, newDrug.Name, responseDrug.Name)
	assert.Equal(t, newDrug.Type, responseDrug.Type)
	assert.Equal(t, newDrug.Image, responseDrug.Image)
	assert.Equal(t, newDrug.Price, responseDrug.Price)
	assert.Equal(t, newDrug.Description, responseDrug.Description)
}
