package routers

import (
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    // "github.com/google/uuid"
    "MumiChat/pkg/setting"
    "MumiChat/middleware"
    // "MumiChat/pkg/e"
    "net/http"
    "encoding/json"
    "log"
    // "fmt"
)
type Packet struct {
    Type string 
    Action string
    Data map[string]string
}

// Upgrade method from an HTTP request handler to get a *Conn:
// Read and Write buffer size
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}
func Routes(conn *websocket.Conn, p []byte) {
    var packet Packet;
    err := json.Unmarshal(p, &packet)
    if err != nil {
        log.Println("error:", err)
    }
    if(packet.Type == "action") {
        Action(conn, packet.Action, packet.Data)
    } else if(packet.Type == "message") {
        middleware.CheckIsLogin(conn, packet.Data)
    }
}
func wsReader(conn *websocket.Conn) {
    for {
    // read in a message
        _, p, err := conn.ReadMessage()
        if err != nil {
            Logout()
            log.Println(err)
            return
        }
    // print out that message for clarity
        // fmt.Println(string(p))
        Routes(conn, p)
        // if err := conn.WriteMessage(messageType, p); err != nil {
        //     log.Println(err)
        //     return
        // }

    }
}

func wsEndpoint(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

    // upgrade this connection to a WebSocket
    // connection
    ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Println(err)
	}
    log.Println("Websocket Client Connected")
    
	err = ws.WriteMessage(1, []byte(`{
        "key": "24AB50C2"
    }`))

	defer ws.Close()
	wsReader(ws)
}

func InitRouter() *gin.Engine {
    CleanLoginStatus()
    r := gin.Default() 

    gin.SetMode(setting.RunMode)
    r.GET("/", wsEndpoint)
    r.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "test",
        })
    })
    // apiv1 := r.Group("/api/v1")
    // {
    //     //获取标签列表
    //     apiv1.GET("/tags", v1.GetTags)
    //     //新建标签
    //     apiv1.POST("/tags", v1.AddTag)
    //     //更新指定标签
    //     apiv1.PUT("/tags/:id", v1.EditTag)
    //     //删除指定标签
    //     apiv1.DELETE("/tags/:id", v1.DeleteTag)
    // }

    return r
}