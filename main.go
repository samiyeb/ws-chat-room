package main

import (
    "golang.org/x/net/websocket"
    "log"
    "net/http"
    "fmt"
    "io"
    "net/url"
    // "strconv"
)


type Server struct {
    serverUrl *url.URL
    connections map[*websocket.Conn]bool
}

func newServer() *Server {
    newUrl, err := url.Parse("http://localhost:8080/ws")
    if err != nil {
        log.Fatal(err)
    }
    return &Server{connections: make(map[*websocket.Conn]bool), serverUrl: newUrl}
}

func (server *Server) handleWS(ws *websocket.Conn){
    fmt.Println("Incoming connection from client", ws.RemoteAddr())

    server.connections[ws] = true

    server.readLoop(ws)
}

func (server *Server) readLoop(ws *websocket.Conn){
    buf := make([]byte, 1024)

    for {
        if n, err := ws.Read(buf); err != nil {
            if err == io.EOF {
                break
            }
            fmt.Println("Error: ", err)
            continue
        } else {
            message := buf[:n]
            server.broadcast(message)
        }
    }
}

func (server *Server) broadcast(bytes []byte) {
    for c, v := range server.connections {
        if v {
            go func(ws *websocket.Conn) {
                if _, err := ws.Write(bytes); err != nil {
                    fmt.Println("Error: ", err)
                } 
            }(c)
        }
    }
}




func main() {
    server := newServer()
    // client := newClient("Sam", 8081)

    // var rwc io.ReadWriteCloser
    // ws, err := websocket.NewClient(&websocket.Config{Location: server.serverUrl, Origin: client.clientUrl }, rwc )
    // if err != nil {
    //     log.Fatal(err)
    // }
    
    http.Handle("/ws", websocket.Handler(server.handleWS))
    
    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}

