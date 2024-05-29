package handlers

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

var (
	list  = flag.Bool("list", false, "List tests")
	count = flag.Int("count", 1, "Run tests `n` times")
	json  = flag.Bool("json", false, "Output test results in JSON format")
	cpu   = flag.String("cpu", "", "Specify CPU `value`")
	race  = flag.Bool("race", false, "Enable data race detection")
)

func TestMain(m *testing.M) {
	flag.Parse()

	if *list {
		testing.Init()
		testing.Main(func(pat, str string) (bool, error) { return true, nil }, nil, nil, nil)
		return
	}

	if *race {
		testing.Init()
		testing.Main(func(pat, str string) (bool, error) { return true, nil }, nil, nil, nil)
		return
	}

	if *cpu != "" {
		if err := os.Setenv("GOMAXPROCS", *cpu); err != nil {
			println("Failed to set GOMAXPROCS:", err)
			return
		}
	}

	if *json {
		testing.Init()
		testing.Main(func(pat, str string) (bool, error) { return true, nil }, nil, nil, nil)
		return
	}

	m.Run()
}

func TestHomeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HomeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestGenerateHandlerWithCount(t *testing.T) {
	for i := 0; i < *count; i++ {
		TestGenerateHandler(t)
	}
}

func TestGenerateHandler(t *testing.T) {
	data := url.Values{}
	data.Set("input", "test")
	data.Set("stylename", "standard")

	req, err := http.NewRequest("POST", "/generate", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GenerateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestGenerateHandlerOnlyAsciiCharts(t *testing.T) {
	data := url.Values{}
	data.Set("input", "абай")
	data.Set("stylename", "standard")

	req, err := http.NewRequest("POST", "/generate", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GenerateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for non-ascii input: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestGenerateHandlerInvalidStyle(t *testing.T) {
	data := url.Values{}
	data.Set("input", "hello")
	data.Set("stylename", "invalid-style")

	req, err := http.NewRequest("POST", "/generate", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GenerateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code for invalid style: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestGenerateHandlerContentType(t *testing.T) {
	req, err := http.NewRequest("POST", "/generate", strings.NewReader(`{"input": "hello", "stylename": "standard"}`))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GenerateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for incorrect content type: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestHomeHandlerInvalidPath(t *testing.T) {
	req, err := http.NewRequest("GET", "/nonexistent", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HomeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code for invalid path: got %v want %v", status, http.StatusNotFound)
	}
}

func TestGenerateHandlerInvalidMethod(t *testing.T) {
	req, err := http.NewRequest("GET", "/generate", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GenerateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code for invalid method: got %v want %v", status, http.StatusMethodNotAllowed)
	}

}

func TestGenerateHandlerEmptyInput(t *testing.T) {
	data := url.Values{}
	data.Set("input", "")
	data.Set("stylename", "standard")

	req, err := http.NewRequest("POST", "/generate", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GenerateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code for empty input: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestGenerateHandlerCaseInsensitivity(t *testing.T) {
	data := url.Values{}
	data.Set("input", "HELLO")
	data.Set("stylename", "standard")

	req, err := http.NewRequest("POST", "/generate", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GenerateHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code for case insensitive input: got %v want %v", status, http.StatusOK)
	}
}

func TestHomeHandlerWithQueryParams(t *testing.T) {
	req, err := http.NewRequest("GET", "/?unused=param", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HomeHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code with query parameters: got %v want %v", status, http.StatusOK)
	}
}

func TestGenerateHandlerXSSAttempt(t *testing.T) {
	data := url.Values{}
	data.Set("input", "<script>alert('xss')</script>")
	data.Set("stylename", "standard")

	req, err := http.NewRequest("POST", "/generate", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GenerateHandler)

	handler.ServeHTTP(rr, req)

}
