package app

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"todo-app/internal/config"
)

// TestTaskFlow drives the real handler: GET issues a CSRF token, a POST
// without it is rejected, a valid POST stores a task that GET then lists.
func TestTaskFlow(t *testing.T) {
	a, err := New(&config.Metadata{
		Project:  config.Project{Name: "todo-app", Kind: "app"},
		Database: config.Database{Driver: "sqlite", DSN: filepath.Join(t.TempDir(), "test.db")},
	})
	if err != nil {
		t.Fatal(err)
	}
	h := a.Handler()

	get := httptest.NewRecorder()
	h.ServeHTTP(get, httptest.NewRequest(http.MethodGet, "/", nil))
	if get.Code != http.StatusOK {
		t.Fatalf("GET / = %d", get.Code)
	}
	cookie := get.Result().Cookies()[0]
	token := regexp.MustCompile(`name="csrf_token" value="([^"]+)"`).FindStringSubmatch(get.Body.String())[1]

	post := func(form url.Values) *httptest.ResponseRecorder {
		req := httptest.NewRequest(http.MethodPost, "/tasks", strings.NewReader(form.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.AddCookie(cookie)
		rec := httptest.NewRecorder()
		h.ServeHTTP(rec, req)
		return rec
	}

	if rec := post(url.Values{"title": {"no token"}}); rec.Code != http.StatusForbidden {
		t.Fatalf("POST without CSRF token = %d, want 403", rec.Code)
	}
	if rec := post(url.Values{"csrf_token": {token}, "title": {"  "}}); !strings.Contains(rec.Header().Get("Location"), "error=") {
		t.Fatalf("blank title should redirect with an error, got %q", rec.Header().Get("Location"))
	}
	if rec := post(url.Values{"csrf_token": {token}, "title": {"Write the tutorial"}}); rec.Code != http.StatusSeeOther {
		t.Fatalf("valid POST = %d, want 303", rec.Code)
	}

	list := httptest.NewRecorder()
	h.ServeHTTP(list, httptest.NewRequest(http.MethodGet, "/", nil))
	if !strings.Contains(list.Body.String(), "Write the tutorial") {
		t.Fatal("new task not listed on GET /")
	}
}
