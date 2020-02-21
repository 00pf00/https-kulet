package util

//client
const (
	//CLIENT_CERT = "./conf/cert/client/kube-apiserver-kubelet-client-cert.crt"
	//CLIENT_KEY  = "./conf/cert/client/kube-apiserver-kubelet-client-key.key"
	//COMMAND_LS = "https://kubernetes-master:10250/exec/default/ng-0/web-server?command=touch&command=ls&input=1&output=1&tty=1"
	CLIENT_CERT = "./conf/cert/client/apiserver-kubelet-client.crt"
	CLIENT_KEY  = "./conf/cert/client/apiserver-kubelet-client.key"
	COMMAND_LS = "https://170.106.72.202:10250/exec/default/ng-0/web-server?command=ls&error=1&output=1"
	//COMMAND_LS = "https://token:10250/exec/default/ng-0/web-server?command=ls&error=1&output=1"
	//COMMAND_BASH = "https://49.51.38.39:10250/exec/default/ng-0/web-server?command=/bin/bash&input=1&output=1&tty=1"
	COMMAND_BASH = "https://127.0.0.1:10250/exec/default/ng-0/web-server?command=/bin/bash&input=1&output=1&tty=1"
	SERVER_CERT  = "./conf/cert/server/kubelet.crt"
	SERVER_KEY   = "./conf/cert/server/kubelet.key"
	//CLIENT_CERT = "./conf/cert/client/kubelet.crt"
	//CLIENT_KEY  = "./conf/cert/client/kubelet.key"
	//COMMAND_LS = "https://192.168.1.24:10250/exec/default/account-ccnpc/account?command=touch&command=ls&input=1&output=1&tty=1"
)
