package handlers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

// TestMainPageHandler checks if the main page is returned successfully.
func TestMainPageHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(MainPageHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("MainPageHandler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := "text/html"
    if contentType := rr.Header().Get("Content-Type"); !strings.Contains(contentType, expected) {
        t.Errorf("MainPageHandler returned wrong content type: got %v want %v", contentType, expected)
    }
}

// TestProjectHighlightHandler verifies that the project highlight section is rendered correctly.
func TestProjectHighlightHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/highlights", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(ProjectHighlightHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("ProjectHighlightHandler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Example check for content within the response body
    // Assuming the response contains specific text related to project highlights
    if !strings.Contains(rr.Body.String(), "Project Highlights") {
        t.Errorf("ProjectHighlightHandler returned unexpected body: body does not contain 'Project Highlights'")
    }
}

// TestCVDownloadHandler ensures that the CV can be downloaded successfully.
func TestCVDownloadHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/cv", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(CVDownloadHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("CVDownloadHandler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    cd := rr.Header().Get("Content-Disposition")
    if !strings.Contains(cd, "attachment; filename=") {
        t.Errorf("CVDownloadHandler returned unexpected Content-Disposition: %v", cd)
    }
}

// TestContactHandler confirms that the contact information is rendered correctly.
func TestContactHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/contact", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(ContactHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("ContactHandler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    // Assuming the response should contain contact information
    if !strings.Contains(rr.Body.String(), "Contact Information") {
        t.Errorf("ContactHandler returned unexpected body: body does not contain 'Contact Information'")
    }
}

// TestBlogLinkHandler checks if the blog link redirects correctly.
func TestBlogLinkHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/blog", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(BlogLinkHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusFound {
        t.Errorf("BlogLinkHandler returned wrong status code: got %v want %v", status, http.StatusFound)
    }

    // Check for redirection URL
    location := rr.Header().Get("Location")
    expectedLocation := "http://example.com/blog" // Assuming example.com is where the blog is hosted
    if location != expectedLocation {
        t.Errorf("BlogLinkHandler redirected to wrong location: got %v want %v", location, expectedLocation)
    }
}