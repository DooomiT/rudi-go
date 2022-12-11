package api

import (
	"fmt"

	assemblyai "github.com/DooomiT/assembly-ai-go/pkg"
	"github.com/DooomiT/rudi-go/api/stt"
	"github.com/DooomiT/rudi-go/api/sttlocal"
	"github.com/asticode/go-asticoqui"
	"github.com/gin-gonic/gin"
)

type API struct {
	router  *gin.Engine
	version string
}

func (api *API) versionedEndpoint(endpoint string) string {
	return fmt.Sprintf("/%s/%s", api.version, endpoint)
}

func New(version string, client assemblyai.AssemblyAI, maxRetries uint) *API {
	router := gin.Default()
	api := API{router, version}
	api.router.POST(api.versionedEndpoint("/stt"), stt.SpeechToText(client))
	return &api
}

func NewLocal(version string, model *asticoqui.Model) *API {
	router := gin.Default()
	api := API{router, version}
	api.router.POST(api.versionedEndpoint("/stt"), sttlocal.SpeechToText(model))
	return &api
}

func (api *API) Run(port string) error {
	return api.router.Run(":" + port)
}
