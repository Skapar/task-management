package repository

import (
	"database/sql"
	"errors"

	"github.com/Skapar/task-management/internal/model"
)

type UserRepository struct {
    DB *sql.DB
}

func (r *UserRepository) GetAll() ([]model.User, error) {
    rows, err := r.DB.Query("SELECT id, name, email, role, created_at FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []model.User
    for rows.Next() {
        var user model.User
        if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt); err != nil {
            return nil, err
        }
        users = append(users, user)
    }
    return users, nil
}

func (r *UserRepository) GetByID(id int) (*model.User, error) {
    row := r.DB.QueryRow("SELECT id, name, email, role, created_at FROM users WHERE id = $1", id)
    var user model.User
    if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Role, &user.CreatedAt); err != nil {
        if err == sql.ErrNoRows {
            return nil, errors.New("user not found")
        }
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) Create(user *model.User) error {
    err := r.DB.QueryRow(
        "INSERT INTO users (name, email, role, created_at) VALUES ($1, $2, $3, $4) RETURNING id",
        user.Name, user.Email, user.Role, user.CreatedAt,
    ).Scan(&user.ID)
    return err
}

func (r *UserRepository) Update(user *model.User) error {
    _, err := r.DB.Exec(
        "UPDATE users SET name = $1, email = $2, role = $3 WHERE id = $4",
        user.Name, user.Email, user.Role, user.ID,
    )
    return err
}

func (r *UserRepository) Delete(id int) error {
    _, err := r.DB.Exec("DELETE FROM users WHERE id = $1", id)
    return err
}
