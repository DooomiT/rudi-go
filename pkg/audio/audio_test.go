package audio

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/assert"
)

func readAudioFile(relPath string) ([]byte, error) {
	dir, _ := os.Getwd()
	var absPath = fmt.Sprint(strings.TrimSuffix(dir, "pkg/audio"), relPath)
	data, err := os.ReadFile(absPath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func TestAudioToWav(t *testing.T) {
	var tests = []struct {
		name      string
		inputFile string
		want      []int16
		wantErr   bool
	}{
		{"Wrong audio format (m4a)", "audio-files/Untitled.m4a", nil, true},
		{"Valid audio fromat (wav)", "audio-files/2830-3980-0043.wav", []int16{}, false},
		{"Valid audio fromat (wav)", "audio-files/4507-16021-0012.wav", []int16{}, false},
		{"Valid audio fromat (wav)", "audio-files/8455-210777-0068.wav", []int16{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := readAudioFile(tt.inputFile)
			assert.NoError(t, err)
			got, gotErr := AudioToWav(data)
			if tt.wantErr {
				assert.Error(t, gotErr)
			} else {
				assert.NoError(t, gotErr)
			}
			snaps.MatchSnapshot(t, got)
		})
	}
}
