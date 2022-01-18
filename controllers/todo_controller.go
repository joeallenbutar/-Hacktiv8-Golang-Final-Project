package controllers

import (
	"Final-Project/database"
	"Final-Project/models"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/swaggo/swag/example/celler/httputil"
)

type APIEnv struct {
	DB *gorm.DB
}

// GetToDo godoc
// @summary Get ToDo
// @Description Get All ToDo's
// @Tags todos
// @Produce json
// @Success 200 {object} models.Todo
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /todos [get]
func (a *APIEnv) GetToDos(c *gin.Context) {
	todos, err := database.GetToDos(a.DB)
	if err != nil {
		fmt.Println(err)
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, todos)
}

// CreateToDo godoc
// @Tags todos
// @Description Create toDo
// @ID create-todo
// @Accept json
// @Produce json
// @Param RequestBody body models.TodoPayload true "request body json"
// @Success 201 {object} models.Todo
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /todos [post]
func (a *APIEnv) CreateToDo(c *gin.Context) {
	todo := models.Todo{}
	err := c.BindJSON(&todo)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	if err := a.DB.Create(&todo).Error; err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

// GetToDo godoc
// @Tags todos
// @Description Get toDo
// @ID get-todo
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} models.Todo
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /todos/{id} [get]
func (a *APIEnv) GetToDo(c *gin.Context) {
	id := c.Params.ByName("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
	}

	todo, exists, err := database.GetToDoByID(i, a.DB)
	if err != nil {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}

	if !exists {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, todo)
}

// UpdateToDo godoc
// @Tags todos
// @Description Update ToDo
// @ID update-todo
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param RequestBody body models.TodoPayload true "request body json"
// @Success 200 {object} models.Todo
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /todos/{id} [put]
func (a *APIEnv) UpdateToDo(c *gin.Context) {

	id := c.Params.ByName("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
	}

	_, exists, err := database.GetToDoByID(i, a.DB)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}

	updatedToDo := models.Todo{}
	err = c.BindJSON(&updatedToDo)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	updatedToDo.ID = uint(i)

	if err := database.UpdateToDo(a.DB, &updatedToDo); err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	a.GetToDo(c)
}

// DeleteToDo godoc
// @Tags todos
// @Description Delete ToDo
// @ID delete-todo
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 204
// @Failure 400 {object} httputil.HTTPError
// @Failure 404 {object} httputil.HTTPError
// @Failure 500 {object} httputil.HTTPError
// @Router /todos/{id} [delete]
func (a *APIEnv) DeleteToDo(c *gin.Context) {
	id := c.Params.ByName("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Fatal(err)
	}

	_, exists, err := database.GetToDoByID(i, a.DB)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	if !exists {
		httputil.NewError(c, http.StatusNotFound, err)
		return
	}

	err = database.DeleteToDo(id, a.DB)
	if err != nil {
		httputil.NewError(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "record deleted successfully")
}
