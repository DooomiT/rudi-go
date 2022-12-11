package sttlocal

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net/http"

	"github.com/DooomiT/rudi-go/api/stt"
	"github.com/asticode/go-asticoqui"
	"github.com/cryptix/wav"
	"github.com/gin-gonic/gin"
)

func audioToWav(audio []byte) ([]int16, error) {
	r, err := wav.NewReader(bytes.NewReader(audio), int64(binary.Size(audio)))
	if err != nil {
		return nil, fmt.Errorf("creating new reader failed: %w", err)
	}
	var d []int16
	for {
		s, err := r.ReadSample()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("reading sample failed: %w", err)
		}
		d = append(d, int16(s))
	}
	return d, nil
}

func SpeechToText(model *asticoqui.Model) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto stt.SpeechToTextDto
		if err := ctx.BindJSON(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}

		data, err := audioToWav(dto.Audio)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}
		res, err := model.SpeechToText(data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": fmt.Sprint("Failed converting speech to text: ", err)},
			})
			return
		}

		results := []string{res}
		respData := stt.SpeechToTextResponse{Text: results[0]}
		ctx.JSON(http.StatusOK, respData)
	}
}
