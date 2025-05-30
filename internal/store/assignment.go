package store

import (
	"context"
	"database/sql"
)

type Assignment struct {
	ProjectID     int `json:"project_id"`
	ParticipantID int `json:"participant_id"`
}

type AssignmentDetail struct {
	ProjectID           int    `json:"project_id"`
	ProjectName         string `json:"project_name"`
	ParticipantID       int    `json:"participant_id"`
	ParticipantName     string `json:"participant_name"`
	ParticipantLastName string `json:"participant_last_name"`
}

type AssignmentStore struct {
	db *sql.DB
}

func (a *AssignmentStore) AssignProject(ctx context.Context, assig *Assignment) error {
	query := `
		INSERT INTO project_assignments 
			(project_id, participant_id) 
			VALUES (?, ?)`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := a.db.ExecContext(
		ctx,
		query,
		assig.ProjectID,
		assig.ParticipantID,
	)

	if err != nil {
		return err
	}

	return nil
}

func (a *AssignmentStore) GetAlls(ctx context.Context) (*[]AssignmentDetail, error) {
	query := `
		SELECT
			projects.project_id, projects.name,
			participants.identification, participants.first_name , participants.last_name
		FROM
			project_assignments 
		JOIN projects  
			ON project_assignments.project_id = projects.project_id
		JOIN participants  
			ON project_assignments.participant_id = participants.identification
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := a.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	assignments := []AssignmentDetail{}
	for rows.Next() {
		var a AssignmentDetail
		if err := rows.Scan(&a.ProjectID, &a.ProjectName, &a.ParticipantID, &a.ParticipantName ,&a.ParticipantLastName); err != nil {
			return nil, err
		}
		assignments = append(assignments, a)
	}

	return &assignments, nil
}
