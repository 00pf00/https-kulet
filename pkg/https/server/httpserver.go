package server

import (
	"00pf00/https-kulet/pkg/https/client"
	"00pf00/https-kulet/pkg/https/client/gorillawebsocket"
	"00pf00/https-kulet/pkg/util"
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
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
	s := &http.Server{
		Addr:              server.Addr,
		Handler:           server,
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

func (server *HttpServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if strings.Index(request.URL.String(), "exec") > 0 {
		EXEC(writer, request)
	} else if strings.Index(request.URL.String(), "cri") > 0 {
		CRI(writer, request)
	}

}

func EXEC(writer http.ResponseWriter, request *http.Request) {
	cert, err := tls.LoadX509KeyPair(util.CLIENT_CERT, util.CLIENT_KEY)
	if err != nil {
		fmt.Printf("client load cert fail certpath = %s keypath = %s \n", util.CLIENT_CERT, util.CLIENT_KEY)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	}
	httpclient := &http.Client{
		Transport:     tr,
		CheckRedirect: RD,
	}

}

func CRI(writer http.ResponseWriter, request *http.Request) {
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
}

//处理回调
func RD(req *http.Request, via []*http.Request) error {
	http.Redirect(writer, request, "/cri/a", http.StatusFound)
	return nil
}
