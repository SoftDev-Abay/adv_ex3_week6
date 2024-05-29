package handlers

import (
	"ascii-art-web/ascii-art/logic"
	"ascii-art-web/ascii-art/validation"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Error struct {
	ErrorCode int
	ErrorMsg  string
}

type Page struct {
	InputText string
	StyleName string
	AsciiText string
}

const testMode = true

func getTemplatePath(relativePath string) string {
	var finalPath string

	if testMode {
		finalPath = filepath.Join("C:/Users/abaya/Desktop/Code/aleam/ascii-art-web-stylize/", relativePath)
	} else {
		dir, err := filepath.Abs(filepath.Dir("."))
		if err != nil {
			log.Fatalf("Error getting absolute directory path: %v", err)
		}
		finalPath = filepath.Join(dir, relativePath)
	}

	return finalPath
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	log.Println("HomeHandler called", r.Method)

	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	t, err := template.ParseFiles(getTemplatePath("client/templates/home.html"))
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func GenerateHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("GenerateHandler called", r.Method)

	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	// Check if content type is application/x-www-form-urlencoded
	if r.Header.Get("Content-Type") != "application/x-www-form-urlencoded" {
		log.Println("Invalid content type:", r.Header.Get("Content-Type"))
		ErrorHandler(w, r, http.StatusBadRequest, "Invalid content type")
		return
	}

	if err := r.ParseForm(); err != nil {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	t, err := template.ParseFiles(getTemplatePath("client/templates/home.html"))
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	inputText := r.FormValue("input")
	styleName := r.FormValue("stylename")

	if inputText == "" || styleName == "" {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return

	}

	if err := validation.ConsistsOnlyFromAsciiChars(inputText); err != nil {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if err := validation.CheckValidStyle(styleName); err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusBadRequest))
		return
	}

	pathToStyle := getTemplatePath("ascii-art/styles/" + styleName + ".txt")

	letters, err := logic.GetAsciiArtLetters(pathToStyle)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	asciiText := logic.ConvertWordsToAsciiArtWithNewLinesStr(inputText, letters)

	pageContent := Page{
		InputText: inputText,
		StyleName: styleName,
		AsciiText: asciiText,
	}

	err = t.Execute(w, pageContent)
	if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	}
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, code int, msg string) {
	t, err := template.ParseFiles(getTemplatePath("client/templates/error.html"))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(code)
	errors := Error{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
	err = t.Execute(w, errors)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
