package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestChatClient checks if the client page is returned successfully.
func TestChatClient(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ChatClientHandler)
	handler.ServeHTTP(rr, req)
	//test status code
	bytes, _ := io.ReadAll(rr.Body)
	t.Logf(string(bytes))
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("MainPageHandler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//test content type
	expected := "text/html"
	if contentType := rr.Header().Get("Content-Type"); !strings.Contains(contentType, expected) {
		t.Errorf("MainPageHandler returned wrong content type: got %v want %v", contentType, expected)
	}
}
