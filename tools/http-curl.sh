#!/bin/bash
#用curl模拟kubectl exec podname ls 命令
url=`curl --insecure -I -s -H "X-Stream-Protocol-Version: v4.channel.k8s.io" -H "X-Stream-Protocol-Version: channel.k8s.io" -H "Connection: Upgrade" -H "Upgrade: websocket" -X POST "https://kubernetes-master:10250/exec/default/ng-0/web-server?command=ls&error=1&output=1" --cert /home/ubuntu/apiserver-kubelet-client.crt --key /home/ubuntu/apiserver-kubelet-client.key -o /dev/null -w %{redirect_url}`
wscat -c "${url}" --cert /home/ubuntu/apiserver-kubelet-client.crt --key /home/ubuntu/apiserver-kubelet-client.key -n 
