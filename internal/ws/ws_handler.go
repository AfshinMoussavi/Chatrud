package ws

import (
	"Chat-Websocket/pkg/loggerPkg"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type wsHandler struct {
	service  IWsService
	upgrader websocket.Upgrader
	logger   loggerPkg.ILogger
}

func NewHandler(s IWsService, logger loggerPkg.ILogger) IWsHandler {
	return &wsHandler{
		service: s,
		logger:  logger,
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (h *wsHandler) HandleWebSocket(c *gin.Context) {
	room := c.Query("room")
	token := c.Query(("token"))
	if room == "" {
		h.logger.Error("No room specified in query")
		c.JSON(http.StatusBadRequest, gin.H{"error": "No room specified in query"})
		return
	}
	if token == "" {
		h.logger.Error("No token specified in query")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
		return
	}

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upgrade to WebSocket"})
		return
	}

	if err := h.service.HandleConnection(room, conn, token); err != nil {
		log.Printf("Error handling connection in room %s: %v", room, err)
		conn.Close()
	}

}
