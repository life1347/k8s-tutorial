# Useful command

## Cluster operation

```
# Show all supported k8s version
$ minikube get-k8s-versions

# Start with given k8s version. (Must delete old k8s cluster first)
$ minikube start --kubernetes-version=<version>

# Stop cluster
$ minikube stop

# Delete old cluster
$ minikube delete
```

## Enable addons

```
# Show all available addons
$ minikube addons list

# Enable/Disable addons in "config". The addons config will be used when new cluster is up.
$ minikube addons enable <addon>
$ minikube addons disable <addon>

# Run addons when cluster is already up.
$ minikube open <addon>
```

## Show minikube IP

```
$ minikube ip
```

## SSH login into minikube VM

```
$ minikube ssh
```

## Docker machine ENV
```
$ minikube docker-env
```
