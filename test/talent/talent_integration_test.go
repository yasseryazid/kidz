package talent_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yasseryazid/boilerpart/internal/infra/repository/talent/memory"
	"github.com/yasseryazid/boilerpart/internal/transport/http/talent"
	usecase "github.com/yasseryazid/boilerpart/internal/usecase/talent"
)

type talentPayload struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	Title     string   `json:"title"`
	Skills    []string `json:"skills"`
	CreatedAt string   `json:"created_at"`
}

type talentResponse struct {
	Data talentPayload `json:"data"`
}

type talentListResponse struct {
	Data []talentPayload `json:"data"`
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	r := gin.New()
	repo := memory.NewTalentRepository()
	svc := usecase.NewService(repo)
	talent.RegisterRoutes(r, svc)

	return r
}

func performRequest(t *testing.T, r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	t.Helper()

	req, err := http.NewRequest(method, path, bytes.NewReader(body))
	if err != nil {
		t.Fatalf("failed to build request: %v", err)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	return rec
}

func TestTalentLifecycleIntegration(t *testing.T) {
	router := setupRouter()

	createBody := []byte(`{"name":"Ada Lovelace","title":"Engineer","skills":["Go","DDD"]}`)
	createRec := performRequest(t, router, http.MethodPost, "/talents", createBody)
	if createRec.Code != http.StatusCreated {
		t.Fatalf("expected status %d, got %d", http.StatusCreated, createRec.Code)
	}

	var created talentResponse
	if err := json.Unmarshal(createRec.Body.Bytes(), &created); err != nil {
		t.Fatalf("failed to decode create response: %v", err)
	}

	if created.Data.ID == "" {
		t.Fatalf("expected talent id to be set")
	}
	if created.Data.Name != "Ada Lovelace" {
		t.Fatalf("expected name to be Ada Lovelace, got %s", created.Data.Name)
	}
	if created.Data.Title != "Engineer" {
		t.Fatalf("expected title to be Engineer, got %s", created.Data.Title)
	}
	if len(created.Data.Skills) != 2 {
		t.Fatalf("expected 2 skills, got %d", len(created.Data.Skills))
	}

	getRec := performRequest(t, router, http.MethodGet, "/talents/"+created.Data.ID, nil)
	if getRec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, getRec.Code)
	}

	var fetched talentResponse
	if err := json.Unmarshal(getRec.Body.Bytes(), &fetched); err != nil {
		t.Fatalf("failed to decode get response: %v", err)
	}
	if fetched.Data.ID != created.Data.ID {
		t.Fatalf("expected id %s, got %s", created.Data.ID, fetched.Data.ID)
	}

	listRec := performRequest(t, router, http.MethodGet, "/talents", nil)
	if listRec.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, listRec.Code)
	}

	var listed talentListResponse
	if err := json.Unmarshal(listRec.Body.Bytes(), &listed); err != nil {
		t.Fatalf("failed to decode list response: %v", err)
	}
	if len(listed.Data) != 1 {
		t.Fatalf("expected 1 talent in list, got %d", len(listed.Data))
	}
	if listed.Data[0].ID != created.Data.ID {
		t.Fatalf("expected listed id %s, got %s", created.Data.ID, listed.Data[0].ID)
	}
}

func TestTalentNotFoundIntegration(t *testing.T) {
	router := setupRouter()

	rec := performRequest(t, router, http.MethodGet, "/talents/999", nil)
	if rec.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, rec.Code)
	}
}

func TestTalentValidationIntegration(t *testing.T) {
	router := setupRouter()

	rec := performRequest(t, router, http.MethodPost, "/talents", []byte(`{"title":"Engineer"}`))
	if rec.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, rec.Code)
	}
}
