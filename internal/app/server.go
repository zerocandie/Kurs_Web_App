package app

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (a *Application) StartServer() {
	log.Println("Server start up")

	r := gin.Default()

	r.GET("/ping/:name", a.Ping)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	log.Println("Server down")
}

type pingReq struct{}
type pingResp struct {
	Status string `json:"status"`
}

// Ping godoc
// @Summary      Show hello text
// @Description  very very friendly response
// @Tags         Tests
// @Produce      json
// @Success      200  {object}  pingResp
// @Router       /ping/{name} [get]
func (a *Application) Ping(gCtx *gin.Context) {
	name := gCtx.Param("name")
	gCtx.String(http.StatusOK, "Hello %s", name)
}
