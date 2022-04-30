package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafikurnia/go-webapp/models"

	"github.com/stretchr/testify/assert"
)

// Test GET /api/v1/contacts
func TestGetAllContacts(t *testing.T) {
	router := getRouterForTest()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/contacts", nil)
	router.ServeHTTP(w, req)

	expectedOutput := `[
		{"id": 1, "name": "Rafi","email": "rafi@rafi.com"},
		{"id": 2, "name": "Kurnia","email": "kurnia@kurnia.com"},
		{"id": 3, "name": "Putra","email": "putra@putra.com"}
	]`

	assert.JSONEq(t, expectedOutput, w.Body.String())
	assert.Equal(t, 200, w.Code)
}

// Test GET /api/v1/contacts/:id
func TestGetAContactById(t *testing.T) {
	router := getRouterForTest()

	t.Run("Valid ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/contacts/1", nil)
		router.ServeHTTP(w, req)

		expectedOutput := `{"id": 1, "name": "Rafi","email": "rafi@rafi.com"}`

		assert.JSONEq(t, expectedOutput, w.Body.String())
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/contacts/10", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
	})

	t.Run("Invalid Parameter", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/contacts/invalid", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}

// Test POST /api/v1/contacts
func TestCreateAContact(t *testing.T) {
	t.Run("Valid Request", func(t *testing.T) {
		router := getRouterForTest()

		w := httptest.NewRecorder()

		newContact := models.NewContact()
		newContact.Name = "Rafi"
		newContact.Email = "rafi@rafi.com"

		inJson, _ := json.Marshal(newContact)

		req, _ := http.NewRequest("POST", "/api/v1/contacts", bytes.NewBuffer(inJson))
		router.ServeHTTP(w, req)

		assert.Equal(t, 201, w.Code)
	})
}

// Test DELETE /api/v1/contacts/:id
func TestDeleteAContactById(t *testing.T) {
	router := getRouterForTest()

	t.Run("Valid ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/v1/contacts/3", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 204, w.Code)
	})

	t.Run("Invalid ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/v1/contacts/10", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
	})

	t.Run("Invalid Parameter", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/v1/contacts/invalid", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}

// Test PUT /api/v1/contacts/:id
func TestUpdateAContactById(t *testing.T) {
	router := getRouterForTest()

	t.Run("Valid Request", func(t *testing.T) {
		w := httptest.NewRecorder()

		newContact := models.NewContact()
		newContact.Name = "Rafi Kurnia"
		newContact.Email = "rafi@kurnia.com"

		inJson, _ := json.Marshal(newContact)

		req, _ := http.NewRequest("PUT", "/api/v1/contacts/1", bytes.NewBuffer(inJson))
		router.ServeHTTP(w, req)

		expectedOutput := `{"id": 1, "name": "Rafi Kurnia","email": "rafi@kurnia.com"}`

		assert.JSONEq(t, expectedOutput, w.Body.String())
		assert.Equal(t, 200, w.Code)
	})

	t.Run("Invalid Request non existence resource", func(t *testing.T) {
		w := httptest.NewRecorder()

		newContact := models.NewContact()
		newContact.Name = "Rafi Kurnia"
		newContact.Email = "rafi@kurnia.com"

		inJson, _ := json.Marshal(newContact)

		req, _ := http.NewRequest("PUT", "/api/v1/contacts/10", bytes.NewBuffer(inJson))
		router.ServeHTTP(w, req)

		assert.Equal(t, 404, w.Code)
	})

	t.Run("Invalid path", func(t *testing.T) {
		w := httptest.NewRecorder()

		newContact := models.NewContact()
		newContact.Name = "Rafi Kurnia"
		newContact.Email = "rafi@kurnia.com"

		inJson, _ := json.Marshal(newContact)

		req, _ := http.NewRequest("PUT", "/api/v1/contacts/invalid", bytes.NewBuffer(inJson))
		router.ServeHTTP(w, req)

		assert.Equal(t, 400, w.Code)
	})
}
