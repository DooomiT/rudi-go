package assemblyai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type AssemblyAi interface {
	// UploadLocalFile uploads binary data to AassemblyAi
	// It returs the upload_url
	UploadLocalFile(content []byte) (string, error)
	// Transcript creates a transcription job at AssemblyAi
	// It returns the id of the job
	Transcript(audioUrl string) (string, error)
	// Transcript polls a transcription job at assemblyAi
	// It returns the result of the job
	PollTranscript(id string, maxTries uint) (string, error)
}

type AssembyAiImpl struct {
	http.Client
	baseUrl string
	token   string
}

func New(client http.Client, baseUrl, token string) AssemblyAi {
	return &AssembyAiImpl{client, baseUrl, token}
}

func isValidStatus(statusCode int) bool {
	okStatusRegex := regexp.MustCompile(`^2..`)
	s := strconv.Itoa(statusCode)
	return okStatusRegex.MatchString(s)
}

func getBody(response *http.Response) ([]byte, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, err
}

func getData[T any](response *http.Response) (*T, error) {
	body, err := getBody(response)
	if err != nil {
		return nil, err
	}
	var data T
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

type UploadLocalFileResponse struct {
	UploadUrl string `json:"upload_url"`
}

func (client *AssembyAiImpl) UploadLocalFile(content []byte) (string, error) {
	req, err := http.NewRequest("POST", client.baseUrl+"/upload", bytes.NewBuffer(content))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", client.token)
	req.Header.Set("transfer-encoding", "chunked")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if !isValidStatus(resp.StatusCode) {
		errBody, _ := getBody(resp)
		return "", fmt.Errorf(string(errBody))
	}
	data, err := getData[UploadLocalFileResponse](resp)
	if err != nil {
		return "", err
	}
	return data.UploadUrl, nil
}

type TranscriptDto struct {
	AudioUrl string `json:"audio_url"`
}

type TranscriptResponse struct {
	Id     string `json:"id"`
	Status string `json:"status"`
	Text   string `json:"text"`
	Error  string `json:"error"`
}

func (client *AssembyAiImpl) PollTranscript(id string, maxTries uint) (string, error) {
	url := fmt.Sprintf("%s/transcript/%s", client.baseUrl, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("authorization", client.token)
	if err != nil {
		return "", err
	}
	var n uint = 0
	for n < maxTries {
		n++
		resp, err := client.Do(req)
		if err != nil {
			return "", err
		}
		if !isValidStatus(resp.StatusCode) {
			errBody, _ := getBody(resp)
			return "", fmt.Errorf(string(errBody))
		}
		data, err := getData[TranscriptResponse](resp)
		if err != nil {
			return "", err
		}
		if data.Status == "error" {
			return "", errors.New(data.Error)
		}
		if data.Status == "completed" {
			return data.Text, nil
		}
		if data.Status == "queued" {
			time.Sleep(time.Second)
		}
	}
	return "", fmt.Errorf("no transcription received, exceeded max tries %d", maxTries)
}

func (client *AssembyAiImpl) Transcript(audioUrl string) (string, error) {
	dto := TranscriptDto{AudioUrl: audioUrl}
	body, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", client.baseUrl+"/transcript", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", client.token)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if !isValidStatus(resp.StatusCode) {
		errBody, _ := getBody(resp)
		return "", fmt.Errorf(string(errBody))
	}
	data, err := getData[TranscriptResponse](resp)

	if err != nil {
		return "", err
	}
	if data.Status == "error" {
		return "", errors.New(data.Error)
	}
	return data.Id, nil
}
