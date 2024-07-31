package repository

import (
	"database/sql"
	"errors"

	"github.com/Skapar/task-management/internal/model"
)

type TaskRepository struct {
    DB *sql.DB
}

func (r *TaskRepository) GetAll() ([]model.Task, error) {
    rows, err := r.DB.Query("SELECT id, title, description, priority, status, assignee_id, project_id, created_at, completed_at FROM tasks")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tasks []model.Task
    for rows.Next() {
        var task model.Task
        if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.AssigneeID, &task.ProjectID, &task.CreatedAt, &task.CompletedAt); err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }
    return tasks, nil
}

func (r *TaskRepository) GetByID(id int) (*model.Task, error) {
    row := r.DB.QueryRow("SELECT id, title, description, priority, status, assignee_id, project_id, created_at, completed_at FROM tasks WHERE id = $1", id)
    var task model.Task
    if err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Priority, &task.Status, &task.AssigneeID, &task.ProjectID, &task.CreatedAt, &task.CompletedAt); err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("task not found")
        }
        return nil, err
    }
    return &task, nil
}

func (r *TaskRepository) Create(task *model.Task) error {
    err := r.DB.QueryRow(
        "INSERT INTO tasks (title, description, priority, status, assignee_id, project_id, created_at, completed_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
        task.Title, task.Description, task.Priority, task.Status, task.AssigneeID, task.ProjectID, task.CreatedAt, task.CompletedAt,
    ).Scan(&task.ID)
    return err
}

func (r *TaskRepository) Update(task *model.Task) error {
    _, err := r.DB.Exec(
        "UPDATE tasks SET title = $1, description = $2, priority = $3, status = $4, assignee_id = $5, project_id = $6, completed_at = $7 WHERE id = $8",
        task.Title, task.Description, task.Priority, task.Status, task.AssigneeID, task.ProjectID, task.CompletedAt, task.ID,
    )
    return err
}

func (r *TaskRepository) Delete(id int) error {
    _, err := r.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
    return err
}
