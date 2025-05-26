package store

import (
	"context"
	"database/sql"
	"fmt"
)


type Participants struct {
	Identification string `json:"identificacion"`
	First_name     string `json:"firs_name"`
	Last_name      string `json:"last_name"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
}

type ParticipantsStore struct {
	db *sql.DB
}

func (p *ParticipantsStore) Create(ctx context.Context, participant *Participants) error {
	query := `
		INSERT INTO participants (identification, first_name, last_name, email, phone)
		VALUES (?, ?, ?, ?, ?)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	result, err := p.db.ExecContext(ctx, query,
		participant.Identification,
		participant.First_name,
		participant.Last_name,
		participant.Email,
		participant.Phone,
	)
	if err != nil {
		return fmt.Errorf("failed to insert participant: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}
	
	participant.Identification = fmt.Sprintf("%d", id)

	return nil
}

func (p *ParticipantsStore) GetAlls(ctx context.Context) (*[]Participants, error) {
	query := `SELECT identification, first_name, last_name, email, phone FROM participants`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := p.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []Participants

	for rows.Next() {
		var participant Participants
		err := rows.Scan(
			&participant.Identification,
			&participant.First_name,
			&participant.Last_name,
			&participant.Email,
			&participant.Phone,
		)
		if err != nil {
			return nil, err
		}
		participants = append(participants, participant)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &participants, nil
}

