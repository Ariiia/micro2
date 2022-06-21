package repository

import "context"

// type Note struct {
// 	id int `json:"id"`
// 	body string `json:"body"`
// }

type Visits interface {
	Inc(ctx context.Context) error
	Get(ctx context.Context) (int, error)
}

type Notes interface {
	GetNote(ctx context.Context, id int) (string, error)
	MakeNote(ctx context.Context, body string) (int, error)
	ChangeNote(ctx context.Context) (int, error)
	DeleteNote(ctx context.Context) (int, error)
	// GetAll(ctx context.Context)  ([]Note, error)
}
