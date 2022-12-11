package rudi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getServer(handler func(res http.ResponseWriter, req *http.Request)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(handler))
}

func TestNew(t *testing.T) {
	server := getServer(func(res http.ResponseWriter, req *http.Request) {})
	defer server.Close()
	client := New(*http.DefaultClient, server.URL)
	assert.NotEmpty(t, client)
}

func TestStt(t *testing.T) {
	server := getServer(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte(`{
			"text": "I am a valid byte array"
		  }`))
	})
	defer server.Close()
	client := New(*http.DefaultClient, server.URL)

	text, err := client.SpeechToText([]byte("SSBhbSBhIHZhbGlkIGJ5dGUgYXJyYXk="))
	assert.NoError(t, err)
	assert.Equal(t, "I am a valid byte array", text)
}

func TestSttError(t *testing.T) {
	server := getServer(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(400)
		res.Write([]byte(`{}`))
	})
	defer server.Close()
	client := New(*http.DefaultClient, server.URL)

	uploadUrl, err := client.SpeechToText([]byte{})
	assert.Error(t, err)
	assert.Equal(t, "", uploadUrl)
}
