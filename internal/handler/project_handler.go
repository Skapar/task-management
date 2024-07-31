package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Skapar/task-management/internal/model"
	"github.com/Skapar/task-management/internal/repository"
	"github.com/gorilla/mux"
)

type ProjectHandler struct {
    Repo *repository.ProjectRepository
}

func (h *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
    projects, err := h.Repo.GetAll()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(projects)
}

func (h *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid project ID", http.StatusBadRequest)
        return
    }

    project, err := h.Repo.GetByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
    var project model.Project
    if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    err := h.Repo.Create(&project)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid project ID", http.StatusBadRequest)
        return
    }

    var project model.Project
    if err := json.NewDecoder(r.Body).Decode(&project); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    project.ID = id

    err = h.Repo.Update(&project)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(project)
}

func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, err := strconv.Atoi(params["id"])
    if err != nil {
        http.Error(w, "Invalid project ID", http.StatusBadRequest)
        return
    }

    err = h.Repo.Delete(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
}
