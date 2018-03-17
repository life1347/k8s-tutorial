* Deployment detail
    * https://tachingchen.com/tw/blog/kubernetes-rolling-update-with-deployment/

* kubectl command

``` bash
$ kubectl run test-depl-1 --image=nginx:latest --labels="app=nginx" --expose --port=80
```