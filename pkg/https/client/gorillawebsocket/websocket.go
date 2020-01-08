package gorillawebsocket

import (
	"00pf00/https-kulet/pkg/util"
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

//command ls 的重定向回调函数
func LSRD(req *http.Request, via []*http.Request) error {
	wss := req.URL
	wss.Scheme = "wss"
	cert, err := tls.LoadX509KeyPair(util.CLIENT_CERT, util.CLIENT_KEY)
	if err != nil {
		fmt.Printf("client load cert fail certpath = %s keypath = %s \n", util.CLIENT_KEY, util.CLIENT_KEY)
		return err
	}
	dailer := &websocket.Dialer{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	}
	wscli, _, err := dailer.Dial(wss.String(), nil)
	if err != nil {
		fmt.Printf("websocket connection failed url = %s \n", wss.Host)
		return err
	}
	running := true
	for running {
		_, msg, err := wscli.ReadMessage()
		if err != nil {
			fmt.Printf("websocket read file err = %v \n", err)
			return err
		}
		fmt.Print(string(msg) + "\n")
	}
	return nil
}

//command bash 的重定向回调函数
func BASHRD(req *http.Request, via []*http.Request) error {
	wss := req.URL
	wss.Scheme = "wss"
	cert, err := tls.LoadX509KeyPair(util.CLIENT_CERT, util.CLIENT_KEY)
	if err != nil {
		fmt.Printf("client load cert fail certpath = %s keypath = %s \n", util.CLIENT_KEY, util.CLIENT_KEY)
		return err
	}
	dailer := &websocket.Dialer{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	}
	wscli, _, err := dailer.Dial(wss.String(), nil)
	if err != nil {
		fmt.Printf("websocket connection failed url = %s \n", wss.Host)
		return err
	}
	stop := make(chan struct{})
	go func(conn *websocket.Conn, stop chan struct{}) {
		running := true
		for running {
			select {
			default:
				err := conn.WriteMessage(websocket.TextMessage, []byte("ls\n"))
				if err != nil {
					fmt.Printf("bash websocket write fail err = %v\n", err)
					stop <- struct{}{}
					return
				}
				time.Sleep(1 * time.Second)
			case <-stop:
				running = false
				conn.Close()

			}
		}
	}(wscli, stop)
	running := true
	for running {
		select {
		default:
			_, msg, err := wscli.ReadMessage()
			if err != nil {
				fmt.Printf("websocket read file err = %v \n", err)
				stop <- struct{}{}
				return err
			}
			fmt.Print(string(msg) + "\n")
		case <-stop:
			wscli.Close()
			running = false
		}

	}
	return nil
}
