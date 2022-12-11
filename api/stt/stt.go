package stt

import (
	"net/http"

	assemblyai "github.com/DooomiT/assembly-ai-go/pkg"
	"github.com/gin-gonic/gin"
)

type SpeechToTextDto struct {
	Audio []byte `json:"audio"`
}

type SpeechToTextResponse struct {
	Text string `json:"text"`
}

func SpeechToText(client assemblyai.AssemblyAI) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto SpeechToTextDto
		if err := ctx.BindJSON(&dto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}
		audioUrl, err := client.UploadLocalFile(dto.Audio)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}
		id, err := client.Transcript(audioUrl)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}
		transcibedAudio, err := client.PollTranscript(id, nil)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"Status":  http.StatusBadRequest,
				"Message": "error",
				"Data":    map[string]interface{}{"data": err.Error()}})
			return
		}
		respData := SpeechToTextResponse{transcibedAudio}
		ctx.JSON(http.StatusOK, respData)
	}
}
