package server

import (
	"00pf00/https-kulet/pkg/util"
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type HttpServer struct {
	Cert string
	Key  string
	Addr string
}

func (server *HttpServer) StartServer() {
	cert, err := tls.LoadX509KeyPair(server.Cert, server.Key)
	if err != nil {
		fmt.Printf("client load cert fail certpath = %s keypath = %s \n", server.Cert, server.Key)
		return
	}
	config := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	mux := &http.ServeMux{}
	mux.HandleFunc("/exec", func(writer http.ResponseWriter, request *http.Request) {
		http.Redirect(writer, request, "/cri", http.StatusFound)
	})
	mux.HandleFunc("/cri", func(writer http.ResponseWriter, request *http.Request) {
		upgrader := websocket.Upgrader{}
		c, err := upgrader.Upgrade(writer, request, nil)
		if err != nil {
			fmt.Printf("upgrade fail err = %v\n", err)
		}
		defer c.Close()
		running := true
		for running {
			err = c.WriteMessage(websocket.TextMessage, []byte{'a'})
			if err != nil {
				fmt.Printf("websocket  write fail err = %v\n", err)
				running = false
			}
			time.Sleep(1 * time.Second)

		}
	})

	s := &http.Server{
		Addr:              server.Addr,
		Handler:           mux,
		TLSConfig:         config,
		ReadTimeout:       0,
		ReadHeaderTimeout: 0,
		WriteTimeout:      0,
		IdleTimeout:       0,
		MaxHeaderBytes:    0,
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	err = s.ListenAndServeTLS("", "")
	if err != nil {
		fmt.Printf("server start fail err = %v\n", err)
	}
}

func NewHttpServer() *HttpServer {
	return &HttpServer{
		Cert: util.SERVER_CERT,
		Key:  util.SERVER_KEY,
		Addr: "0.0.0.0:10250",
	}
}
