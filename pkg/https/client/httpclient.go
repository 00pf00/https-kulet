package client

import (
	"00pf00/https-kulet/pkg/util"
	"crypto/tls"
	"fmt"
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
	}
	request, err := http.NewRequest("POST",client.Url,nil)

	if err != nil {
		fmt.Printf("get request fail url = %s \n",client.Url)
	}

	request.Header.Add("X-Stream-Protocol-Version","v2.channel.k8s.io")
	request.Header.Add("X-Stream-Protocol-Version","channel.k8s.io")
	response,err := httpclient.Do(request)
	if err != nil {
		fmt.Printf("response fail err = %v \n",err)
		return
	}
	heads := response.Header
	fmt.Printf("Locatin = %v",heads.Get("Location"))
}
