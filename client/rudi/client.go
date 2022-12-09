package rudi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/DooomiT/rudi-go/api/stt"
)

type Rudi interface {
	SpeechToText(audio []byte) (string, error)
}

type RudiImpl struct {
	http.Client
	baseUrl string
}

func New(client http.Client, baseUrl string) Rudi {
	return &RudiImpl{client, baseUrl}
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

type SpeechToTextResponse struct {
	Text string `json:"text"`
}

func (client *RudiImpl) SpeechToText(content []byte) (string, error) {
	dto := stt.SpeechToTextDto{Audio: content}
	body, err := json.Marshal(dto)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest("POST", client.baseUrl+"/v1/stt", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if !isValidStatus(resp.StatusCode) {
		errBody, _ := getBody(resp)
		return "", fmt.Errorf(string(errBody))
	}
	data, err := getData[SpeechToTextResponse](resp)
	if err != nil {
		return "", err
	}
	return data.Text, nil
}
