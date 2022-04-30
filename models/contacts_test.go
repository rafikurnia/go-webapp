package models

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/rafikurnia/go-webapp/utils"

	"github.com/stretchr/testify/assert"
)

// A mock test of data structure of Contact struct that is used to define how
// the data is stored on the actual database, on GetAll() method.
func TestContactsGetAll(t *testing.T) {
	dbModels := &models{
		Contacts: &mockContactModel{},
	}

	dbModels.Contacts.Init()
	data, _ := dbModels.Contacts.GetAll()
	actualOutput, _ := json.Marshal(data)

	expectedOutput := `[
		{"id": 1, "name": "Rafi","email": "rafi@rafi.com"},
		{"id": 2, "name": "Kurnia","email": "kurnia@kurnia.com"},
		{"id": 3, "name": "Putra","email": "putra@putra.com"}
	]`

	assert.JSONEq(t, expectedOutput, string(actualOutput))
}

// Test to retrieve a contact data by its ID
func TestContactsGetById(t *testing.T) {
	dbModels := &models{
		Contacts: &mockContactModel{},
	}

	dbModels.Contacts.Init()

	t.Run("Valid ID", func(t *testing.T) {
		data, _ := dbModels.Contacts.GetById(1)
		actualOutput, _ := json.Marshal(data)

		expectedOutput := `{"id": 1, "name": "Rafi","email": "rafi@rafi.com"}`

		assert.JSONEq(t, expectedOutput, string(actualOutput))
	})

	t.Run("Invalid ID", func(t *testing.T) {
		_, err := dbModels.Contacts.GetById(9)

		if !errors.Is(err, utils.ErrorNotFound) {
			t.Errorf("Expect %v, got %v", utils.ErrorNotFound, err)
		}
	})
}

// Test to create a new contact
func TestContactsCreate(t *testing.T) {
	dbModels := &models{
		Contacts: &mockContactModel{},
	}

	dbModels.Contacts.Init()

	t.Run("Valid Operation", func(t *testing.T) {
		newData := contact{Name: "Rafi Kurnia", Email: "rafi@kurnia.com"}

		data, _ := dbModels.Contacts.Create(&newData)
		actualOutput, _ := json.Marshal(data)

		expectedOutput := `{"id": 4, "name": "Rafi Kurnia","email": "rafi@kurnia.com"}`

		assert.JSONEq(t, expectedOutput, string(actualOutput))
	})
}

// Test to delete a contact by its ID
func TestContactsDeleteById(t *testing.T) {
	dbModels := &models{
		Contacts: &mockContactModel{},
	}

	dbModels.Contacts.Init()

	t.Run("Valid Operation", func(t *testing.T) {
		dbModels.Contacts.DeleteById(3)

		data, _ := dbModels.Contacts.GetAll()
		actualOutput, _ := json.Marshal(data)

		expectedOutput := `[
			{"id": 1, "name": "Rafi","email": "rafi@rafi.com"},
			{"id": 2, "name": "Kurnia","email": "kurnia@kurnia.com"}
		]`

		assert.JSONEq(t, expectedOutput, string(actualOutput))
	})

	t.Run("Delete on Invalid ID", func(t *testing.T) {
		data, _ := dbModels.Contacts.GetAll()
		output, _ := json.Marshal(data)

		numberOfDataBeforeDeletion := len(output)

		dbModels.Contacts.DeleteById(10)

		data, _ = dbModels.Contacts.GetAll()
		output, _ = json.Marshal(data)

		numberOfDataAfterDeletion := len(output)

		assert.Equal(t, numberOfDataBeforeDeletion, numberOfDataAfterDeletion)
	})
}

// Test to update a contact by its ID
func TestContactsUpdateById(t *testing.T) {
	dbModels := &models{
		Contacts: &mockContactModel{},
	}

	dbModels.Contacts.Init()

	t.Run("Valid Operation", func(t *testing.T) {
		newData := &contact{Name: "Rafi Kurnia Putra", Email: "rafi.kurnia.putra@rafi.com"}

		dbModels.Contacts.UpdateById(1, newData)

		data, _ := dbModels.Contacts.GetById(1)
		actualOutput, _ := json.Marshal(data)

		expectedOutput := `{"id": 1, "name": "Rafi Kurnia Putra","email": "rafi.kurnia.putra@rafi.com"}`

		assert.JSONEq(t, expectedOutput, string(actualOutput))
	})
}
