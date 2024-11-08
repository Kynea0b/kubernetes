
start:
	minikube start

create-pod: start
	kubectl create -f pod-nginx.yaml

show: start
	kubectl get pod

crate-echo:
	kubectl create deployment hello-minikube --image=registry.k8s.io/echoserver:1.10
	#https://kubernetes.io/ja/docs/setup/learning-environment/minikube/

create-echo2:
	kubectl create deployment hello-node --image=registry.k8s.io/e2e-test-images/agnhost:2.39 -- /agnhost netexec --http-port=8080
	#https://kubernetes.io/ja/docs/tutorials/hello-minikube/

stop:
	minikube stop
