package stt

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DooomiT/rudi-go/client/assemblyai"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestSpeechToText(t *testing.T) {
	client := assemblyai.NewMock("https://some-url.com/some-id", nil, "abcd", nil, "I am a valid byte array", nil)
	requestBody := []byte(`{"audio": "SSBhbSBhIHZhbGlkIGJ5dGUgYXJyYXk="}`)
	reader := bytes.NewReader(requestBody)
	response := `{"text":"I am a valid byte array"}`
	r := SetRouter()
	r.POST("/stt", SpeechToText(client, 1))
	req, _ := http.NewRequest("POST", "/stt", reader)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, response, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSpeechToTextBadRequest(t *testing.T) {
	response := `{"Data":{"data":"invalid request"},"Message":"error","Status":400}`
	r := SetRouter()
	r.POST("/stt", SpeechToText(nil, 1))
	req, _ := http.NewRequest("POST", "/stt", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, response, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSpeechToTextUploadFileError(t *testing.T) {
	client := assemblyai.NewMock("", errors.New("upload file failed"), "", nil, "", nil)
	response := `{"Data":{"data":"upload file failed"},"Message":"error","Status":400}`
	requestBody := []byte(`{"audio": "SSBhbSBhIHZhbGlkIGJ5dGUgYXJyYXk="}`)
	reader := bytes.NewReader(requestBody)
	r := SetRouter()
	r.POST("/stt", SpeechToText(client, 1))
	req, _ := http.NewRequest("POST", "/stt", reader)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, response, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSpeechToTextTranscriptionError(t *testing.T) {
	client := assemblyai.NewMock("https://some-url.com/some-id", nil, "", errors.New("transcription failed"), "", nil)
	response := `{"Data":{"data":"transcription failed"},"Message":"error","Status":400}`
	requestBody := []byte(`{"audio": "SSBhbSBhIHZhbGlkIGJ5dGUgYXJyYXk="}`)
	reader := bytes.NewReader(requestBody)
	r := SetRouter()
	r.POST("/stt", SpeechToText(client, 1))
	req, _ := http.NewRequest("POST", "/stt", reader)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, response, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSpeechToTextPollError(t *testing.T) {
	client := assemblyai.NewMock("https://some-url.com/some-id", nil, "", nil, "", errors.New("transcription failed"))
	response := `{"Data":{"data":"transcription failed"},"Message":"error","Status":400}`
	requestBody := []byte(`{"audio": "SSBhbSBhIHZhbGlkIGJ5dGUgYXJyYXk="}`)
	reader := bytes.NewReader(requestBody)
	r := SetRouter()
	r.POST("/stt", SpeechToText(client, 1))
	req, _ := http.NewRequest("POST", "/stt", reader)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, response, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
