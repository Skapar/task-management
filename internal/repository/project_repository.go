package repository

import (
	"database/sql"
	"errors"

	"github.com/Skapar/task-management/internal/model"
)

type ProjectRepository struct {
    DB *sql.DB
}

func (r *ProjectRepository) GetAll() ([]model.Project, error) {
    rows, err := r.DB.Query("SELECT id, title, description, start_date, end_date, manager_id FROM projects")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var projects []model.Project
    for rows.Next() {
        var project model.Project
        if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID); err != nil {
            return nil, err
        }
        projects = append(projects, project)
    }
    return projects, nil
}

func (r *ProjectRepository) GetByID(id int) (*model.Project, error) {
    row := r.DB.QueryRow("SELECT id, title, description, start_date, end_date, manager_id FROM projects WHERE id = $1", id)
    var project model.Project
    if err := row.Scan(&project.ID, &project.Title, &project.Description, &project.StartDate, &project.EndDate, &project.ManagerID); err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("project not found")
        }
        return nil, err
    }
    return &project, nil
}

func (r *ProjectRepository) Create(project *model.Project) error {
    err := r.DB.QueryRow(
        "INSERT INTO projects (title, description, start_date, end_date, manager_id) VALUES ($1, $2, $3, $4, $5) RETURNING id",
        project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID,
    ).Scan(&project.ID)
    return err
}

func (r *ProjectRepository) Update(project *model.Project) error {
    _, err := r.DB.Exec(
        "UPDATE projects SET title = $1, description = $2, start_date = $3, end_date = $4, manager_id = $5 WHERE id = $6",
        project.Title, project.Description, project.StartDate, project.EndDate, project.ManagerID, project.ID,
    )
    return err
}

func (r *ProjectRepository) Delete(id int) error {
    _, err := r.DB.Exec("DELETE FROM projects WHERE id = $1", id)
    return err
}
