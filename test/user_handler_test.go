package test

import (
	"auth-service/internal/domain/model"
	"auth-service/internal/handler"
	repository "auth-service/internal/repository/db"
	"auth-service/internal/usecase"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestRouter(t *testing.T) (*gin.Engine, *gorm.DB) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	assert.NoError(t, err)

	db.AutoMigrate(&model.User{})

	repo := repository.NewUserRepository(db)
	uc := usecase.NewUserUsecase(repo)
	h := handler.NewUserHandler(uc)

	r := gin.Default()
	api := r.Group("/v1")
	userGroup := api.Group("/user")
	{
		userGroup.POST("/register", h.Register)
		userGroup.POST("/login", h.Login)
	}
	return r, db
}

func TestRegisterSuccess(t *testing.T) {
	r, _ := setupTestRouter(t)

	payload := map[string]string{
		"name":     "testing",
		"email":    "test@example.com",
		"password": "password123",
	}
	body, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/user/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
}

func TestLoginSuccess(t *testing.T) {
	r, db := setupTestRouter(t)

	// Insert user manual ke DB
	db.Create(&model.User{
		Email:    "test@example.com",
		Password: "$2a$10$7u1z2eOjTfYFZ/eDqk3VY.QXK3WbBZlqWaZXw3SYj3EsLOpblbR9W", // hash dari 'password123'
	})

	payload := map[string]string{
		"email":    "test@example.com",
		"password": "password123",
	}
	body, _ := json.Marshal(payload)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/user/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}
