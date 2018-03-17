* Pod detail
    * https://tachingchen.com/tw/blog/kubernetes-pod/
    * https://tachingchen.com/tw/blog/kubernetes-assigning-pod-to-nodes/

* kubectl command

``` bash
$ kubectl apply -f pod.yaml
$ kubectl logs -f -c <container> <pod>
$ kubectl port-forward <pod> <local port>:<container port>
```