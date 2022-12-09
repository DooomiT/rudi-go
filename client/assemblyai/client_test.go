package assemblyai

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
	client := New(*http.DefaultClient, server.URL, "some-token")
	assert.NotEmpty(t, client)
}

func TestUploadLocalFile(t *testing.T) {
	server := getServer(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte(`{
			"upload_url": "https://cdn.assemblyai.com/upload/f4932e0c-4f0a-40b8-8994-bdae0c0980fb"
		  }`))
	})
	defer server.Close()
	client := New(*http.DefaultClient, server.URL, "some-token")

	uploadUrl, err := client.UploadLocalFile([]byte{})
	assert.NoError(t, err)
	assert.Equal(t, "https://cdn.assemblyai.com/upload/f4932e0c-4f0a-40b8-8994-bdae0c0980fb", uploadUrl)
}

func TestUploadLocalFileBadRequest(t *testing.T) {
	server := getServer(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(400)
		res.Write([]byte(`{}`))
	})
	defer server.Close()
	client := New(*http.DefaultClient, server.URL, "some-token")

	uploadUrl, err := client.UploadLocalFile([]byte{})
	assert.Error(t, err)
	assert.Equal(t, "", uploadUrl)
}

func TestUploadLocalFileBadContent(t *testing.T) {
	server := getServer(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte(`{"upload_url": 1}`))
	})
	defer server.Close()
	client := New(*http.DefaultClient, server.URL, "some-token")

	uploadUrl, err := client.UploadLocalFile([]byte{})
	assert.Error(t, err)
	assert.Equal(t, "", uploadUrl)
}

func TestTranscribe(t *testing.T) {
	server := getServer(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte(`{
			"id": "5551722-f677-48a6-9287-39c0aafd9ac1",
			"status": "completed",
			"text": "You know Demons on TV like that and and for people to expose themselves to being rejected on TV or humiliated by fear factor or."
		}`))
	})
	defer server.Close()
	client := New(*http.DefaultClient, server.URL, "some-token")

	text, err := client.Transcript("https://some-url.com/some-id")
	assert.NoError(t, err)
	assert.Equal(t, "5551722-f677-48a6-9287-39c0aafd9ac1", text)
}

func TestTranscribeError(t *testing.T) {
	server := getServer(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte(`{
			"id": "5551722-f677-48a6-9287-39c0aafd9ac1",
			"status": "error",
			"text": "null"
		}`))
	})
	defer server.Close()
	client := New(*http.DefaultClient, server.URL, "some-token")

	text, err := client.Transcript("https://some-url.com/some-id")
	assert.Error(t, err)
	assert.Equal(t, "", text)
}

func TestPollTranscribe(t *testing.T) {
	server := getServer(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte(`{
			"id": "5551722-f677-48a6-9287-39c0aafd9ac1",
			"status": "completed",
			"text": "You know Demons on TV like that and and for people to expose themselves to being rejected on TV or humiliated by fear factor or."
		}`))
	})
	defer server.Close()
	client := New(*http.DefaultClient, server.URL, "some-token")

	text, err := client.PollTranscript("https://some-url.com/some-id", 1)
	assert.NoError(t, err)
	assert.Equal(t, "You know Demons on TV like that and and for people to expose themselves to being rejected on TV or humiliated by fear factor or.", text)
}

func TestPollTranscribeError(t *testing.T) {
	server := getServer(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte(`{
			"id": "5551722-f677-48a6-9287-39c0aafd9ac1",
			"status": "error",
			"text": "null"
		}`))
	})
	defer server.Close()
	client := New(*http.DefaultClient, server.URL, "some-token")

	text, err := client.PollTranscript("5551722-f677-48a6-9287-39c0aafd9ac1", 1)
	assert.Error(t, err)
	assert.Equal(t, "", text)
}

func TestPollTranscribeTimeout(t *testing.T) {
	server := getServer(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Write([]byte(`{
			"id": "5551722-f677-48a6-9287-39c0aafd9ac1",
			"status": "queued",
			"text": "null"
		}`))
	})
	defer server.Close()
	client := New(*http.DefaultClient, server.URL, "some-token")

	text, err := client.PollTranscript("5551722-f677-48a6-9287-39c0aafd9ac1", 1)
	assert.Error(t, err)
	assert.Equal(t, "", text)
}
