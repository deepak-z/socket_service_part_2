package main

import (
    "github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	websocket "github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
    conn, err := wsupgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Failed to set websocket upgrade: %+v", err)
        return
    }

    for {
        t, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }
        conn.WriteMessage(t, msg)
    }
}

func main() {

    r := gin.Default()


	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	r.LoadHTMLFiles("index.html")

r.GET("/", func(c *gin.Context) {
    c.HTML(200, "index.html", nil)
})

    r.Run("localhost:12312")
}