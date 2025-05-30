package store

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

type Project struct {
	ProjectID   int       `json:"project_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date" example:"2024-10-10T00:00:00Z"`
	EndDate     time.Time `json:"end_date"  example:"2024-10-10T00:00:00Z"`
	Value       float64   `json:"value"`
}

type ProjectsStore struct {
	db *sql.DB
}

func (p *ProjectsStore) Create(ctx context.Context, project *Project) error {
	query := `
		INSERT INTO projects (project_id, name, description, start_date, end_date, value)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := p.db.ExecContext(ctx, query,
		project.ProjectID,
		project.Name,
		project.Description,
		project.StartDate,
		project.EndDate,
		project.Value,
	)
	if err != nil {
		return fmt.Errorf("failed to insert project: %w", err)
	}

	return nil
}

func (p *ProjectsStore) GetAlls(ctx context.Context) (*[]Project, error) {
	query := `
		SELECT project_id, name, description, start_date, end_date, value FROM projects
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to query projects: %w", err)
	}
	defer rows.Close()

	var projects []Project

	for rows.Next() {
		var project Project
		err := rows.Scan(
			&project.ProjectID,
			&project.Name,
			&project.Description,
			&project.StartDate,
			&project.EndDate,
			&project.Value,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan project: %w", err)
		}
		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return &projects, nil
}

func (p *ProjectsStore) SearchByName(ctx context.Context, name string) (*[]Project, error) {
    query := `SELECT project_id, name, description, start_date, end_date, value 
              FROM projects WHERE name LIKE ?`

    rows, err := p.db.QueryContext(ctx, query, "%"+name+"%")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var projects []Project
    for rows.Next() {
        var proj Project
        err := rows.Scan(&proj.ProjectID, &proj.Name, &proj.Description, &proj.StartDate, &proj.EndDate, &proj.Value)
        if err != nil {
            return nil, err
        }
        projects = append(projects, proj)
    }

    return &projects, nil
}

