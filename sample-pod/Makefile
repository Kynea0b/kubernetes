
start:
	minikube start

create-pod: start
	kubectl create -f pod-nginx.yaml

show: start
	kubectl get pod

crate-echo:
	kubectl create deployment hello-minikube --image=registry.k8s.io/echoserver:1.10


stop:
	minikube stop
