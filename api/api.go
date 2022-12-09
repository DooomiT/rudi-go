package api

import (
	"fmt"

	"github.com/DooomiT/rudi-go/api/stt"
	"github.com/DooomiT/rudi-go/client/assemblyai"
	"github.com/gin-gonic/gin"
)

type API struct {
	router  *gin.Engine
	version string
}

func (api *API) versionedEndpoint(endpoint string) string {
	return fmt.Sprintf("/%s/%s", api.version, endpoint)
}

func New(version string, client assemblyai.AssemblyAi, maxRetries uint) *API {
	router := gin.Default()
	api := API{router, version}
	api.router.POST(api.versionedEndpoint("/stt"), stt.SpeechToText(client, maxRetries))
	return &api
}

func (api *API) Run(port string) error {
	return api.router.Run(":" + port)
}
