package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/rafikurnia/go-webapp/utils"

	"gorm.io/gorm"
)

// Data schema of contact table
type contact struct {
	ID    uint   `json:"id"   gorm:"primaryKey;autoIncrement"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Data schema of contact table for the API documentation in Swagger
type contactInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// An interface that links consumer and producer of contact data to enable loosely coupled
// integration
type contactLinker interface {
	Init() error
	GetAll() (*[]contact, error)
	GetById(id uint) (*contact, error)
	Create(data *contact) (*contact, error)
	DeleteById(id uint) error
	UpdateById(id uint, data *contact) (*contact, error)
}

// A struct that connects consumer of contact table to the database via
// gorm library. It implements contactLinker interface.
type contactModel struct {
	DB *gorm.DB
}

// A method to initialize table in the database. The migration is handled by
// gorm. It returns error if the migration fails.
func (c *contactModel) Init() error {
	if err := c.DB.AutoMigrate(&contact{}); err != nil {
		return fmt.Errorf("Init() -> %w", err)
	}
	return nil
}

// A method to retrieve all data in contact table in the database.
// It returns a slice of all contacts and any error encountered.
func (c *contactModel) GetAll() (*[]contact, error) {
	contacts := &[]contact{}

	if err := c.DB.Find(contacts).Error; err != nil {
		return nil, fmt.Errorf("GetAll() -> %w", err)
	}
	return contacts, nil
}

// A method to retrieve a contact by its ID.
// It returns a contact object and any error encountered.
func (c *contactModel) GetById(id uint) (*contact, error) {
	contact := &contact{}

	if err := c.DB.First(contact, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, utils.ErrorNotFound
		}
		return nil, fmt.Errorf("GetById(id uint) -> %w", err)
	}
	return contact, nil
}

// A method to add a new contact to DB
// It returns the newly created contact object and any error encountered.
func (c *contactModel) Create(data *contact) (*contact, error) {
	result := c.DB.Create(&data)
	if err := result.Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, fmt.Errorf("%v (%w)", err, utils.ErrorDuplicateEntry)
		}

		return nil, fmt.Errorf("Create(data *contact) -> %w", err)
	}
	return data, nil
}

// A method to delete a contact from DB
// It any error encountered.
func (c *contactModel) DeleteById(id uint) error {
	if err := c.DB.Delete(&contact{}, id).Error; err != nil {
		return fmt.Errorf("Create(data *contact) -> %w", err)
	}
	return nil
}

// A method to delete a contact from DB
// It any error encountered.
func (c *contactModel) UpdateById(id uint, data *contact) (*contact, error) {
	dataToBeUpdated, err := c.GetById(id)
	if err != nil {
		return nil, err
	}

	dataToBeUpdated.Name = data.Name
	dataToBeUpdated.Email = data.Email

	if err := c.DB.Save(&dataToBeUpdated).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key value") {
			return nil, fmt.Errorf("%v (%w)", err, utils.ErrorDuplicateEntry)
		}

		return nil, fmt.Errorf("UpdateById(id uint, data *contact) -> %w", err)
	}

	return dataToBeUpdated, nil
}

// A struct that provies mock data of contact model.
// It implements contactLinker interface.
type mockContactModel struct {
	contacts []contact
}

// Add two entries of contact data as initializer
func (m *mockContactModel) Init() error {
	m.contacts = make([]contact, 0)
	m.contacts = append(m.contacts, contactSeedData...)
	return nil
}

// Return all contact data of mockContactModel
func (m *mockContactModel) GetAll() (*[]contact, error) {
	return &m.contacts, nil
}

// Return a contact data of mockContactModel by its ID
func (m *mockContactModel) GetById(id uint) (*contact, error) {
	for _, c := range m.contacts {
		if c.ID == id {
			return &c, nil
		}
	}

	return nil, utils.ErrorNotFound
}

// Mock method that add a new entry with prior checking on duplicate key or code
func (m *mockContactModel) Create(data *contact) (*contact, error) {
	var maxID uint = 0
	for _, c := range m.contacts {
		if c.ID > maxID {
			maxID = c.ID
		}
	}

	if data.ID == 0 {
		data.ID = maxID + 1
	}

	m.contacts = append(m.contacts, *data)
	return data, nil
}

// Mock method to delete a contact by its ID
func (m *mockContactModel) DeleteById(id uint) error {
	newData := []contact{}
	for _, c := range m.contacts {
		if c.ID != id {
			newData = append(newData, c)
		}
	}

	m.contacts = newData
	return nil
}

// Mock method to update a contact by its ID
func (m *mockContactModel) UpdateById(id uint, data *contact) (*contact, error) {
	isFound := false
	index := 0
	for i, c := range m.contacts {
		if c.ID == id {
			m.contacts[i].Name = data.Name
			m.contacts[i].Email = data.Email
			isFound = true
			index = i
			break
		}
	}

	if !isFound {
		return nil, utils.ErrorNotFound
	}

	return &m.contacts[index], nil
}

func NewContact() *contact {
	return &contact{}
}
