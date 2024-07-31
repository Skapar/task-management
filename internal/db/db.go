package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    var err error

    dsn := os.Getenv("DATABASE_URL")
    if dsn == "" {
        dsn = "user=postgres password=yourpassword dbname=task_management sslmode=disable"
    }

    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to open a DB connection: %v", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatalf("Failed to ping DB: %v", err)
    }

    log.Println("Database connected successfully")

    runMigrations()
}

func CloseDB() {
    if err := DB.Close(); err != nil {
        log.Fatalf("Failed to close DB connection: %v", err)
    }
}

func tableExists(tableName string) bool {
    var exists bool
    query := `SELECT EXISTS (
        SELECT FROM information_schema.tables 
        WHERE table_schema = 'public' AND table_name = $1
    )`
    err := DB.QueryRow(query, tableName).Scan(&exists)
    if err != nil {
        log.Fatalf("Failed to check if table %s exists: %v", tableName, err)
    }
    return exists
}

func runMigrations() {
    migrationDir := "migrations"
    files, err := os.ReadDir(migrationDir)
    if err != nil {
        log.Fatalf("Failed to read migrations directory: %v", err)
    }

    for _, file := range files {
        if filepath.Ext(file.Name()) == ".sql" {
            tableName := getTableNameFromFile(file.Name())
            if tableExists(tableName) {
                log.Printf("Table %s already exists, skipping migration", tableName)
                continue
            }

            path := filepath.Join(migrationDir, file.Name())
            script, err := os.ReadFile(path)
            if err != nil {
                log.Fatalf("Failed to read migration file %s: %v", path, err)
            }

            _, err = DB.Exec(string(script))
            if err != nil {
                log.Fatalf("Failed to execute migration script %s: %v", path, err)
            }

            log.Printf("Successfully ran migration script: %s", path)
        }
    }
}

func getTableNameFromFile(filename string) string {
    parts := strings.Split(filename, "_")
    return parts[2]
}