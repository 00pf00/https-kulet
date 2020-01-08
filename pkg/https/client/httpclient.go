package client

import (
	"00pf00/https-kulet/pkg/https/client/gorillawebsocket"
	"00pf00/https-kulet/pkg/util"
	"crypto/tls"
	"fmt"
	"golang.org/x/net/websocket"
	"net"
	"net/http"
	"time"
)

type HttpClient struct {
	CertPath string
	KeyPath  string
	Url      string
}

func NewClient() *HttpClient {
	return &HttpClient{
		CertPath: util.CLIENT_CERT,
		KeyPath:  util.CLIENT_KEY,
	}
}

//模拟kubectl exec podname ls
func (client *HttpClient) LS() {
	client.Url = util.COMMAND_LS
	cert, err := tls.LoadX509KeyPair(client.CertPath, client.KeyPath)
	if err != nil {
		fmt.Printf("client load cert fail certpath = %s keypath = %s \n", client.CertPath, client.KeyPath)
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
		CheckRedirect: gorillawebsocket.LSRD,
	}
	request, err := http.NewRequest("POST", client.Url, nil)

	if err != nil {
		fmt.Printf("get request fail url = %s \n", client.Url)
	}

	request.Header.Add("X-Stream-Protocol-Version", "v2.channel.k8s.io")
	request.Header.Add("X-Stream-Protocol-Version", "channel.k8s.io")
	request.Header.Add("Connection", "Upgrade")
	request.Header.Add("Upgrade", "websocket")
	body, err := httpclient.Do(request)
	if err != nil && body.StatusCode != http.StatusFound {
		fmt.Printf("response fail err = %v \n", err)
		return
	}
}

//模拟kubectl exec -it podname /bin/bash
func (client *HttpClient) BASH() {
	client.Url = util.COMMAND_BASH
	cert, err := tls.LoadX509KeyPair(client.CertPath, client.KeyPath)
	if err != nil {
		fmt.Printf("client load cert fail certpath = %s keypath = %s \n", client.CertPath, client.KeyPath)
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
		CheckRedirect: gorillawebsocket.BASHRD,
	}
	request, err := http.NewRequest("POST", client.Url, nil)

	if err != nil {
		fmt.Printf("get request fail url = %s \n", client.Url)
	}

	request.Header.Add("X-Stream-Protocol-Version", "v2.channel.k8s.io")
	request.Header.Add("X-Stream-Protocol-Version", "channel.k8s.io")
	request.Header.Add("Connection", "Upgrade")
	request.Header.Add("Upgrade", "websocket")
	body, err := httpclient.Do(request)
	if err != nil && body.StatusCode != http.StatusFound {
		fmt.Printf("response fail err = %v \n", err)
		return
	}
}

func LSRD(req *http.Request, via []*http.Request) error {
	wss := req.URL
	wss.Scheme = "wss"
	orignUrl := via[0].URL
	cert, err := tls.LoadX509KeyPair(util.CLIENT_CERT, util.CLIENT_KEY)
	if err != nil {
		fmt.Printf("client load cert fail certpath = %s keypath = %s \n", util.CLIENT_KEY, util.CLIENT_KEY)
		return err
	}
	config := &websocket.Config{
		Location: wss,
		Origin:   orignUrl,
		Protocol: nil,
		Version:  13,
		TlsConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
		Header: nil,
		Dialer: &net.Dialer{
			KeepAlive: 1 * time.Second,
		},
	}
	wscli, err := websocket.DialConfig(config)
	if err != nil {
		fmt.Printf("websocket connection failed url = %s \n", wss.Host)
		return err
	}
	running := true
	for running {
		var msg = make([]byte, 1024)
		n, err := wscli.Read(msg)
		fmt.Printf("read %d \n", n)
		if err != nil {
			fmt.Printf("websocket read file err = %v \n", err)
			return err
		}
		fmt.Print(string(msg[:n]))
	}
	return nil
}
