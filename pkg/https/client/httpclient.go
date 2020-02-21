package client

import (
	"00pf00/https-kulet/pkg/https/client/gorillawebsocket"
	"00pf00/https-kulet/pkg/util"
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
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
	//tr := &http.Transport{
	//	TLSClientConfig: &tls.Config{
	//		Certificates:       []tls.Certificate{cert},
	//		InsecureSkipVerify: true,
	//	},
	//}
	tlsconfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	//httpclient := &http.Client{
	//	Transport: tr,
	//	//CheckRedirect: gorillawebsocket.LSRD,
	//}
	request, err := http.NewRequest("POST", client.Url, nil)

	if err != nil {
		fmt.Printf("get request fail url = %s \n", client.Url)
	}

	request.Header.Add("X-Stream-Protocol-Version", "v4.channel.k8s.io")
	request.Header.Add("X-Stream-Protocol-Version", "v3.channel.k8s.io")
	request.Header.Add("X-Stream-Protocol-Version", "v2.channel.k8s.io")
	request.Header.Add("X-Stream-Protocol-Version", "channel.k8s.io")
	request.Header.Add("Connection", "Upgrade")
	request.Header.Add("Upgrade", "SPDY/3.1")
	conn,err := tls.Dial("tcp","170.106.72.202:10250",tlsconfig)
	//body, err := httpclient.Do(request)
	//if err != nil && body.StatusCode != http.StatusFound {
	//	fmt.Printf("response fail err = %v \n", err)
	//	return
	//}
	if err != nil {

	}
	request.Write(conn)
	rawResponse := bytes.NewBuffer(make([]byte, 0, 1024))
	rawResponse.Reset()
	respReader :=bufio.NewReader(io.TeeReader(io.LimitReader(conn,1024),rawResponse))
	resp,err := http.ReadResponse(respReader,nil)
	if err != nil {

	}
	fmt.Printf("%s",string(resp.StatusCode))
}

func (client *HttpClient) LSDR() {
	wss := "https://token:10250/exec/default/ng-0/web-server?command=ls&error=1&output=1"
	cert, err := tls.LoadX509KeyPair(util.CLIENT_CERT, util.CLIENT_KEY)
	if err != nil {
		fmt.Printf("client load cert fail certpath = %s keypath = %s \n", util.CLIENT_KEY, util.CLIENT_KEY)
		return
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates:       []tls.Certificate{cert},
			InsecureSkipVerify: true,
		},
	}
	httpclient := &http.Client{
		Transport: tr,
	}
	request, err := http.NewRequest("GET", wss, nil)
	request.Header.Add("Connection", "Upgrade")
	request.Header.Add("Upgrade", "websocket")
	request.Header.Add("Sec-Websocket-Version", "13")
	request.Header.Add("Sec-Websocket-Key", "abcd")

	if err != nil {
		fmt.Printf("get request fail url = %s \n", client.Url)
	}
	body, err := httpclient.Do(request)
	if err != nil {
		fmt.Printf("response fail err = %v \n", err)
		return
	}
	if body.StatusCode == http.StatusSwitchingProtocols {
		conn, ok := body.Body.(io.ReadWriteCloser)
		if ok {
			wrunning := true
			for wrunning {
				rb := make([]byte,1024)
				n, err := conn.Read(rb)
				if err != nil {
					fmt.Printf("httpclient write fail err = %v\n", err)
					wrunning = false
				}
				fmt.Printf("receive msg = %s \n",rb[:n])
			}
		}
	}
	running := true
	for running {
		msg := make([]byte, 1024)
		_, err := body.Body.Read(msg)
		if err != nil {
			fmt.Printf("body read fail err = %v\n", err)
			running = false
		}
		fmt.Printf("msg = %s\n", string(msg))
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

//模拟kubectl exec podname ls
func (client *HttpClient) GET() {
	//client.Url = "https://49.51.38.39:10250/containerLogs/default/ng-0/web-server?follow=true"
	client.Url = "https://token:10250/containerLogs/default/ng-0/web-server?follow=true"
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
	//request, err := http.NewRequest("GET", client.Url, nil)
	////request.Header.Add("Content-Type","application/json")
	//if err != nil {
	//	fmt.Printf("get request fail url = %s \n", client.Url)
	//}
	//
	//body, err := httpclient.Do(request)
	body ,err := httpclient.Get(client.Url)
	if err != nil && body.StatusCode != http.StatusFound {
		fmt.Printf("response fail err = %v \n", err)
		return
	}
	running := true
	for running {
		msg := make([]byte,512)
		n,err := body.Body.Read(msg)
		if err != nil {
			fmt.Printf("err = %v",err)
			running = false
		}
		fmt.Printf("msg = %s\n",string(msg[:n]))
	}

	//rsp, err := ioutil.ReadAll(body.Body)
	//if err == nil {
	//	fmt.Printf("rsp = %s\n", string(rsp))
	//}
}
