# deploy

kubectl apply -f cockroach-statefulset.yaml
kubectl apply -f cockroach-service.yaml



#-----

kubectl -n lemontea get pods
kubectl -n lemontea get statefulsets.apps
kubectl -n lemontea get service
kubectl -n lemontea get poddisruptionbudgets.policy

kubectl apply -f cockroach-init-job.yaml


kubectl -n lemontea get jobs.batch


kubectl -n lemontea get pods

