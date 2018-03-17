* Service detail
    * https://tachingchen.com/tw/blog/kubernetes-service/

* kubectl command

``` bash
$ kubectl apply -f *.yaml
$ kubectl expose pod <pod> —type=<svc type> —name=<svc name> --port <expose port> —target-port=<container port>
```
