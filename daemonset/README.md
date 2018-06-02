### Deployment detail
    * https://tachingchen.com/tw/blog/kubernetes-rolling-update-with-deployment/

### Create deployment from CLI

``` bash
$ kubectl run test-depl-1 --image=nginx:latest --labels="app=nginx" --expose --port=80
```

### Rolling update commands

* apply

```
$ kubectl apply -f <file> --record
```

* set image

```
$ kubectl set image deployment <deployment> <container>=<image> --record
```

* replace

```
$ kubectl replace -f <yaml> --record
```

* edit

```
$ kubectl edit deployment <deployment> --record
```

* Check rolling update status

```
$ kubectl rollout status deployment nginx
```

* Pause rolling update

```
$ kubectl rollout pause deployment <deployment>
```

* Resume rolling update

```
$ kubectl rollout resume deployment <deployment>
```

* Show rolling update history

```
$ kubectl rollout history deployment <deployment>
```

* Rollback to previous/specific revision

```
$ kubectl rollout undo deployment <deployment>
$ kubectl rollout undo deployment <deployment> --to-revision=<revision>
```
