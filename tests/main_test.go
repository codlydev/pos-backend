package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"pos-backend/config"
	"pos-backend/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

var authToken string

func TestMain(m *testing.M) {
	_ = godotenv.Load("../.env")
	config.InitDB()
	authToken = getTestToken()
	os.Exit(m.Run())
}

func SetupRouter() *gin.Engine {
	return routes.SetupRouter()
}

func TestRegisterAndLogin(t *testing.T) {
	router := SetupRouter()

	registerPayload := map[string]string{
		"username": "user1",
		"password": "pass123",
	}
	body, _ := json.Marshal(registerPayload)

	// Register
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("➡️  [POST] /register Payload: %v", registerPayload)
	t.Logf("📥 Response Code: %d", w.Code)
	t.Logf("📦 Response Body: %s", w.Body.String())
	assert.Equal(t, 200, w.Code)

	// Login
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("➡️  [POST] /login Payload: %v", registerPayload)
	t.Logf("📥 Response Code: %d", w.Code)
	t.Logf("📦 Response Body: %s", w.Body.String())
	assert.Equal(t, 200, w.Code)

	respBody, _ := io.ReadAll(w.Body)
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)
	tok, ok := result["token"].(string)
	assert.True(t, ok)
	t.Logf("✅ Login Token: %s", tok)
}

func TestCreateProduct(t *testing.T) {
	router := SetupRouter()

	productPayload := map[string]interface{}{
		"name":  "Test Product",
		"price": 19.99,
	}
	body, _ := json.Marshal(productPayload)
	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("➡️  [POST] /products Payload: %v", productPayload)
	t.Logf("📥 Response Code: %d", w.Code)
	t.Logf("📦 Response Body: %s", w.Body.String())

	assert.Equal(t, 201, w.Code)
}

func TestGetProducts(t *testing.T) {
	router := SetupRouter()

	req, _ := http.NewRequest("GET", "/products", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("➡️  [GET] /products")
	t.Logf("📥 Response Code: %d", w.Code)
	t.Logf("📦 Response Body: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

func TestCreateSale(t *testing.T) {
	router := SetupRouter()

	salePayload := map[string]interface{}{
		"product_id": 1,
		"quantity":   2,
	}
	body, _ := json.Marshal(salePayload)
	req, _ := http.NewRequest("POST", "/sales", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("➡️  [POST] /sales Payload: %v", salePayload)
	t.Logf("📥 Response Code: %d", w.Code)
	t.Logf("📦 Response Body: %s", w.Body.String())

	assert.Equal(t, 201, w.Code)
}

func TestGetSales(t *testing.T) {
	router := SetupRouter()

	req, _ := http.NewRequest("GET", "/sales", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authToken))

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	t.Logf("➡️  [GET] /sales")
	t.Logf("📥 Response Code: %d", w.Code)
	t.Logf("📦 Response Body: %s", w.Body.String())

	assert.Equal(t, 200, w.Code)
}

// 🔑 Helper function to register and login user and return token
func getTestToken() string {
	router := SetupRouter()

	payload := map[string]string{
		"username": "testuser",
		"password": "testpass",
	}
	body, _ := json.Marshal(payload)

	// Register
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Login
	req, _ = http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		panic("❌ Login failed during test setup")
	}

	respBody, _ := io.ReadAll(w.Body)
	var result map[string]interface{}
	json.Unmarshal(respBody, &result)

	token, ok := result["token"].(string)
	if !ok {
		panic("❌ Token not found in login response")
	}

	return token
}
