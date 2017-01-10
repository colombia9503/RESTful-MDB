package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/colombia9503/RESTful-MDB/common"
	"github.com/colombia9503/RESTful-MDB/models"
	"github.com/gorilla/mux"
)

var Task = new(tasksController)

type tasksController struct{}

//insertar
func (tc *tasksController) Create(w http.ResponseWriter, r *http.Request) {
	var t models.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	task, err := models.Tasks.Create(t.Name, t.Desc)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusCreated)
}

//Buscar un elemento
func (tc *tasksController) FindOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	println("Find one: %s", id)
	task, err := models.Tasks.FindOne(id)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}

//buscar todos
func (tc *tasksController) FindAll(w http.ResponseWriter, r *http.Request) {
	tasks, err := models.Tasks.FindAll()
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(tasks)
	if err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonOk(w, res, http.StatusOK)
}

//Actualizar
func (tc *tasksController) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	println("Update: %s", id)
	var t models.Task

	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}

	if err := models.Tasks.Update(id, t.Name, t.Desc); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonStatus(w, http.StatusNoContent)
}

//Eliminar
func (tc *tasksController) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	println("Delete: %s", id)
	if err := models.Tasks.DeleteById(id); err != nil {
		common.JsonError(w, err, http.StatusBadRequest)
		return
	}
	common.JsonStatus(w, http.StatusNoContent)

}
