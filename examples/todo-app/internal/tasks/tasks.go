// Package tasks holds the task store and the HTML form handlers.
package tasks

import (
	"context"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/krewire/krewire/pkg/orm"
	"github.com/krewire/krewire/pkg/validation"
	"github.com/krewire/krewire/pkg/web"

	"todo-app/internal/models"
)

// Store wraps the database with the queries the app needs.
type Store struct{ DB *orm.DB }

// List returns open tasks first, newest first.
func (s *Store) List(ctx context.Context) ([]models.Task, error) {
	var out []models.Task
	err := s.DB.Model(&models.Task{}).WithContext(ctx).
		OrderBy("done", "ASC").OrderBy("id", "DESC").
		FindAll(&out)
	return out, err
}

// Create inserts a task and fills in its ID.
func (s *Store) Create(ctx context.Context, t *models.Task) error {
	_, err := s.DB.Model(t).WithContext(ctx).Insert(t)
	return err
}

// Toggle flips the done flag of one task.
func (s *Store) Toggle(ctx context.Context, id int64) error {
	var t models.Task
	if err := s.DB.Model(&t).WithContext(ctx).Find(&t, id); err != nil {
		return err
	}
	t.Done = !t.Done
	_, err := s.DB.Model(&t).WithContext(ctx).Update(&t)
	return err
}

// Delete removes one task.
func (s *Store) Delete(ctx context.Context, id int64) error {
	_, err := s.DB.Model(&models.Task{}).WithContext(ctx).Delete(&models.Task{}, id)
	return err
}

// BoardData is the data the TaskBoard.kiw view renders.
type BoardData struct {
	Title   string
	Tasks   []models.Task
	Summary string // e.g. "2 open"
	Token   string // CSRF token for the hidden form fields
	Error   string
	Draft   string
}

// Board loads the data for GET /. Validation errors come back from the POST
// handler as query parameters (post/redirect/get).
func (s *Store) Board(r *http.Request) (any, error) {
	list, err := s.List(r.Context())
	if err != nil {
		return nil, err
	}
	remaining := 0
	for _, t := range list {
		if !t.Done {
			remaining++
		}
	}
	return BoardData{
		Title:   "Tasks",
		Tasks:   list,
		Summary: strconv.Itoa(remaining) + " open",
		Token:   web.CSRFFrom(r.Context()),
		Error:   r.URL.Query().Get("error"),
		Draft:   r.URL.Query().Get("title"),
	}, nil
}

// newTask is the form input, validated with pkg/validation struct tags.
type newTask struct {
	Title string `validate:"required,max=120"`
}

// CreateHandler handles POST /tasks.
func (s *Store) CreateHandler(w http.ResponseWriter, r *http.Request, _ web.Params) {
	in := newTask{Title: strings.TrimSpace(r.PostFormValue("title"))}
	if err := validation.Struct(&in); err != nil {
		msg := "Title is required."
		if in.Title != "" {
			msg = "Title must be 120 characters or fewer."
		}
		q := url.Values{"error": {msg}, "title": {in.Title}}
		web.Redirect(w, r, "/?"+q.Encode(), http.StatusSeeOther)
		return
	}
	t := &models.Task{Title: in.Title, CreatedAt: time.Now().UTC()}
	if err := s.Create(r.Context(), t); err != nil {
		web.Error(w, err)
		return
	}
	web.Redirect(w, r, "/", http.StatusSeeOther)
}

// ToggleHandler handles POST /tasks/{id}/toggle.
func (s *Store) ToggleHandler(w http.ResponseWriter, r *http.Request, p web.Params) {
	id, err := strconv.ParseInt(p["id"], 10, 64)
	if err != nil {
		web.Error(w, web.BadRequest("invalid task id"))
		return
	}
	if err := s.Toggle(r.Context(), id); err != nil {
		web.Error(w, err)
		return
	}
	web.Redirect(w, r, "/", http.StatusSeeOther)
}

// DeleteHandler handles POST /tasks/{id}/delete.
func (s *Store) DeleteHandler(w http.ResponseWriter, r *http.Request, p web.Params) {
	id, err := strconv.ParseInt(p["id"], 10, 64)
	if err != nil {
		web.Error(w, web.BadRequest("invalid task id"))
		return
	}
	if err := s.Delete(r.Context(), id); err != nil {
		web.Error(w, err)
		return
	}
	web.Redirect(w, r, "/", http.StatusSeeOther)
}
