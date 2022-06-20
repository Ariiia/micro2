package postgres

import (
	"context"

	"github.com/jackc/pgx/v4"
)

type Repository struct {
	db *pgx.Conn
}

func New(db *pgx.Conn) *Repository {
	return &Repository{
		db: db,
	}
}

type Note struct {
	body	string	`json:"body"`
}

func (r *Repository) getNote(ctx context.Context) ( error){
	//addd
}

func (r *Repository) makeNote(ctx context.Context) ( error){
	note := Note{}
	err := ctx.BodyParser(&note)
	if err != nil {
		return err
	}
	err := r.db.Create(&book).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) changeNote(ctx context.Context) ( error){
	//addd

}

func (r *Repository) deleteNote(ctx context.Context) ( error){
	//addd

}


func (r *Repository) Get(ctx context.Context) (int, error) {
	rows, err := r.db.Query(ctx, "SELECT counter FROM visits")
	if err != nil {
		return 0, err
	}
	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, rows.Err()
}

func (r *Repository) Inc(ctx context.Context) error {
	current, err := r.Get(ctx)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(ctx, "UPDATE visits SET counter = $1", current+1)
	return err
}
