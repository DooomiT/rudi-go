package sttlocal

import (
	"fmt"
	"net/http"

	"github.com/DooomiT/rudi-go/api/stt"
	"github.com/DooomiT/rudi-go/pkg/audio"

	"github.com/asticode/go-asticoqui"
	"github.com/gin-gonic/gin"
)

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

		data, err := audio.AudioToWav(dto.Audio)
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
