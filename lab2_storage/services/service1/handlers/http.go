package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitlab.com/kpi-lab/microservices-demo/services/service1/repository"
	"gitlab.com/kpi-lab/microservices-demo/services/service1/repository/postgres"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type NotesServer struct {
	db postgres.Notes
}

func NewNotesServer(db postgres.Notes) *NotesServer {
	return &NotesServer{
		db: db,
	}
}

type Server struct {
	db repository.Visits
}

func NewVisitsServer(db repository.Visits) *Server {
	return &Server{
		db: db,
	}
}

func (s *NotesServer) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	lines, err := s.db.GetAll(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(lines)

	}
}

func (s *NotesServer) GetNote(w http.ResponseWriter, r *http.Request) {

	//ADD CODE
	var err error

	keys, ok := r.URL.Query()["id"]
	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'id' is missing")
		err = errors.New("URL param ID is completely and utterly missing")
		return
	}

	// var n int
	log.Println("getting note(GET)")
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		}
	}()

	id, err := strconv.Atoi(keys[0])

	note, err := s.db.GetNote(r.Context(), id)

	_, err = w.Write([]byte("{note: " + note + "}"))
}

func (s *NotesServer) MakeNote(w http.ResponseWriter, r *http.Request) {
	var err error
	log.Println("making note(POST)")
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError) + " \n" + err.Error()))
		}
	}()

	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	_, err = s.db.MakeNote(r.Context(), string(body))
	if err != nil {
		log.Println("failed to create note: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Created note")
	_, err = w.Write([]byte(msg))
}

func (s *NotesServer) ChangeNote(w http.ResponseWriter, r *http.Request) {
	//ADD CODE
	var err error
	//var n int

	log.Println("editing note(PUT)")
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		}
	}()

	_, err = s.db.ChangeNote(r.Context())
}

func (s *NotesServer) DeleteNote(w http.ResponseWriter, r *http.Request) {
	//ADD CODE
	var err error
	//var n int
	log.Println("deleting note(DELETE)")
	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		}
	}()

	_, err = s.db.DeleteNote(r.Context())
}

func (s *Server) Ping(w http.ResponseWriter, r *http.Request) {
	var current int
	var err error
	log.Println("ping request")

	defer func() {
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
		}
	}()

	err = s.db.Inc(r.Context())
	if err != nil {
		log.Println("failed to increment visits: %w", err)
	}
	current, err = s.db.Get(r.Context())
	if err != nil {
		log.Println("failed to get visits: %w", err)
	}

	w.WriteHeader(http.StatusOK)
	msg := fmt.Sprintf("Service1 is healthy, visited: %d times", current)
	_, err = w.Write([]byte(msg))
}
