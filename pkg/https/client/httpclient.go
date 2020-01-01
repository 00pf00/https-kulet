package client

import (
	"00pf00/https-kulet/pkg/util"
	"crypto/tls"
	"fmt"
	"github.com/golang/net/websocket"
	"net/http"
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
		Url:      util.SERVER_ADDR,
	}
}

func (client *HttpClient) Post() {
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
		CheckRedirect: Redirect,
	}
	request, err := http.NewRequest("POST", client.Url, nil)

	if err != nil {
		fmt.Printf("get request fail url = %s \n", client.Url)
	}

	request.Header.Add("X-Stream-Protocol-Version", "v2.channel.k8s.io")
	request.Header.Add("X-Stream-Protocol-Version", "channel.k8s.io")
	request.Header.Add("Connection", "Upgrade")
	request.Header.Add("Upgrade", "websocket")
	_, err = httpclient.Do(request)
	if err != nil {
		fmt.Printf("response fail err = %v \n", err)
		return
	}
}

func Redirect(req *http.Request, via []*http.Request) error {
	host := req.URL.Host
	path := req.URL.Path
	ws := "ws://" + host + path
	origin := "https://" + host
	wscli, err := websocket.Dial(ws, "", origin)
	if err != nil {
		fmt.Printf("websocket connection failed url = %s \n", ws)
		return err
	}
	running := true
	for running {
		var msg = make([]byte, 512)
		n, err := wscli.Read(msg)
		if err != nil {
			fmt.Printf("websocket read file err = %v \n", err)
			return err
		}
		fmt.Print(string(msg[:n]))
	}

	fmt.Printf("redirect url = %s \n", req.URL.String())
	return nil
}
