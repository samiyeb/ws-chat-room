package main

import (
    // "golang.org/x/net/websocket"
    // "log"
    // "net/http"
    // "fmt"
    // "io"
    // "net/url"
    // "strconv"
)

// type Client struct {
//     name string
//     clientUrl *url.URL
// }

// func newClient(newName string, newPort int) (*Client) {
//     port := strconv.Itoa(newPort)
//     newUrl, err := url.Parse("http://localhost:" + port + "/ws")
//     if err != nil {
//         log.Fatal(err)
//     }
//     return &Client{name: newName, clientUrl: newUrl}
// }

// func (client *Client) sendMessage(ws *websocket.Conn, msg []byte){
//     for {
//         if _, err := ws.Write(msg); err != nil {
//             if err == io.EOF {
//                 break
//             }
//             fmt.Println("Error: ", err)
//             continue
//         }
//     }
// }