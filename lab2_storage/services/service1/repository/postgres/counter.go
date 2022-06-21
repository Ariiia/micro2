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
	id   int    `json:"id"`
	body string `json:"body"`
}
type Notes interface {
	GetAll(ctx context.Context) ([]Note, error)
	GetNote(ctx context.Context, id int) (string, error)
	MakeNote(ctx context.Context, body string) (int, error)
	ChangeNote(ctx context.Context) (int, error)
	DeleteNote(ctx context.Context) (int, error)
}

func (r *Repository) GetAll(ctx context.Context) ([]Note, error) {
	var notes []Note
	query := `select id, body from notes;`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return notes, err
	}
	defer rows.Close()

	for rows.Next() {
		var ID int
		var BODY string
		err = rows.Scan(&ID, &BODY)
		if err != nil {
			return notes, err
		}

		note := Note{id: ID, body: BODY}
		notes = append(notes, note)
	}
	return notes, nil
}

func (r *Repository) GetNote(ctx context.Context, id int) (string, error) {
	var note string
	err := r.db.QueryRow(ctx, "SELECT body FROM notes WHERE id = $1", id).Scan(&note)

	if err != nil {
		return "There has been some kind of an error, look: " + err.Error(), err
	}

	return note, nil
}

func (r *Repository) MakeNote(ctx context.Context, body string) (int, error) {

	_, err := r.db.Exec(context.Background(),
		"insert into notes(body) values($1)", body)

	return 0, err
	/*note := Note{}
	err := ctx.BodyParser(&note)
	if err != nil {
		return err
	}
	err := r.db.Create(&book).Error
	if err != nil {
		return err
	}
	return nil*/
}

func (r *Repository) ChangeNote(ctx context.Context) (int, error) {
	//addd
	panic("not imnplemented")
}

func (r *Repository) DeleteNote(ctx context.Context) (int, error) {
	//addd
	panic("not imnplemented")

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
