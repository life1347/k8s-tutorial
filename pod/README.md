* Pod detail
    * https://tachingchen.com/tw/blog/kubernetes-pod/
    * https://tachingchen.com/tw/blog/kubernetes-assigning-pod-to-nodes/

* kubectl command

``` bash
$ kubectl apply -f pod.yaml
```

* 查看 pod log 

``` bash
$ kubectl logs -f -c <container> <pod>
```

* Portforward 測試
    * https://kubernetes.io/docs/tasks/access-application-cluster/port-forward-access-application-cluster/

``` bash
$ kubectl port-forward <pod> <local port>:<container port>
```

* 進入 container 執行 shell
    * https://kubernetes.io/docs/tasks/debug-application-cluster/get-shell-running-container/

``` bash
# kubectl exec -it -c nginx nginx-pod-1 sh
$ kubectl exec -it -c <container> <pod> <shell>
```
