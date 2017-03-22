# Sample go app to be deployed in different cloud implementation

## Pre requisite

After cloning

```
./build-docker.sh
docker push vincentserpoul/testdeploygolang
```

You will need to change the name in the different yml config files


## In docker swarm

```
docker deploy --compose-file ./compose.yml testdeploy
```

## In kubernetes

```
kubectl apply -f kube.yml
```

if you are using [minikube](https://github.com/kubernetes/minikube):
```
minikube start
minikube ip
kubectl get svc
```

get the IP and the forwarded port to access your service

## In Pivotal cloud foundry

PCF will autodetect everything

```
cf push
```

if you are using [pcfdev](https://pivotal.io/platform/pcf-tutorials/getting-started-with-pivotal-cloud-foundry-dev/introduction)

```
cf login -a api.local.pcfdev.io --skip-ssl-validation
cf push --hostname testdeploy
```

You should be able to access the deployment at testdeploy.local.pcfdev.io
