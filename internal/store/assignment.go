package store

import (
	"context"
	"database/sql"
)

type Assignment struct {
	ProjectID     int `json:"project_id"`
	ParticipantID int `json:"participant_id"`
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
