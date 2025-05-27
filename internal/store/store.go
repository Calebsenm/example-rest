package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)


var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Projects interface {
		Create(context.Context , *Project) error
		GetAlls(context.Context) (*[]Project, error)
	}
	Participants interface {
		Create(context.Context , *Participants) error
		GetAlls(context.Context) (*[]Participants, error)
	}

	Assignment interface{ 
		AssignProject(context.Context , *Assignment) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Projects:       &ProjectsStore{db},
		Participants: &ParticipantsStore{db},
		Assignment:  &AssignmentStore{db},
	}
}
