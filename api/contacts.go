package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/rafikurnia/go-webapp/models"
	"github.com/rafikurnia/go-webapp/utils"
)

type contacts struct {
	*api
}

// List all contacts
// @Summary  Get a list of contacts
// @Tags     Contacts
// @Accept   json
// @Produce  json
// @Success  200               {object}  []models.Contact
// @Failure  500               {object}  utils.HTTPError
// @Router   /contacts         [get]
func (c *contacts) getAll(ctx *gin.Context) {
	c.api.getAll(ctx, "contacts")
}

// Get a contact with the specified ID
// @Summary  Returns a contact with the specified ID
// @Tags     Contacts
// @Accept   json
// @Produce  json
// @Param    id                path      uint  true  "The ID of the contact"  Format(uint)
// @Success  200               {object}  models.Contact
// @Failure  400               {object}  utils.HTTPError
// @Failure  404               {object}  utils.HTTPError
// @Failure  500               {object}  utils.HTTPError
// @Router   /contacts/{id}    [get]
func (c *contacts) getById(ctx *gin.Context) {
	c.api.getById(ctx, "contacts")
}

// Create a contact
// @Summary  Create a contact
// @Tags     Contacts
// @Accept   json
// @Produce  json
// @Param    contact           body      models.ContactInput  true  "Each field is required, otherwise it will be set as an empty string."
// @Success  201               {object}  models.Contact
// @Failure  400               {object}  utils.HTTPError
// @Failure  500               {object}  utils.HTTPError
// @Router   /contacts         [post]
func (c *contacts) create(ctx *gin.Context) {
	if err := ctx.ShouldBindBodyWith(models.NewContact(), binding.JSON); err != nil {
		log.Printf("create(contacts) -> %v", err)
		utils.Throws(ctx, http.StatusBadRequest, err)
		return
	}

	c.api.create(ctx, "contacts")
}

// Delete a contact by ID
// @Summary  Delete a contact with the specified ID
// @Tags     Contacts
// @Accept   json
// @Produce  json
// @Param    id                path      uint  true  "The ID of the contact"  Format(uint)
// @Success  204               {object}  nil
// @Failure  400               {object}  utils.HTTPError
// @Failure  404               {object}  utils.HTTPError
// @Failure  500               {object}  utils.HTTPError
// @Router   /contacts/{id}  [delete]
func (c *contacts) deleteById(ctx *gin.Context) {
	c.api.deleteById(ctx, "contacts")
}

// Update a contact by ID
// @Summary  Update a contact with the specified ID
// @Tags     Contacts
// @Accept   json
// @Produce  json
// @Param    id                path      uint  true  "The ID of the contact"  Format(uint)
// @Param    contact           body      models.ContactInput  true  "Each field is required, otherwise it will be set as an empty string."
// @Success  200               {object}  models.Contact
// @Failure  400               {object}  utils.HTTPError
// @Failure  404               {object}  utils.HTTPError
// @Failure  500               {object}  utils.HTTPError
// @Router   /contacts/{id}  [put]
func (c *contacts) updateById(ctx *gin.Context) {
	if err := ctx.ShouldBindBodyWith(models.NewContact(), binding.JSON); err != nil {
		log.Printf("updateById(contacts) -> %v", err)
		utils.Throws(ctx, http.StatusBadRequest, err)
		return
	}

	c.api.updateById(ctx, "contacts")
}
