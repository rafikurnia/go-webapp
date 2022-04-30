package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"github.com/rafikurnia/go-webapp/models"
	"github.com/rafikurnia/go-webapp/utils"
)

// This API struct will contain methods that are used by all API resources. It is designed to allow
// code reuse
type api struct{}

// GET {resources}
func (a *api) getAll(c *gin.Context, resources string) {
	mapData := utils.ResourcesMap[resources]

	r := reflect.ValueOf(models.DB)
	f := reflect.Indirect(r).FieldByName(mapData["dbFieldName"])
	m := f.MethodByName("GetAll")

	ret := m.Call([]reflect.Value{})

	data := ret[0].Interface()
	err := ret[1].Interface()

	if err != nil {
		log.Printf("getAll(%v) -> %v", resources, err)
		utils.Throws(c, http.StatusInternalServerError, err.(error))
		return
	}

	c.IndentedJSON(http.StatusOK, data)
}

// GET {resources}/:id
func (a *api) getById(c *gin.Context, resources string) {
	mapData := utils.ResourcesMap[resources]

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("getById(%v) -> %v", resources, err)
		utils.Throws(c, http.StatusBadRequest, errors.New(
			fmt.Sprintf("'%v' is not a valid ID.", c.Param("id"))))
		return
	}

	r := reflect.ValueOf(models.DB)
	f := reflect.Indirect(r).FieldByName(mapData["dbFieldName"])
	m := f.MethodByName("GetById")

	ret := m.Call([]reflect.Value{reflect.ValueOf(uint(id))})

	data := ret[0].Interface()
	errs := ret[1].Interface()

	if errs != nil {
		log.Printf("getById(%v) -> %v", resources, errs.(error))

		if errors.Is(errs.(error), utils.ErrorNotFound) {
			utils.Throws(c, http.StatusNotFound, errors.New(
				fmt.Sprintf("A %v with ID=%d is not found.",
					mapData["singleton"], id)))
			return
		}

		utils.Throws(c, http.StatusInternalServerError, errs.(error))
		return
	}
	c.IndentedJSON(http.StatusOK, data)
}

// POST {resources}
func (a *api) create(c *gin.Context, resources string) {
	mapData := utils.ResourcesMap[resources]

	var model interface{}

	if err := c.ShouldBindBodyWith(&model, binding.JSON); err != nil {
		log.Printf("create(%v) -> %v", resources, err)
		utils.Throws(c, http.StatusBadRequest, err)
		return
	}

	r := reflect.ValueOf(models.DB)
	f := reflect.Indirect(r).FieldByName(mapData["dbFieldName"])
	m := f.MethodByName("Create")

	ret := []reflect.Value{}

	switch resources {
	case "contacts":
		d := models.NewContact()
		jsonStr, _ := json.Marshal(model)
		json.Unmarshal(jsonStr, d)
		ret = m.Call([]reflect.Value{reflect.ValueOf(d)})
	}

	data := ret[0].Interface()
	errs := ret[1].Interface()

	if errs != nil {
		log.Printf("create(%v) -> %v", resources, errs.(error))
		if errors.Is(errors.Unwrap(errs.(error)), utils.ErrorDuplicateEntry) {
			utils.Throws(c, http.StatusBadRequest,
				errors.New("Duplicate detected on values which supposed to be unique."))
			return
		}

		utils.Throws(c, http.StatusInternalServerError, errs.(error))
		return
	}

	c.IndentedJSON(http.StatusCreated, data)
}

// DELETE {resources}/:id
func (a *api) deleteById(c *gin.Context, resources string) {
	mapData := utils.ResourcesMap[resources]

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("deleteById(%v) -> %v", resources, err)
		utils.Throws(c, http.StatusBadRequest, errors.New(
			fmt.Sprintf("'%v' is not a valid ID.", c.Param("id"))))
		return
	}

	r := reflect.ValueOf(models.DB)
	f := reflect.Indirect(r).FieldByName(mapData["dbFieldName"])
	m := f.MethodByName("GetById")

	ret := m.Call([]reflect.Value{reflect.ValueOf(uint(id))})

	errs := ret[1].Interface()
	if errs != nil {
		log.Printf("deleteById(%v) -> %v", resources, errs.(error))

		if errors.Is(errs.(error), utils.ErrorNotFound) {
			utils.Throws(c, http.StatusNotFound, errors.New(
				fmt.Sprintf("A %v with ID=%d is not found.",
					mapData["singleton"], id)))
			return
		}

		utils.Throws(c, http.StatusInternalServerError, errs.(error))
		return
	}

	m = f.MethodByName("DeleteById")

	ret = m.Call([]reflect.Value{reflect.ValueOf(uint(id))})

	if errs = ret[0].Interface(); err != nil {
		log.Printf("deleteById(%v) -> %v", resources, err.(error))
		utils.Throws(c, http.StatusInternalServerError, err.(error))
		return
	}
	c.IndentedJSON(http.StatusNoContent, nil)
}

// PUT {resources}/:id
func (a *api) updateById(c *gin.Context, resources string) {
	mapData := utils.ResourcesMap[resources]

	id, err := strconv.ParseUint(c.Param("id"), 10, 0)
	if err != nil {
		log.Printf("updateById(%v) -> %v", resources, err)
		utils.Throws(c, http.StatusBadRequest, errors.New(
			fmt.Sprintf("'%v' is not a valid ID.", c.Param("id"))))
		return
	}

	var model interface{}

	if err := c.ShouldBindBodyWith(&model, binding.JSON); err != nil {
		log.Printf("updateById(%v) -> %v", resources, err)
		utils.Throws(c, http.StatusBadRequest, err)
		return
	}

	r := reflect.ValueOf(models.DB)
	f := reflect.Indirect(r).FieldByName(mapData["dbFieldName"])
	m := f.MethodByName("UpdateById")

	ret := []reflect.Value{}

	switch resources {
	case "contacts":
		d := models.NewContact()
		jsonStr, _ := json.Marshal(model)
		json.Unmarshal(jsonStr, d)
		ret = m.Call([]reflect.Value{reflect.ValueOf(uint(id)), reflect.ValueOf(d)})
	}

	data := ret[0].Interface()
	errs := ret[1].Interface()

	if errs != nil {
		log.Printf("create(%v) -> %v", resources, errs.(error))

		if errors.Is(errs.(error), utils.ErrorNotFound) {
			utils.Throws(c, http.StatusNotFound, errors.New(
				fmt.Sprintf("A %v with ID=%d is not found.",
					mapData["singleton"], id)))
			return
		}

		if errors.Is(errors.Unwrap(errs.(error)), utils.ErrorDuplicateEntry) {
			utils.Throws(c, http.StatusBadRequest,
				errors.New("Duplicate detected on values which supposed to be unique."))
			return
		}

		utils.Throws(c, http.StatusInternalServerError, errs.(error))
		return
	}

	c.IndentedJSON(http.StatusOK, data)
}
