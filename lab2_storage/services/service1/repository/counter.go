package repository

import "context"

type Visits interface {
	Inc(ctx context.Context) error
	Get(ctx context.Context) (int, error)
}


type Notes interface {
	//Inc(ctx context.Context) error
	getNote(ctx context.Context) (error)
	makeNote(ctx context.Context) (int, error)
	changeNote(ctx context.Context) (int, error)
	deleteNote(ctx context.Context) (int, error)
}
